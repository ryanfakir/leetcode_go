package main

type LRUCache struct {
	m     map[int]*Node
	cap   int
	level int
	head  *Node
	tail  *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func (cache *LRUCache) addHead(n *Node) {
	n.prev = cache.head
	n.next = cache.head.next
	cache.head.next.prev = n
	cache.head.next = n
}

func (cache *LRUCache) deleteTrail(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func Constructor(capacity int) LRUCache {
	head := new(Node)
	tail := new(Node)
	head.next = tail
	tail.prev = head
	return LRUCache{make(map[int]*Node), capacity, 0, head, tail}
}

func (this *LRUCache) Get(key int) int {
	if this.m[key] != nil {
		res := this.m[key]
		this.deleteTrail(res)
		this.addHead(res)
		return res.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.m[key] != nil {
		node := this.m[key]
		node.value = value
		this.deleteTrail(this.m[key])
	} else {
		newNode := new(Node)
		newNode.key = key
		newNode.value = value
		this.m[key] = newNode
		if this.level >= this.cap {
			delete(this.m, this.tail.prev.key)
			this.deleteTrail(this.tail.prev)
		} else {
			this.level++
		}
	}
	this.addHead(this.m[key])
}
