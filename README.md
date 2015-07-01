# amqp_purge
Golang program to connect to an AMQP (RabbitMQ) queue and purge all of it's unacknowledged messages.

Replace the queue URI, and queue_name to suit your needs, put it on a cron (or whatever), and purge away.

This could be made far more flexible with little effort, but this suited my needs at the time. I will likely never expand on what is available here.
