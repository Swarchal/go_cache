# go cache

An in-memory JSON store with auto-expiring keys.

A cache is created with a lifetime (in seconds) associated with it, any object you store in
the cache will have a timestamp attached to it. When trying to retrieve the
object the `object_timestamp` + `cache_lifetime` is compared against the current time to check
if the object has expired.

```go
type Cache struct {
    store    map[string]Entry
    lifetime time.Duration
}

type Entry struct {
    data      []byte
    timestamp time.Time
}


Cache.Add(key string, data []byte)
Cache.Delete(key string)
Cache.Get(key string) []byte
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

    r := cache.Get("bob")

    if r != nil {
        fmt.Println(string(r))
    }
}
```


