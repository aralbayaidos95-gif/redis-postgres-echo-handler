package storage

import (
	"context"
	"encoding/json"
	"study/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RDB *redis.Client
}

func NewRedis(connStrRDB string) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr: connStrRDB,
	})

	return &Redis{RDB: rdb}
}

func (r *Redis) CreateUser(user models.User) error {
	data, _ := json.Marshal(user)

	return r.RDB.Set(context.Background(), user.Name, data, time.Minute*15).Err()
}

func (r Redis) GetUser(name string) (models.User, error) {
	var user models.User
	data,err:=r.RDB.Get(context.Background(),name).Result()

	if err!=nil{
		return user,err
	}

	if err=json.Unmarshal([]byte(data),&user);err!=nil{
		return user,err
	}

	return user,nil
}
