package _46__LRU_Cache

type LRUCache struct {
	capacity, size int
	head, tail     *node
	m              map[int]*node
}

type node struct {
	key, val   int
	prev, next *node
}

func Constructor(capacity int) LRUCache {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	l := LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		m:        make(map[int]*node),
	}
	return l
}

func (this *LRUCache) removeNode(n *node) {
	prev, next := n.prev, n.next
	prev.next = next
	next.prev = prev

	this.size--
	delete(this.m, n.key)
}

func (this *LRUCache) insertNode(n *node) {
	prev := this.tail.prev
	prev.next = n
	n.prev = prev
	n.next = this.tail
	this.tail.prev = n

	this.size++
	this.m[n.key] = n
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.removeNode(n)
		this.insertNode(n)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.m[key]; ok {
		this.removeNode(n)
		this.insertNode(&node{key: key, val: value})
		return
	}
	this.insertNode(&node{key: key, val: value})
	if this.size > this.capacity {
		this.removeNode(this.head.next)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
