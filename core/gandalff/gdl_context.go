package gandalff

import (
	"fmt"
	"sync"
)

type Context struct {
	// StringPool is a pool of strings that are used by the series.
	// This is used to reduce the number of allocations and to allow for fast comparisons.
	stringPool *StringPool

	threadsNumber int
	nullString    string
	timeFormat    string
}

func NewContext() *Context {
	return &Context{
		stringPool:    NewStringPool(),
		threadsNumber: THREADS_NUMBER,
		nullString:    NULL_STRING,
		timeFormat:    "2006-01-02 15:04:05",
	}
}

func (ctx *Context) GetThreadsNumber() int {
	return ctx.threadsNumber
}

func (ctx *Context) SetThreadsNumber(n int) {
	ctx.threadsNumber = n
}

func (ctx *Context) GetNullString() string {
	return ctx.nullString
}

func (ctx *Context) SetNullString(s string) {
	ctx.nullString = s
}

func (ctx *Context) GetTimeFormat() string {
	return ctx.timeFormat
}

func (ctx *Context) SetTimeFormat(s string) {
	ctx.timeFormat = s
}

type StringPool struct {
	sync.RWMutex
	pool          map[string]*string
	nullStringPtr *string
}

func NewStringPool() *StringPool {
	pool := &StringPool{pool: make(map[string]*string)}
	pool.nullStringPtr = pool.Put(NULL_STRING)

	return pool
}

// Get returns the address of the string if it exists in the pool, otherwise nil.
func (sp *StringPool) Get(s string) *string {
	if entry, ok := sp.pool[s]; ok {
		return entry
	}
	return nil
}

// Put returns the address of the string if it exists in the pool, otherwise it adds it to the pool and returns its address.
func (sp *StringPool) Put(s string) *string {
	if entry, ok := sp.pool[s]; ok {
		return entry
	}

	// Create a new string and add it to the pool
	addr := &s
	sp.pool[s] = addr
	return addr
}

// PutSync returns the address of the string if it exists in the pool, otherwise it adds it to the pool and returns its address.
// This version is thread-safe.
func (sp *StringPool) PutSync(s string) *string {
	sp.RLock()
	entry, ok := sp.pool[s]
	sp.RUnlock()
	if ok {
		return entry
	}

	sp.Lock()
	defer sp.Unlock()
	if entry, ok := sp.pool[s]; ok {
		// Someone else inserted the string while we were waiting
		return entry
	}

	// Create a new string and add it to the pool
	sp.pool[s] = &s
	return sp.pool[s]
}

func (sp *StringPool) Len() int {
	return len(sp.pool)
}

func (sp *StringPool) ToString() string {
	out := "StringPool["
	for _, v := range sp.pool {
		out += *v + ", "
	}
	out = out[:len(out)-2] + "]"
	return out
}

func (sp *StringPool) debugPrint() {
	for k, v := range sp.pool {
		fmt.Printf("%s: %p\n", k, v)
	}
}
