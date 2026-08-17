package store

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"shiftboard/internal/model"
)

type Repository struct {
	mu     sync.RWMutex
	orders map[string]model.WorkOrder
}

func NewRepository() *Repository {
	return &Repository{orders: make(map[string]model.WorkOrder)}
}

func (repo *Repository) Save(order model.WorkOrder) error {
	if err := order.Validate(); err != nil {
		return err
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()
	if _, exists := repo.orders[order.ID]; exists {
		return fmt.Errorf("work order %s already exists", order.ID)
	}
	if order.Stage == "" {
		order.Stage = model.StageQueued
	}
	repo.orders[order.ID] = order.Clone()
	return nil
}

func (repo *Repository) SaveAll(ctx context.Context, orders []model.WorkOrder) error {
	saved := make([]string, 0, len(orders))
	for index, order := range orders {
		// 先检查取消，避免取消后仍把工单落库
		if err := ctx.Err(); err != nil {
			repo.rollback(saved)
			return fmt.Errorf("import interrupted: %w", model.ImportCanceledError{Processed: index})
		}
		if err := repo.Save(order); err != nil {
			repo.rollback(saved)
			return err
		}
		saved = append(saved, order.ID)
	}
	return nil
}

// rollback 回滚本批已新增的工单，保证取消导入后面板内容不变
func (repo *Repository) rollback(saved []string) {
	if len(saved) == 0 {
		return
	}
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for _, id := range saved {
		delete(repo.orders, id)
	}
}

func (repo *Repository) Replace(order model.WorkOrder) error {
	if err := order.Validate(); err != nil {
		return err
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()
	if _, exists := repo.orders[order.ID]; !exists {
		return fmt.Errorf("work order %s was not found", order.ID)
	}
	repo.orders[order.ID] = order.Clone()
	return nil
}

func (repo *Repository) Get(id string) (model.WorkOrder, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	order, exists := repo.orders[id]
	if !exists {
		return model.WorkOrder{}, fmt.Errorf("work order %s was not found", id)
	}
	return order.Clone(), nil
}

func (repo *Repository) List() []model.WorkOrder {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	orders := make([]model.WorkOrder, 0, len(repo.orders))
	for _, order := range repo.orders {
		orders = append(orders, order.Clone())
	}
	sort.Slice(orders, func(left, right int) bool {
		return orders[left].ID < orders[right].ID
	})
	return orders
}
