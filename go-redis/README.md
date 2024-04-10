To run locally: 
```
docker compose up --build
```

from another terminal session:
- To see full response:
```
curl localhost:9090/api\?q=san%20francisco -v | jq
```
- To avoid full response and just check cacheHit:
```
curl localhost:9090/api\?q=san%20francisco -v | jq | grep cache
```
> each 15 seconds the cached value expires, so running this a couple of times, can see the working of this program.
<details><summary>curl reqs and redis-details</summary>
<details><summary>curl reqs</summary>

```bash
curl localhost:9090/api -v
curl  curl https://nominatim.openstreetmap.org/search?q=\san%20francisco\&format\=json | jq

curl localhost:9090/api\?q=san%20francisco | jq

curl localhost:9090/api\?q=san%20francisco -v | jq
```

</details>

<details><other refs</summary>
json to struct:
https://mholt.github.io/json-to-go/

</details>

- https://github.com/redis/go-redis

```go
// set the value

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }
//  https://github.com/vishalvivekm/gopher/blob/c0c7dd0b923de5caa4fbbeec71418696bb028523/go-redis/main.go#L68-L72
/*
    err = a.cache.Set(ctx, q, bytes.NewBuffer(buf).Bytes(), time.Second*15).Err()
		if err != nil {
			return nil, false, err
		}*/
```

```go
    // get the value of key, if it exists
    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist

// https://github.com/vishalvivekm/gopher/blob/c0c7dd0b923de5caa4fbbeec71418696bb028523/go-redis/main.go#L47
/*
value, err := a.cache.Get(ctx, q).Result()
....
...
*/
```

</details>
