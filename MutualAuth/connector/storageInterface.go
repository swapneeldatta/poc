package connector

import (
	"log"

	"dattus.com/errorhandler"
	r "github.com/go-redis/redis"
)

const passwordForRedisInstance = ""
const dbNumberToUse = 0

const pingURL = "post_ping"
const dataURL = "post_data"
const geURL = "post_event"

// InitializeClient attempts to build a connection to the local redis instance
func InitializeClient() *r.Client {
	uri := "localhost:6379"
	return InitializeClientWithURI(&uri)
}

// InitializeClientWithURI attempts to build a connection to the local redis instance
func InitializeClientWithURI(uri *string) *r.Client {
	client := r.NewClient(&r.Options{
		Addr:     *uri,
		Password: passwordForRedisInstance,
		DB:       dbNumberToUse,
	})
	return client
}

// GetPortalURL returns the portal URL from cache
func GetPortalURL(r *r.Client) *string {
	result, e := r.HGet("url", dataURL).Result()
	errorhandler.HandleError("fetching portal URL", &e, false)
	return &result
}

// GetPingURL returns the ping URL from cache
func GetPingURL(r *r.Client) *string {
	result, e := r.HGet("url", pingURL).Result()
	errorhandler.HandleError("fetching ping URL", &e, false)
	return &result
}

// GetEventURL returns the ping URL from cache
func GetEventURL(r *r.Client) *string {
	result, e := r.HGet("url", geURL).Result()
	errorhandler.HandleError("fetching general event URL", &e, false)
	return &result
}

// GetURLs returns all URLs from cache
func GetURLs(r *r.Client) (*[]interface{}, *error) {
	result, e := r.HMGet("url", pingURL, dataURL, geURL).Result()
	errorhandler.HandleError("fetching URLs", &e, false)
	return &result, &e
}

// GetChunkSize returns the chunkSize from cache
func GetChunkSize(r *r.Client) (uint64, *error) {
	result, e := r.Get("chunk_size").Uint64()
	errorhandler.HandleError("fetching chunk size", &e, false)
	return result, &e
}

// Ping attempts a ping to the redis server
func Ping(client *r.Client) {
	pong, e := client.Ping().Result()
	errorhandler.HandleError("pinging redis server", &e, true)
	log.Println("Redis replies :", string(pong))
}

// GetAPIKeys fetches the api keys of all hubs configured
func GetAPIKeys(client *r.Client) ([]string, *error) {
	keys, e := client.HKeys("hub").Result()
	errorhandler.HandleError("pinging redis server", &e, false)
	return keys, &e
}
