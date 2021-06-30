package melems

import (
	"github.com/mem-memov/mbook"
)

type Storage struct {
	entryLength uint
	book *mbook.Book
}

func NewStorage(entryLength uint, increment uint) (*Storage, error) {
	book, err := mbook.NewBook(increment * entryLength)
	if err != nil {
		return nil, err
	}

	s := &Storage{
		entryLength: entryLength,
		book: book,
	}

	return s, nil
}

func (s *Storage) Create() *Entry {
	elements := make([]uint, s.entryLength, s.entryLength)
	return newEntry(elements)
}

func (s *Storage) Read(position uint, entry *Entry) error {
	return nil
	//entry.read(position, )
}
