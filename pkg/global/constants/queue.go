package constants

var Queue = &nastQueue{
	ActiveLevel: &queue{
		Subject:   "activeLevel-subject",
		QueueName: "activeLevel-queue",
	},
}

type nastQueue struct {
	ActiveLevel *queue
}

type queue struct {
	Subject   string
	QueueName string
}
