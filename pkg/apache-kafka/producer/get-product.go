package producer

import "github.com/IBM/sarama"

func (p *Producer) GetProducer() (*sarama.AsyncProducer, error) {

	producer, err := sarama.NewAsyncProducer(p.addrs, p.config)

	if err != nil {
		p.callback(*p.GetErrorMessage())
		return nil, err
	}

	return &producer, err
}
