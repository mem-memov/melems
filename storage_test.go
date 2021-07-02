package melems

import "testing"

func TestStorage(t *testing.T) {
	data := []struct {
		name string
		entryLength uint
		increment uint
	} {
		{"happy path", 6, 100},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			storage, _ := NewStorage(d.entryLength, d.increment)
			entryWritten := storage.Create()
			_ = entryWritten.Set(0, 1)
			_ = entryWritten.Set(1, 2)
			_ = entryWritten.Set(2, 3)
			_ = entryWritten.Set(3, 4)
			_ = entryWritten.Set(4, 5)
			_ = entryWritten.Set(5, 6)
			_ = storage.Write(0, entryWritten)
			entryRead := storage.Create()
			_ = storage.Read(0, entryRead)
			if entryRead != entryWritten {
				t.Errorf("read %#v what is not written %#v", entryRead, entryWritten)
			}
		})
	}
}
