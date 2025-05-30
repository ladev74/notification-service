package service

import (
	"context"

	"go.uber.org/zap"

	"notification/internal/config"
)

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type Timestamp struct {
	Time string `json:"time"`
}

type TempEmailWithTime struct {
	Time    string `json:"time"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type EmailWithTime struct {
	Time  string
	Email Email
}

type EmailService struct {
	config *config.MailSender
	logger *zap.Logger
}

type EmailSender interface {
	SendMessage(ctx context.Context, email Email) error
}

func New(config *config.MailSender, logger *zap.Logger) *EmailService {
	return &EmailService{
		config: config,
		logger: logger,
	}
}
