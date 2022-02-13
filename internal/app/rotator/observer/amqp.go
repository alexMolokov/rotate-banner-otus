package observer

import (
	"encoding/json"
	"fmt"

	"github.com/alexMolokov/rotate-banner-otus/internal/app/rotator"
	"github.com/alexMolokov/rotate-banner-otus/internal/config"
	"github.com/streadway/amqp"
)

type AMQPMessage struct {
	Type          string `json:"type"`
	SlotID        int64  `json:"slotId"`
	BannerID      int64  `json:"bannerId"`
	SocialGroupID int64  `json:"socialGroupId"`
	Date          string `json:"date"`
}

type AMQPObserver struct {
	cfg     config.QueueConf
	logger  approtator.Logger
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewAMQPObserver(cfg config.QueueConf, logger approtator.Logger) *AMQPObserver {
	return &AMQPObserver{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *AMQPObserver) Init() error {
	conn, err := amqp.Dial(a.getURI())
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	a.conn = conn

	channel, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}
	a.channel = channel

	if err := channel.ExchangeDeclare(
		a.cfg.ExchangeName, // name
		a.cfg.ExchangeType, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // noWait
		nil,                // arguments
	); err != nil {
		return fmt.Errorf("exchange declare: %w", err)
	}
	return nil
}

func (a *AMQPObserver) Close() {
	if a.conn == nil {
		return
	}

	if err := a.conn.Close(); err != nil {
		a.logger.Error("can't close amqp connection")
	}
}

func (a *AMQPObserver) Handle(em approtator.EventMessage) {
	if a.channel == nil {
		return
	}

	message := AMQPMessage{
		Type:          em.Type,
		SlotID:        em.SlotID,
		BannerID:      em.BannerID,
		SocialGroupID: em.SgID,
		Date:          em.Date,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		a.logger.Error("can't marshall message, %w", err)
	}

	if err := a.channel.Publish(
		a.cfg.ExchangeName, // publish to an exchange
		"banner_statistic", // routing to 0 or more queues
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            bytes,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		a.logger.Error("exchange publish: %w", err)
	}
}

func (a *AMQPObserver) getURI() string {
	return "amqp://" + a.cfg.User + ":" + a.cfg.Password + "@" + a.cfg.Host + ":" + string(rune(a.cfg.Port))
}
