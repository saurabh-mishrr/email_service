package controllers

import (
	//configs "emailer_service/configs"
	"emailer_service/helpers"
	"log"
	"net/http"
	"strconv"
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/go-mailer"
	"github.com/streadway/amqp"

)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Produce(c *gin.Context) {
	mailPort, _ := strconv.Atoi(helpers.DotEnvVal("MAIL_PORT"))
	mailerConfig := mailer.Config{
		Host: helpers.DotEnvVal("MAIL_HOST"),
		Port: mailPort,
		Username: helpers.DotEnvVal("MAIL_USERNAME"),
		Password: helpers.DotEnvVal("MAIL_PASSWORD"),
		FromAddr: helpers.DotEnvVal("MAIL_USERNAME"),
		FromAlias: helpers.DotEnvVal("MAIL_NAME"),
		UseCommand: false,
	}

	sender := mailer.New(mailerConfig)

	subject := "Dummy Mail"

	content := `<h1>Hello</h1> <br/><br/> <span style="color:red"> This is the rich message body </span>`

	to := []string{"test1@mailhog.local"}

	err := sender.Send(subject, content, to...)

	if err != nil {
		println("error while sending the e-mail: " + err.Error())
	}


	//c.BindJSON(data)
	//body, _ := ioutil.ReadAll(c.Request.Body)
	c.JSON(http.StatusOK, gin.H{"data": "done"})
	/* conn, err := amqp.Dial("amqp://guest:guest@email-service:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	failOnError(err, "Failed to publish a message")
	c.String(http.StatusOK, "Message queued successfully") */
}

func Receive(c *gin.Context) {
	conn, err := amqp.Dial("amqp://guest:guest@email-service:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
