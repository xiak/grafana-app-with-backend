package keywords

import (
	"container/list"
	"strings"
	"sync"
)

type Keywords interface {
	Add(keyword string)
	AddRaw(keywords []string)
	Remove(keyword string)
	Find(input string) []string
	Length() int
}

type KeywordList struct {
	keywords *list.List
	mutex    sync.Mutex
	opts     options
}

// KeywordListOption KeywordList options
type KeywordListOption func(o *options)

// options put here
type options struct {
	// keywords retrieve number
	retrieve int64
}

// Retrieve keywords retrieve numbers
func Retrieve(num int64) KeywordListOption {
	return func(k *options) {
		k.retrieve = num
	}
}

// NewKeywordList init
func NewKeywordList(opts ...KeywordListOption) Keywords {
	op := options{
		retrieve: 5,
	}
	for _, o := range opts {
		o(&op)
	}
	return &KeywordList{
		keywords: list.New(),
		opts:     op,
	}
}

// Add a string to keywords list
func (kl *KeywordList) Add(keyword string) {
	kl.mutex.Lock()
	defer kl.mutex.Unlock()
	for e := kl.keywords.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == keyword {
			return
		}
	}
	kl.keywords.PushBack(keyword)
}

// Add some string to keywords list
func (kl *KeywordList) AddRaw(keywords []string) {
	for _, keyword := range keywords {
		kl.Add(keyword)
	}
}

// Find some string from keywords list
func (kl *KeywordList) Find(input string) []string {
	kl.mutex.Lock()
	defer kl.mutex.Unlock()
	var recommendations []string
	var cnt int64 = 0
	for e := kl.keywords.Front(); e != nil; e = e.Next() {
		keyword := e.Value.(string)
		if cnt < kl.opts.retrieve && strings.Contains(keyword, input) {
			recommendations = append(recommendations, keyword)
			cnt++
		}
	}
	return recommendations
}

// Remove a string from keywords list
func (kl *KeywordList) Remove(keyword string) {
	kl.mutex.Lock()
	defer kl.mutex.Unlock()
	for e := kl.keywords.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == keyword {
			kl.keywords.Remove(e)
			break
		}
	}
}

// Length get keyword list length
func (kl *KeywordList) Length() int {
	kl.mutex.Lock()
	defer kl.mutex.Unlock()
	return kl.keywords.Len()
}
