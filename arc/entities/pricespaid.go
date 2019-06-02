package entities

import "fmt"

// PriceRecord has the required elements for storing price paid records (hash of the ID, the ID and the actual data record)
type PriceRecord struct {
	Hash       string
	Identifier string
	Data       string
}

// PriceStore is the main entity for prices paid records.  The entity unit of work for InlineView is a set of price paid
// records belonging to some arbitrary grouping defined in the business logic (usecases) and is always some type of
// immutable file-based medium.
type PriceStore struct {
	Store    []PriceRecord
	Index    map[string]string
	ModIndex bool
}

// SyncIndexWithStore ensures that each price paid record has a matching index entry.
func (s *PriceStore) SyncIndexWithStore() {
	pre := len(s.Index)
	// for new stores that contain price record entries, the index could be uninitialised.
	if s.Index == nil && len(s.Store) > 0 {
		s.Index = map[string]string{}
	}

	for _, val := range s.Store {
		if _, found := s.Index[val.Hash]; !found {
			s.Index[val.Hash] = val.Identifier
		}
	}
	if len(s.Index) > pre {
		s.ModIndex = true
	} else {
		s.ModIndex = false
	}
}

// IndexActions defines the functional contract for the data persistence of an index
type IndexActions interface {
	ReadIndex(string) (map[string]string, error)
	WriteIndex(string, map[string]string) error
}

// ManageIndex holds the index values and uses IndexActions to perform the retrieval and saving
type ManageIndex struct {
	IndexData  map[string]string
	IndexStore IndexActions
}

// ReadIndex uses the read method on the interface as a proxy for the infrastructure read operation.
func (ia *ManageIndex) ReadIndex(storeName string) {
	var err error
	ia.IndexData, err = ia.IndexStore.ReadIndex(storeName)
	if err != nil {
		fmt.Println(err)
	}
}

// WriteIndex uses the write method on the interface as a proxy for the infrastructure write operation
func (ia *ManageIndex) WriteIndex(storeName string) {
	ia.IndexStore.WriteIndex(storeName, ia.IndexData)
}

// NewIndex is the constructor that returns a ManageIndex object.
func NewIndex(ia IndexActions) *ManageIndex {
	return &ManageIndex{
		IndexData:  map[string]string{},
		IndexStore: ia,
	}
}
