package paging

import (
	"reflect"
	"testing"
)

func TestNewPaginator(t *testing.T) {
	tests := []struct {
		pageIndex int
		pageSize  int
		total     int
		items     interface{}
		expected  *Paginator
	}{
		{
			pageIndex: 1,
			pageSize:  10,
			total:     100,
			items:     []int{1, 2, 3},
			expected: &Paginator{
				PageIndex: 1,
				PageSize:  10,
				Pages:     10,
				Total:     100,
				Items:     []int{1, 2, 3},
			},
		},
		{
			pageIndex: 2,
			pageSize:  5,
			total:     12,
			items:     []string{"a", "b"},
			expected: &Paginator{
				PageIndex: 2,
				PageSize:  5,
				Pages:     3,
				Total:     12,
				Items:     []string{"a", "b"},
			},
		},
	}

	for _, tt := range tests {
		result := NewPaginator(tt.pageIndex, tt.pageSize, tt.total, tt.items)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("NewPaginator(%d, %d, %d, %v) = %v; want %v", tt.pageIndex, tt.pageSize, tt.total, tt.items, result, tt.expected)
		}
	}
}

func TestGetPages(t *testing.T) {
	tests := []struct {
		total    int
		pageSize int
		expected int
	}{
		{total: 100, pageSize: 10, expected: 10},
		{total: 55, pageSize: 10, expected: 6},
		{total: 0, pageSize: 10, expected: 0},
		{total: 10, pageSize: 10, expected: 1},
		{total: 11, pageSize: 10, expected: 2},
	}

	for _, tt := range tests {
		result := getPages(tt.total, tt.pageSize)
		if result != tt.expected {
			t.Errorf("getPages(%d, %d) = %d; want %d", tt.total, tt.pageSize, result, tt.expected)
		}
	}
}
