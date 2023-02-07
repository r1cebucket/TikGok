package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/CyanAsterisk/TikGok/server/cmd/chat/dao"
	"github.com/CyanAsterisk/TikGok/server/cmd/chat/model"
	"github.com/CyanAsterisk/TikGok/server/shared/kitex_gen/chat"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

// Publisher implements an amqp publisher.
type Publisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewPublisher(conn *amqp.Connection, exchange string) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}

	if err = declareExchange(ch, exchange); err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}
	return &Publisher{
		ch:       ch,
		exchange: exchange,
	}, nil
}

// Publish publishes a message.
func (p *Publisher) Publish(_ context.Context, msg *chat.DouyinMessageActionRequest) error {
	body, err := sonic.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal: %v", err)
	}

	return p.ch.Publish(p.exchange, "", false, false, amqp.Publishing{
		Body: body,
	})
}

// Subscriber implements an amqp subscriber.
type Subscriber struct {
	conn     *amqp.Connection
	exchange string
}

// NewSubscriber creates an amqp subscriber.
func NewSubscriber(conn *amqp.Connection, exchange string) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	defer ch.Close()

	if err = declareExchange(ch, exchange); err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}

	return &Subscriber{
		conn:     conn,
		exchange: exchange,
	}, nil
}

// SubscribeRaw subscribes and returns a channel with raw amqp delivery.
func (s *Subscriber) SubscribeRaw(_ context.Context) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, func() {}, fmt.Errorf("cannot allocate channel: %v", err)
	}
	closeCh := func() {
		err := ch.Close()
		if err != nil {
			klog.Errorf("cannot close channel %s", err.Error())
		}
	}

	q, err := ch.QueueDeclare("", false, true, false, false, nil)
	if err != nil {
		return nil, closeCh, fmt.Errorf("cannot declare queue: %v", err)
	}

	cleanUp := func() {
		_, err := ch.QueueDelete(q.Name, false, false, false)
		if err != nil {
			klog.Errorf("cannot delete queue %s : %s", q.Name, err.Error())
		}
		closeCh()
	}

	err = ch.QueueBind(q.Name, "", s.exchange, false, nil)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot bind: %v", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot consume queue: %v", err)
	}
	return msgs, cleanUp, nil
}

// Subscribe subscribes and returns a channel with CarEntity data.
func (s *Subscriber) Subscribe(c context.Context) (chan *chat.DouyinMessageActionRequest, func(), error) {
	msgCh, cleanUp, err := s.SubscribeRaw(c)
	if err != nil {
		return nil, cleanUp, err
	}

	logicCh := make(chan *chat.DouyinMessageActionRequest)
	go func() {
		for msg := range msgCh {
			var carEn chat.DouyinMessageActionRequest
			err := sonic.Unmarshal(msg.Body, &carEn)
			if err != nil {
				klog.Errorf("cannot unmarshal %s", err.Error())
			}
			logicCh <- &carEn
		}
		close(logicCh)
	}()
	return logicCh, cleanUp, nil
}

func SubscribeRoutine(subscriber *Subscriber) {
	msgs, cleanUp, err := subscriber.Subscribe(context.Background())
	defer cleanUp()
	if err != nil {
		klog.Error("cannot subscribe", err)
	}
	for m := range msgs {
		err = dao.ChatAction(&model.Message{
			ToUserId:   m.ToUserId,
			FromUserId: m.UserId,
			Content:    m.Content,
			CreateDate: time.Now(),
		})
		if err != nil {
			klog.Error("cannot chat action", err)
		}
	}
}

func declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil)
}
