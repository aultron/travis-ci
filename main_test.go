package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		z int
	}{
		{1, 1, 2},
		{3, 3, 6},
		{2, 2, 4},
		{12, 2, 14},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.z {
			t.Errorf("incorrect sum of (%d+%d), got: %d, expected: %d.", table.x, table.y, total, table.z)
		}
	}
}