package nats

import (
	"errors"
	"reflect"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
)

type (
	ConsumeHandle func(m *nats.Msg) error

	// ConsumeHandler Consumer interface, used to define the methods required by the consumer
	ConsumeHandler interface {
		HandleMessage(m *nats.Msg) error
	}

	// ConsumerQueue Consumer queue, used to maintain the relationship between a consumer group and queue
	ConsumerQueue struct {
		QueueName         string         // queue name
		Subject           string         // Subscribe subject
		Consumer          ConsumeHandler // consumer object
		AckExplicitPolicy bool           // Ack explicit
	}

	// ConsumerManager Consumer manager for managing multiple consumer queues
	ConsumerManager struct {
		mutex    sync.RWMutex      // read-write lock
		conn     *nats.EncodedConn // nats-streaming connect
		queues   []ConsumerQueue   // consumer queue list
		options  []nats.Option     // Connection configuration items
		doneChan chan struct{}     // close channel
	}
)

// MustNewConsumerManager
func MustNewConsumerManager(cfg *NatsConfig, cq []*ConsumerQueue) queue.MessageQueue {
	nc, err := nats.Connect(cfg.Host, cfg.Options...)
	if err != nil {
		panic("failed to connect nats ,error " + err.Error())
	}
	nec, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		panic("failed to connect nats ,error " + err.Error())
	}

	cm := &ConsumerManager{
		conn:     nec,
		options:  cfg.Options,
		doneChan: make(chan struct{}),
	}
	if len(cq) == 0 {
		logx.Errorf("failed consumerQueue register to  stan, error: cq len is 0")
	}
	for _, item := range cq {
		err = cm.registerQueue(item)
		if err != nil {
			logx.Errorf("failed to register stan, error: %v", err)
		}
	}

	return cm
}

// Start consuming messages in the queue
func (cm *ConsumerManager) Start() {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	if len(cm.queues) == 0 {
		logx.Errorf("no consumer queues found")
	}
	for _, consumerQueue := range cm.queues {
		go cm.subscribe(consumerQueue)
	}
	logx.Info("start consuming messages in the queue")
	<-cm.doneChan
}

// Stop close connect
func (cm *ConsumerManager) Stop() {
	if cm.conn != nil {
		cm.conn.Close()
	}
}

// RegisterQueue Register a consumer queue
func (cm *ConsumerManager) registerQueue(queue *ConsumerQueue) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	if queue.QueueName == "" {
		return errors.New("queue name is required")
	}
	if queue.Subject == "" {
		return errors.New("subject is required")
	}
	if queue.Consumer == nil {
		return errors.New("consumer is required")
	}

	cm.queues = append(cm.queues, *queue)
	return nil
}

// subscribe news
func (cm *ConsumerManager) subscribe(queue ConsumerQueue) {
	sub, err := cm.conn.QueueSubscribe(queue.Subject, queue.QueueName, func(m *nats.Msg) {
		err := queue.Consumer.HandleMessage(m)
		if err != nil {
			logx.Errorf("error handling message: %v", err)
		} else {
			if queue.AckExplicitPolicy {
				err := m.Ack()
				if err != nil {
					logx.Errorf("error acking message: %v", err)
				}
			}
		}
	})
	if err != nil {
		logx.Errorf("error subscribing to queue %s: %v", queue.QueueName, err)
		return
	}
	<-cm.doneChan

	err = sub.Unsubscribe()
	if err != nil {
		logx.Errorf("error unsubscribing from queue %s: %v", queue.QueueName, err)
	}

	// delete consumer queue
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	for i, q := range cm.queues {
		if reflect.DeepEqual(q, queue) {
			cm.queues = append(cm.queues[:i], cm.queues[i+1:]...)
			break
		}
	}
}
