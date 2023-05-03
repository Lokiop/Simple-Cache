package main

import "fmt"

//Random Size
const SIZE = 5

//Node for linked List
type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

//Queue for storing Head, Tail, Length
type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

//Cache
type Cache struct {
	Queue Queue
	Hash  Hash
}

//NewCache initializes the Cache
func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

//NewQueue initializes the Queue for the Cache
func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{
		Head: head,
		Tail: tail,
	}
}

//HashMap for the Faster Cache
type Hash map[string]*Node

//Chcek sees if the string is present in the Cache
func (c *Cache) Check(word string) {
	node := &Node{}

	if val, ok := c.Hash[word]; ok {
		node = c.Remove(val)
	} else {
		node.Val = word
	}

	c.Add(node)
}

//Remove removes the entry from the end of the cache
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right

	right.Left = left
	left.Right = right
	c.Queue.Length -= 1
	delete(c.Hash, n.Val)

	return n
}

//Add adds the node at the start of the cache
func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	head := c.Queue.Head
	tmp := head.Right

	head.Right = n
	tmp.Left = n
	n.Left = head
	n.Right = tmp

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
	c.Hash[n.Val] = n
}

//Displays the content of the cache
func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}

	fmt.Println("]")
}

func main() {
	fmt.Println("Start Cache")

	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "dragonfruit", "potato", "tree", "tomato", "tree", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
