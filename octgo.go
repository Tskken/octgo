// Package octgo will be a oct-tree implementation in Go aimed at game collision detection.
//
// Currently under development and not implemented currently.
package octgo

// OctGo - Core oct-tree structure for OctGo
type OctGo struct {
	*node

	maxDepth uint16
}

type nodes []*nodes

type node struct {
	parent   *node
	bound    Bound
	entities Entities
	children nodes
	depth    uint16
}
