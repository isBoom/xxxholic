package task

import (
	"xxxholic/cache"
	"xxxholic/model"
)

func RestarRank(name string) error {
	for _, value := range model.VideoTypes {
		if err:=cache.RedisClient.Del(cache.GetRankName(name,value)).Err();err!=nil{
			return err
		}
	}
	return cache.RedisClient.Del(name).Err()
}