package repository

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/redis/go-redis/v9"
)

// Redisの操作を抽象化したインターフェース
type IRedisRepository interface {
	SaveSession(c context.Context, sessionKey string, redisValuealue string) error
	GetValue(c context.Context, sessionKey string) (string, error)
	GenerateSessionKey() (string, error)
	DeleteSession(c context.Context, sessionKey string) error
}

// Redisの具体的な実装
type redisRepository struct {
	client *redis.Client
}

// 新しいRedisリポジトリを作成
func NewRedisRepository(client *redis.Client) IRedisRepository {
	return &redisRepository{client}
}

func (r *redisRepository) SaveSession(c context.Context, sessionKey string, redisValue string) error {

	// RedisID: AccountID
	// SessionKey: ランダム生成
	if err := r.client.Set(c, sessionKey, redisValue, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisRepository) GetValue(c context.Context, sessionKey string) (string, error) {
	val, err := r.client.Get(c, sessionKey).Result()

	if err == redis.Nil {
		return "", errors.New("SessionKeyが登録されていません。")
	}
	if err != nil {
		return "", err
	}

	return val, nil
}
func (r *redisRepository) GenerateSessionKey() (string, error) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
func (r *redisRepository) DeleteSession(c context.Context, sessionKey string) error {
	if err := r.client.Del(c, sessionKey).Err(); err != nil {
		return err
	}
	return nil
}
