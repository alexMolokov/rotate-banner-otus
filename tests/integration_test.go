package integrationtests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/alexMolokov/rotate-banner-otus/internal/config"
	rotatorstorage "github.com/alexMolokov/rotate-banner-otus/internal/storage/rotator"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/require"
)

var ConfigFile = os.Getenv("TEST_CONFIG_FILE")

func init() {
	if ConfigFile == "" {
		ConfigFile = "../configs/rotator.json"
	}
}

func TestStorage(t *testing.T) {
	cfg, err := config.NewRotatorConfig(ConfigFile)
	require.NoError(t, err, "can' t make config")

	storage := rotatorstorage.NewRotatorStorage(cfg.DB)
	err = storage.Connect()
	defer storage.Close()

	require.NoError(t, err, "can't connect to db")

	cfgQueue := cfg.Queue
	AmpqDSN := fmt.Sprintf("amqp://%s:%s@%s:%d", cfgQueue.User, cfgQueue.Password, cfgQueue.Host, cfgQueue.Port)

	_, err = amqp.Dial(AmpqDSN)
	require.NoError(t, err, "can't connect to rabbit mq")

	t.Run("get banner ", func(t *testing.T) {
		id := int64(1)

		banner, err := storage.GetBannerByID(context.Background(), id)
		require.NoError(t, err, "can't get banner by id")

		require.Equal(t, id, banner.ID)
	})

	t.Run("get slot", func(t *testing.T) {
		id := int64(1)

		slot, err := storage.GetSlotByID(context.Background(), id)
		require.NoError(t, err, "can't get slot")

		require.Equal(t, id, slot.ID)
	})

	t.Run("get social group", func(t *testing.T) {
		id := int64(1)

		sg, err := storage.GetSocialGroupByID(context.Background(), id)
		require.NoError(t, err, "can't get social group")
		require.Equal(t, id, sg.ID)
	})
}
