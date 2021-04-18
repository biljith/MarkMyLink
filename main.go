package main

import (
	"MarkMyLink/controller"
	"net/http"
	"os"
)


//helper function to check the return value for each amqp call
//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}
//
//func receive(){
//	//connect to RabbitMQ server
//	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//	failOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	//create a channel
//	ch, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	//declare a queue for us to consume
//	q, err := ch.QueueDeclare(
//		"task_queue", // name
//		true,   // durable
//		false,   // delete when unused
//		false,   // exclusive
//		false,   // no-wait
//		nil,     // arguments
//	)
//	failOnError(err, "Failed to declare a queue")
//
//	err = ch.Qos(
//		1,     // prefetch count
//		0,     // prefetch size
//		false, // global
//	)
//	failOnError(err, "Failed to set QoS")
//
//	msgs, err := ch.Consume(
//		q.Name, // queue
//		"",     // consumer
//		false,  // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		nil,    // args
//	)
//	failOnError(err, "Failed to register a consumer")
//
//	forever := make(chan bool)
//
//	go func() {
//		for d := range msgs {
//			log.Printf("Received a message: %s", d.Body)
//			log.Printf("Done")
//			d.Ack(false)
//		}
//	}()
//
//	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
//	<-forever
//}

func main() {
	//receive()
	// Documentation for serving static files - 
	// https://www.alexedwards.net/blog/serving-static-sites-with-go
	http.HandleFunc(
		"/signup", controller.Signup)
	http.HandleFunc(
		"/login", controller.Login)
	http.HandleFunc(
		"/bookmarks", controller.GetBookmarks)
	fs := http.FileServer(http.Dir("./client/build"))
	http.Handle("/", fs)

	//http.HandleFunc("/", index)
	// handles the URL 'localhost:8080/bookmarks
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Commenting out below code. Uncomment after adding react UI.
	// http.HandleFunc(
	// 	"/bookmarks", controller.Index)
	// http.HandleFunc(
	// 	"/signup", controller.Signup)
	// http.HandleFunc(
	// 	"/login", controller.Login)
	http.HandleFunc(
		"/addBookmark", controller.AddBookmark)
	 http.HandleFunc(
	 	"/updateBookmark", controller.UpdateBookmark)
	 http.HandleFunc(
	 	"/deleteBookmark", controller.DeleteBookmark)
	// http.HandleFunc(
	// 	"/bookmarks", controller.Index)
	// this is where the web application will listen
	http.ListenAndServe(":" + port, nil)
}