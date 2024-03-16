package main

import (
	"fmt"

	"github.com/IBM/sarama"
	consumer "github.com/rafaelsouzaribeiro/apache-kafka/consumer"
	producer "github.com/rafaelsouzaribeiro/apache-kafka/producer"
	//authjwt "github.com/rafaelsouzaribeiro/jwt-auth"
)

func main() {

	go Producer()
	go Consumer()
	select {}
	// Espera que todas as goroutines terminem

	// 	mx := http.NewServeMux()
	// 	cre, err := authjwt.NewCredential(3600, "secretkey", nil)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	Consumer()
	// 	Producer()
	// 	Main()
	// 	// Protected routes
	// 	mx.HandleFunc("/route1", cre.AuthMiddleware(rota1Handler))
	// 	mx.HandleFunc("/route2", cre.AuthMiddleware(rota2Handler))

	// 	// Public route
	// 	mx.HandleFunc("/public-route", rotaPublicaHandler)

	// 	http.ListenAndServe(":8080", mx)
	// }

	// func rota1Handler(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Rota 1 protegida por token")
	// }

	// func rota2Handler(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Rota 2 protegida por token")
	// }

	// func rotaPublicaHandler(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Rota pública acessível sem token")
	// }

	// func Main() {
	// 	cre, err := authjwt.NewCredential(1, "secretkey", nil)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	claims := map[string]interface{}{
	// 		"firstname": "Paulo",
	// 		"lastname":  "Eduardo",
	// 	}

	// 	token, err := cre.CreateToken(claims)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	claims, err = cre.ExtractClaims(token)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	//		println(fmt.Sprint(claims["lastname"]), fmt.Sprint(claims["firstname"]))
	//	}
}

func Consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Consumer.Offsets.AutoCommit.Enable = true

	con := consumer.NewConsumer([]string{"springboot:9092"}, "contact-adm",
		[]string{"contact-adm-insert"}, config, func(messages []string) {
			// Processe as mensagens recebidas aqui
			fmt.Println("Mensagens recebidas:", messages[0])
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

	err = produc.SendMessage(prod)

	if err != nil {
		fmt.Printf("Error sending message: %v", message)
	}

}
