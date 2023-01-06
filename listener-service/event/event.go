package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// In RabbitMQ, there are four types of exchanges:
// 1. direct: The message is routed to the queues whose binding key exactly matches the routing key of the message.
// 2. fanout: The message is broadcasted to all the queues that are bound to the exchange.
// 3. topic: The message is routed to one or many queues based on matching between a message routing key and the pattern that was used to bind a queue to an exchange.
// 4. headers: The message is routed based on the header values, rather than the routing key.
func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable?
		false,        // auto-deleted?
		false,        // internal?
		false,        // no-wait?
		nil,          // arguements?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	// server will automatically generate a unique queue name for you
	return ch.QueueDeclare(
		"",    // name?
		false, // durable?
		false, // delete when unused?
		true,  // exclusive?
		false, // no-wait?
		nil,   // arguments?
	)
}
