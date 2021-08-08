package db

import (
	"context"
	"encoding/json"
	"time"
)

const duration = time.Minute * 5 // covert to env variable and use same for refresh token

func (store *CacheStore) SetPrinciple(ctx context.Context, key string, principle Priniple) error {

	data, err := json.Marshal(principle)
	if err != nil {
		return err
	}

	err = store.client.Set(ctx, key, string(data), duration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (store *CacheStore) GetPrinciple(ctx context.Context, key string) (*Priniple, error) {

	data, err := store.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var principle Priniple
	err = json.Unmarshal([]byte(data), &principle)
	if err != nil {
		return nil, err
	}

	return &principle, nil
}
