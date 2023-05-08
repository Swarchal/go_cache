# go cache

An in-memory JSON store with auto-expiring keys.

A cache is created with a lifetime (in seconds) associated with it, any object you store in
the cache will have a timestamp attached to it. When trying to retrieve the
object the `object_timestamp` + `cache_lifetime` is compared against the current time to check
if the object has expired.

It is up to the caller to remove any expired keys.

```go
Cache.Add(key string) []bytes
Cache.Delete(key string)
Cache.Get(key string) ([]bytes, error)
Cache.Exists(key string) bool
Cache.IsFresh(key string) bool
```
------------------

Example:

```go
import (
  "encoding/json"
  "fmt"
  "time"

  "github.com/swarchal/go_cache"
)


type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}


func main() {
  // create cache with a 60 second lifetime
  cache := MakeCache(60)

  // store JSON in cache
	bob := Person{"Bob", 60}
	b, err := json.Marshal(&bob)
	if err != nil {
		panic(err)
	}
	cache.Add("bob", b)

	out, err := cache.Get("bob")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}
```


