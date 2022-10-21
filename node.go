package node

import (
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}
