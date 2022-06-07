package main

/*
$ curl http://localhost:9999/_geecache/scores/Tom
630

$ curl http://localhost:9999/_geecache/scores/kkk
kkk not exist
*/

import (
	"fmt"
	"log"
	"net/http"
	"reecache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	reecache.NewGroup("scores", 2<<10, reecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := reecache.NewHTTPPool(addr)
	log.Println("rreecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
