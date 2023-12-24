package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second;
	cases := []struct {
		key string
		data []byte
	}{
		{
			key: "https://example.com",
			data: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			data: []byte("testdata2"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func (t *testing.T) {
			cache := NewCache(interval)
			cache.Set(c.key, c.data)
			data, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			}
			if string(data) != string(c.data) {
				t.Errorf("Expected %v, got %v", string(c.data), string(data))
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)
	cache.Set("https://example.com", []byte("testdata"))
	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("Expected to find key")
		return
	}
	time.Sleep(waitTime)
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("Expected not to find key")
		return
	}
}