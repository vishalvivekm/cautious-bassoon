package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	log.Info("starting server..")
	api := NewAPI()
	http.HandleFunc("/api", api.Handler)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}

func (a *API) Handler(writer http.ResponseWriter, request *http.Request) {
	log.Info("in the handler")

	q := request.URL.Query().Get("q")
	data, cacheHit, err := a.getData(request.Context(), q)
	if err != nil {
		fmt.Printf("error calling data source: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
	resp := APIReponse{
		Cache: cacheHit,
		Data:  data,
	}
	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		fmt.Printf("error encoding the response: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
func (a *API) getData(ctx context.Context, q string) ([]NominatimResponse, bool, error) {
	// San Francisco , san%20francisco

	// is the query cached
	value, err := a.cache.Get(ctx, q).Result()
	if err == redis.Nil {
		// key does not exist
		// we want to call external data source
		escapedQ := url.PathEscape(q)
		address := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", escapedQ)
		resp, err := http.Get(address)
		if err != nil {
			return nil, false, err
		}
		data := make([]NominatimResponse, 0)
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, false, err
		}

		buf, err := json.Marshal(data)
		if err != nil {
			return nil, false, err
		}

		// set the value
		err = a.cache.Set(ctx, q, bytes.NewBuffer(buf).Bytes(), time.Second*15).Err()
		if err != nil {
			return nil, false, err
		}
		// return the response
		return data, false, nil
	} else if err != nil {
		// some error with redis
		fmt.Printf("error calling redis: %v\n", err)
		return nil, false, err
	} else {
		// cache hit
		// found the value
		data := make([]NominatimResponse, 0)
		// build response
		err := json.Unmarshal(bytes.NewBufferString(value).Bytes(), &data)
		if err != nil {
			return nil, false, err
		}
		// return response
		return data, true, nil
	}

}

type API struct {
	cache *redis.Client
}

func NewAPI() *API {
	redisAddress := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &API{
		cache: rdb,
	}
}

type APIReponse struct {
	Cache bool                `json:"cache"`
	Data  []NominatimResponse `json:"data"`
}
type NominatimResponse struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	PlaceRank   int      `json:"place_rank"`
	Importance  float64  `json:"importance"`
	Addresstype string   `json:"addresstype"`
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Boundingbox []string `json:"boundingbox"`
}

/*
curl localhost:8080/api -v


curl  curl https://nominatim.openstreetmap.org/search?q=\san%20francisco\&format\=json | jq

curl localhost:9090/api\?q=san%20francisco | jq

curl localhost:9090/api\?q=san%20francisco -v | jq
*/
/*
json to struct:
https://mholt.github.io/json-to-go/
*/
