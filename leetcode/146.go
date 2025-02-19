package main

type DLinkedNode struct {
	key   int
	value int
	pre   *DLinkedNode
	next  *DLinkedNode
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkedNode
	head, tail     *DLinkedNode
}

func initDLinkedNode(k, v int) *DLinkedNode {
	return &DLinkedNode{
		key:   k,
		value: v,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode, capacity),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}

	//create dummy head and tail
	l.head.next = l.tail
	l.tail.pre = l.head

	return l
}

func (this *LRUCache) Get(key int) int {
	//check from map
	if node, ok := this.cache[key]; ok {
		//exist -> move node to head
		this.moveToHead(node)
		return node.value
	}
	//non-exist -> return -1
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//check from map
	if node, ok := this.cache[key]; ok {
		//exist -> overwrite the value, move node to head
		node.value = value
		this.moveToHead(node)

	} else {
		//non-exist -> create a new node , add this node before head
		node = initDLinkedNode(key, value)
		//add node before head (注意dummy head永远为head)
		this.addToHead(node)
		//update cache
		this.cache[key] = node
		//if over capacity, delete tail
		if this.size >= this.capacity {
			//*************remove from map***************
			node = this.tail.pre
			this.removeNode(node)
			delete(this.cache, node.key)
		} else {
			this.size++
		}
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	node.next, node.pre = nil, nil
}
