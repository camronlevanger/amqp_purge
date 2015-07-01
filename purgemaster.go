package main

import (
    "fmt"
    "github.com/golang/glog"
    "github.com/streadway/amqp"
)

// main entry point. reads configs, sets up cache, etc.
func main() {

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    _, declare_err := ch.QueueDeclare(
        "queue_name", // name
        true,         // durable
        false,        // delete when usused
        false,        // exclusive
        false,        // no-wait
        nil,          // arguments
    )
    failOnError(declare_err, "Failed to declare a queue")

    purged, err := ch.QueuePurge("queue_name", false)
    failOnError(err, "Failed to register a consumer")
    glog.Infof("%d messages purged from queue", purged)
}

// function to be called on fatal errors, this kills the app
func failOnError(err error, msg string) {
    if err != nil {
        glog.Fatalf("%s: %s", msg, err)
        panic(fmt.Sprintf("%s: %s", msg, err))
    }
}
