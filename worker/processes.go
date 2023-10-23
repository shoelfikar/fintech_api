package worker

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
}

type RedisTaskProcessor struct {
	server *asynq.Server
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt) TaskProcessor {

	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Println("error")
			}),
		},
	)

	return &RedisTaskProcessor{
		server: server,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	return processor.server.Start(mux)
}