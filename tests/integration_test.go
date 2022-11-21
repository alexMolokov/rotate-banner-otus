package integrationtests

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/alexMolokov/rotate-banner-otus/internal/config"
	rotatorstorage "github.com/alexMolokov/rotate-banner-otus/internal/storage/rotator"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/require"
)

type BannerSlot struct {
	BannerID int64 `db:"banner_id"`
	SlotID   int64 `db:"slot_id"`
}

type StatSlot struct {
	Display int64 `db:"display"`
	Click   int64 `db:"click"`
}

var ConfigFile = os.Getenv("TEST_CONFIG_FILE")

func init() {
	if ConfigFile == "" {
		ConfigFile = "../configs/rotator.json"
	}
}

func TestStorage(t *testing.T) {
	cfg, err := config.NewRotatorConfig(ConfigFile)
	require.NoError(t, err, "can' t make config")

	cfgDB := cfg.DB

	storage := rotatorstorage.NewRotatorStorage(cfgDB)
	err = storage.Connect()
	defer storage.Close()

	require.NoError(t, err, "can't connect to db")

	cfgQueue := cfg.Queue
	AmpqDSN := fmt.Sprintf("amqp://%s:%s@%s:%d", cfgQueue.User, cfgQueue.Password, cfgQueue.Host, cfgQueue.Port)

	_, err = amqp.Dial(AmpqDSN)
	require.NoError(t, err, "can't connect to rabbit mq")

	PostgresDSN := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfgDB.Host, cfgDB.Port, cfgDB.User, cfgDB.Password, cfgDB.Name)

	db, err := sqlx.ConnectContext(context.Background(), "postgres", PostgresDSN)
	require.NoError(t, err, "can't connect to db")

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

	t.Run("banner to slot link", func(t *testing.T) {
		bannerID := int64(1)
		slotID := int64(1)
		sgID := int64(1)

		err := storage.AddBannerToSlot(context.Background(), bannerID, slotID)
		require.NoError(t, err, "can't add banner to slot")

		var bannerSlot BannerSlot
		err = db.Get(&bannerSlot,
			"SELECT banner_id, slot_id FROM banner_to_slot WHERE banner_id=$1 AND slot_id=$2",
			bannerID, slotID)
		require.NoError(t, err, "can't get banner to slot link")

		err = storage.CountTransition(context.Background(), bannerID, slotID, sgID)
		require.NoError(t, err, "can't count transition")

		err = storage.CountDisplay(context.Background(), bannerID, slotID, sgID)
		require.NoError(t, err, "can't count display")

		var statSlot StatSlot
		err = db.Get(&statSlot,
			"SELECT display, click FROM stat WHERE banner_id=$1 AND slot_id=$2 AND social_group_id=$3",
			bannerID, slotID, sgID)
		require.NoError(t, err, "can't get stat")
		require.Greater(t, statSlot.Display, int64(1), "display must be more 1")
		require.Greater(t, statSlot.Click, int64(0), "click must be more 0")

		require.Equal(t, bannerID, bannerSlot.BannerID)
		require.Equal(t, slotID, bannerSlot.SlotID)

		err = storage.RemoveBannerFromSlot(context.Background(), bannerID, slotID)
		require.NoError(t, err, "can't remove banner to slot link")
		err = db.Get(&bannerSlot,
			"SELECT banner_id, slot_id FROM banner_to_slot WHERE banner_id=$1 AND slot_id=$2",
			bannerID, slotID)
		require.ErrorIs(t, err, sql.ErrNoRows)
	})
}
