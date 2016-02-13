package services

import (
	. "app/common"
	. "github.com/andboson/configlog"
	"gopkg.in/redis.v2"
)

var RClient *redis.Client
var RClientBasket3 *redis.Client

func init() {
	RClient = RedisClient()
}

func RedisClient() *redis.Client {
	if RClient != nil {
		return RClient
	}

	debug, _ := AppConfig.Bool("debug")

	redisHost, _ := AppConfig.String("redis.redis_host")
	redisPort, _ := AppConfig.String("redis.redis_port")
	redisPass, _ := AppConfig.String("redis.redis_pass")

	RClient = redis.NewTCPClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass, // no password set
		DB:       0,         // use default DB
	})

	pong, err := RClient.Ping().Result()
	if err != nil {
		Log.Fatalf("Unable connect to Redis Server! %v %v", pong, err)
	}

	if debug {
		Log.Printf("new Redis client started")
	}

	return RClient
}
