package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Millisecond * 10)
	if c.cache == nil {
		t.Errorf("Cache was not created")
	}
}

func TestAddToCache(t *testing.T) {
	c := NewCache(time.Millisecond * 10)
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
		inputKey: "location",
		inputVal: []byte("12345"),
		},
		{
		inputKey: "key2",
		inputVal: []byte("value2"),
		},
		{
		inputKey: "key3",
		inputVal: []byte("value3"),
		},
	}

	for _, cas := range cases {
		c.Add(cas.inputKey, cas.inputVal)
		actual, ok := c.Get(cas.inputKey)

		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s doesn't match: %s", string(actual), cas.inputVal)
			continue
		}
	}
}

func TestReap(t *testing.T){
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))

	time.Sleep(interval + time.Millisecond)

	_, ok:= cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should have been reaped", keyOne)
	}

}