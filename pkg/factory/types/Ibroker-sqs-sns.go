package types

import "github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"

type SQSSNS interface {
	Receive(MessageChan chan<- payload.SNSSQSMessage)
	Send()
}
