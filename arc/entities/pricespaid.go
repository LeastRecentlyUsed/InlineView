package entities

// PriceStore is the main entity for prices paid records.  The entity unit of work for InlineView is a set of price paid
// records belonging to some arbitrary grouping defined in the business logic (usecases) and is always some type of
// immutable file-based medium.
type PriceStore struct {
	Store    []PriceRecord
	Index    map[string]string
	ModIndex bool
}

// PriceRecord has the required elements for storing price paid records (hash of the ID, the ID and the actual data record)
type PriceRecord struct {
	Hash       string
	Identifier string
	Data       string
}

// SyncIndexWithStore ensures that each price paid record has a matching index entry.
func (s *PriceStore) SyncIndexWithStore() {
	pre := len(s.Index)
	// for new stores that contain entries, the index could be uninitialised.
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
