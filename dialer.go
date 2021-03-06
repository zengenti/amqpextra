package amqpextra

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Dialer interface {
	Dial() (*amqp.Connection, error)
}

func NewDialer(url string, config amqp.Config) Dialer {
	return DialerFunc(func() (*amqp.Connection, error) {
		return amqp.DialConfig(url, config)
	})
}

type DialerFunc func() (*amqp.Connection, error)

func (f DialerFunc) Dial() (*amqp.Connection, error) {
	return f()
}

func NewMultiHostDialer(
	username string,
	password string,
	hosts []string,
	port int,
	vhost string,
	config amqp.Config,
) Dialer {
	i := 0
	l := len(hosts)

	return DialerFunc(func() (*amqp.Connection, error) {
		url := fmt.Sprintf(
			"amqp://%s:%s@%s:%d/%s",
			username,
			password,
			hosts[i],
			port,
			vhost,
		)

		i = (i + 1) % l

		return amqp.DialConfig(url, config)
	})
}
