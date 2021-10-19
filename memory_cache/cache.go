package memory_cache

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	once       sync.Once
	InnerCache LRUCache
)

func ResourceInit() {
	once.Do(func() {
		InnerCache = InnerCacheConstructor(2000) // 初始化内存缓存
	})
}

// GenKey 生成一个缓存key
func GenKey(data []string) string {
	return strings.Join(data, ":")
}

/* 内存缓存 */

type LRUCache struct {
	size       int
	Capacity   int
	Cache      map[string]*DLinkedNode
	Head, Tail *DLinkedNode
}

type DLinkedNode struct {
	key, value string
	expire     time.Time
	Prev, Next *DLinkedNode
}

func InitDLinkedNode(key, value string, duration int) *DLinkedNode {
	ad, err := time.ParseDuration(genRandomStrNum(duration))
	if err != nil {
		ad, _ = time.ParseDuration("300s")
	}
	expireTime := time.Now().Add(ad)
	return &DLinkedNode{
		key:    key,
		value:  value,
		expire: expireTime,
	}
}

func (cls *LRUCache) Get(key string) string {
	if _, ok := cls.Cache[key]; !ok {
		return ""
	}
	node := cls.Cache[key]

	// 是否过期
	if time.Now().After(node.expire) {
		cls.removeNode(node)
		delete(cls.Cache, node.key)
		cls.size--
		return ""
	}
	cls.moveToHead(node)
	return node.value
}

func (cls *LRUCache) Put(key string, value string, duration int) {
	if _, ok := cls.Cache[key]; !ok {
		node := InitDLinkedNode(key, value, duration)
		cls.Cache[key] = node
		cls.addToHead(node)
		cls.size++
		if cls.size > cls.Capacity {
			removed := cls.removeTail()
			delete(cls.Cache, removed.key)
			cls.size--
		}
	} else {
		node := cls.Cache[key]
		node.value = value
		cls.moveToHead(node)
	}
}

func (cls *LRUCache) addToHead(node *DLinkedNode) {
	node.Prev = cls.Head
	node.Next = cls.Head.Next
	cls.Head.Next.Prev = node
	cls.Head.Next = node
}

func (cls *LRUCache) removeNode(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (cls *LRUCache) moveToHead(node *DLinkedNode) {
	cls.removeNode(node)
	cls.addToHead(node)
}

func (cls *LRUCache) removeTail() *DLinkedNode {
	node := cls.Tail.Prev
	cls.removeNode(node)
	return node
}

// 生成一个5～8m的随机过期时间
func genRandomStrNum(duration int) string {
	var seconds string
	if duration == 0 {
		// 生成一个5~8分钟的随机过期时间
		n, err := rand.Int(rand.Reader, big.NewInt(8-5))
		if err != nil {
			n = big.NewInt(0)
		}
		nInt := n.Int64()
		seconds = strconv.Itoa(int(60*(nInt+5))) + "s"
	} else {
		seconds = strconv.Itoa(duration) + "s"
	}
	return seconds
}

// InnerCacheConstructor 初始化内存缓存
func InnerCacheConstructor(capacity int) LRUCache {
	l := LRUCache{
		Cache:    map[string]*DLinkedNode{},
		Head:     InitDLinkedNode("head", "", 300),
		Tail:     InitDLinkedNode("tail", "", 300),
		Capacity: capacity,
	}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head
	return l
}
