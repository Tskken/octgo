package octgo

// Action function for Entity
type Action func()

// Entities is a list of Entity
type Entities []*Entity

// Entity to be stored in the oct-tree
type Entity struct {
	ID uint64
	Bound
	Action
}
