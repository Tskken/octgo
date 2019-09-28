package octgo

type Action func()

type Entities []*Entity

type Entity struct {
	ID uint64
	Bound
	Action
}