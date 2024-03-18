<strong>version v1.4.0</strong>
<br />
<strong>added message failure:</strong>
<br />

```go

func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	message := []byte("Hello World!")

	produc := producer.NewProducer([]string{"springboot:9092"}, "contact-adm-insert",
		sarama.ByteEncoder(message), config, func(messages string) {
			fmt.Println("Message failure:", messages)
		})
	prod, err := produc.GetProducer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*prod).Close(); err != nil {
			panic(err)
		}
	}()

	produc.SendMessage(prod)

}


```

<strong>Version v1.3.0</strong>
<br />
<strong>Run in separate terminal:</strong>
<br />
<strong>Consumer</strong>
<br />

```go

package main

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/consumer"
)

func main() {
	go Consumer()

	select {}

}

func Consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Consumer.Offsets.AutoCommit.Enable = true

	con := consumer.NewConsumer([]string{"springboot:9092"}, "contact-adm",
		[]string{"contact-adm-insert"}, config, func(messages []string) {
			// Processe as mensagens recebidas aqui
			fmt.Println("Message received:", messages[0])
		})

	client, err := con.GetConsumer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			panic(err)
		}

	}()

	cancel, err := con.VerifyConsumer(client)
	defer cancel()

	if err != nil {
		panic(err)
	}

	con.VerifyError(client)

}

```

<br />
<strong>Producer</strong>
<br />

```go

package main

import (
	"sync"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/producer"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		Producer()
		wg.Done()
	}()

	wg.Wait()
}

func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	message := []byte("Hello World2!")

	produc := producer.NewProducer([]string{"springboot:9092"}, "contact-adm-insert",
		sarama.ByteEncoder(message), config)
	prod, err := produc.GetProducer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*prod).Close(); err != nil {
			panic(err)
		}
	}()

	produc.SendMessage(prod)

}

```