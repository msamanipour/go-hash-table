package main

import "fmt"

const (
	// The size of the hash table array
	ArraySize = 7
)

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each solt of the table
type bucket struct {
	head *bucketNode
}

// bucketNode is a lined list node that hold the same value
type bucketNode struct {
	key  string
	next *bucketNode
}

//Insert a key to hash table
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search value in hash table and return true of false
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

// delete given value in hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//Bucket functions
// insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

//search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

func hash(key string) int {
	sum := 0
	for _, val := range key {
		sum += int(val)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return &result
}

func main() {
	hashTable := Init()
	list := []string{
		"Moahammad",
		"Milad",
		"Ali",
		"Farzad",
		"Mohsen",
		"Behnam",
		"Masoud",
	}
	for _, val := range list {
		hashTable.Insert(val)
	}
}
