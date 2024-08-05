package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	rdb *redis.Client
}

func NewRepository(rdb *redis.Client) *Repository {
	return &Repository{rdb: rdb}
}

func (r *Repository) Create(ctx context.Context, key string, entity interface{}) error {
	json, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	return r.rdb.Set(ctx, key, json, 0).Err()
}

func (r *Repository) Get(ctx context.Context, key string, entity interface{}) error {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), entity)
}

func (r *Repository) Update(ctx context.Context, key string, entity interface{}) error {
	return r.Create(ctx, key, entity)
}

func (r *Repository) Delete(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}

func (r *Repository) GetAll(ctx context.Context, prefix string) ([]string, error) {
	keys, _, err := r.rdb.Scan(ctx, 0, prefix+"*", 0).Result()
	if err != nil {
		return nil, err
	}

	var results []string
	for _, key := range keys {
		val, err := r.rdb.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		results = append(results, val)
	}

	return results, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *User) error {
	return r.Create(ctx, fmt.Sprintf("user:%s", user.ID), user)
}

func (r *Repository) CreateEvent(ctx context.Context, event *Event) error {
	return r.Create(ctx, fmt.Sprintf("event:%s", event.ID), event)
}

func (r *Repository) CreateRole(ctx context.Context, role *Role) error {
	return r.Create(ctx, fmt.Sprintf("role:%s", role.ID), role)
}
