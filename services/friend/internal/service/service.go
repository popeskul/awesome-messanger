package service

import "github.com/popeskul/awesome-messanger/services/friend/internal/core/ports"

type service struct {
	processor ports.OutboxProcessor
}

func NewService(processor ports.OutboxProcessor) ports.Service {
	return &service{
		processor: processor,
	}
}

func (s *service) OutboxProcessor() ports.OutboxProcessor {
	return s.processor
}
