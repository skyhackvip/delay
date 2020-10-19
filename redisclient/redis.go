package redisclient

import (
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func ZAdd(key string, score int, data string) error {
	z := redis.Z{Score: float64(score), Member: data}
	_, err := client.ZAdd(key, z).Result()
	return err
}

func ZRangeFirst(key string) ([]interface{}, error) {
	val, err := client.ZRangeWithScores(key, 0, 0).Result()
	if err != nil {
		return nil, err
	}
	if len(val) == 1 {
		score := val[0].Score
		member := val[0].Member
		data := []interface{}{score, member}
		return data, nil
	}
	return nil, nil
}

func ZRem(key string, member string) error {
	_, err := client.ZRem(key, member).Result()
	return err
}

func Set(key string, data string) error {
	_, err := client.Set(key, data, -1).Result()
	return err
}

func Get(key string) (string, error) {
	val, err := client.Get(key).Result()
	return val, err
}

func Del(key string) error {
	_, err := client.Del(key).Result()
	return err
}
