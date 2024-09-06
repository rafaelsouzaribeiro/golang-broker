package producer

import "github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"

func (p *Producer) GetErrorMessage() *payload.Message {

	return &p.message
}
