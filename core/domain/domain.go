package domain

import (
	"sync"
)

// order struct containing user order details
type Orders struct {
	OrderId        int     `json:"order_id" orm:"pk;auto;"`
	UserId         int     `json:"user_id"`
	ItemIds        string  `json:"item_ids"`
	TotalAmount    float64 `json:"total_amount"`
	Status         string  `json:"status"`
	ProcessingTime float64 `json:"processing_time"`
}

// Metric struct containing metric details
type Metrics struct {
	TotalOrdersProcessed  int            `json:"total_orders_processed"`
	AverageProcessingTime float64        `json:"average_processing_time"`
	OrderStatusCounts     map[string]int `json:"order_status_counts"`
	mu                    sync.Mutex     `json:"-"` // Mutex for thread-safe metrics updates
}

// Created methods of metric structs since we want Mutex field as private
func (m *Metrics) MutexLock() {
	m.mu.Lock()
}

func (m *Metrics) MutexUnLock() {
	m.mu.Unlock()
}
