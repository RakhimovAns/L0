package redisorderrepo

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestNewRepo(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repo := New(client)

	if repo == nil {
		t.Fatal("expected repo, got nil")
	}

	if repo.redisClient != client {
		t.Errorf("expected redisClient to be %v, got %v", client, repo.redisClient)
	}
}

func TestOrderCacheKey(t *testing.T) {
	id := "12345"
	got := orderCacheKey(id)
	want := orderCacheKeyPrefix + id

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
