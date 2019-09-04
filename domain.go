package rabbit

import (
	"github.com/streadway/amqp"
)

func NewConnect(url string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(url)
	return connection, err
}

func NewChannel(con *amqp.Connection) (*amqp.Channel, error) {
	ch, err := con.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func CreateQueue(ch *amqp.Channel, q Queue) (amqp.Queue, error) {
	return ch.QueueDeclare(
		q.Name,
		q.Durable,
		q.AutoDelete,
		q.Durable,
		q.NoWait,
		q.Args,
	)
}

func CreateAmqpChannel(ch *amqp.Channel,consumer Consumer) (<-chan amqp.Delivery,error) {
	return ch.Consume(
		consumer.QueueName,
		consumer.Name,
		consumer.AutoAck,
		consumer.Exclusive,
		consumer.NoLocal,
		consumer.NoWait,
		consumer.Args,
	)}


func CreateExchange(ch *amqp.Channel, e Exchange) error {
	return ch.ExchangeDeclare(
		e.Name,
		e.Kind,
		e.Durable,
		e.AutoDeleted,
		e.Internal,
		e.NoWait,
		e.Args,
	)}


func CreateBind(ch *amqp.Channel ,b Bind) error {
	return ch.QueueBind(
		b.QueueName,
		b.RoutingKey,
		b.Exchange,
		b.NoWait,
		b.Args,
	)}










type Exchange struct {
	Name        string
	Kind        string
	Durable     bool
	AutoDeleted bool
	Internal    bool
	NoWait      bool
	Args        amqp.Table
}

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type Consumer struct {
	QueueName      string
	Name  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

type Bind struct {
	QueueName      string
	RoutingKey  string
	Exchange   string
	NoWait    bool
	Args      amqp.Table
}

type Message struct {
	ExchangeName string
	RoutingKey string
	Mandatory bool
	Immediate bool
	Publish amqp.Publishing
}

func (msg *Message) Send(ch *amqp.Channel) error {
	return ch.Publish(
		msg.ExchangeName,
		msg.RoutingKey,
		msg.Mandatory,
		msg.Immediate,
		msg.Publish,
		)}