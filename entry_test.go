package melems

import "testing"

func TestEntry(t *testing.T) {
	data := []struct {
		name   string
		length uint
	}{
		{"empty entry", 0},
		{"smallest entry", 1},
		{"entry of 6 values", 6},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			elements := make([]uint, d.length, d.length)
			entry := newEntry(elements)
			for i := uint(0); i < d.length; i++ {
				value, _ := entry.Get(i)
				if value != 0 {
					t.Errorf("inital value not zero at index %d", i)
				}
				_ = entry.Set(i, i+1)

				value, _ = entry.Get(i)
				if value != i+1 {
					t.Errorf("inital value not set at index %d", i)
				}
			}
		})
	}
}
