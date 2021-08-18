package cache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/rubberyconf/rubberyconf/internal/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestInRedisOptions(t *testing.T) {

	tests := getCommonScenarios()

	conf := config.GetConfiguration()
	if conf == nil {
		path, _ := os.Getwd()
		conf = config.NewConfiguration(filepath.Join(path, "../../config/local.yml"))
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer redisC.Terminate(ctx)

	endpoint, err := redisC.Endpoint(ctx, "")
	if err != nil {
		t.Error(err)
	}

	// hack config to work with testcontainers
	conf.Redis.Url = endpoint
	conf.Redis.Username = ""
	conf.Redis.Password = ""

	storage := NewDataStorageRedis()

	for _, tt := range tests {
		testname := fmt.Sprintf("key: %s, value: %s, duration: %d", tt.key, tt.value, tt.duration)
		t.Run(testname, func(t *testing.T) {
			result, err := storage.SetValue(ctx, tt.key, tt.value, tt.duration)
			if !result || err != nil {
				t.Errorf("Imposible to store this object")
			}
			time.Sleep(500 * time.Millisecond)
			_, found, err := storage.GetValue(ctx, tt.key)
			if err != nil && found == tt.found {
				t.Errorf("got '%t', want '%t'", found, tt.found)
			}
		})
	}
}
