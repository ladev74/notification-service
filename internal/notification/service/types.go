package service

import (
	"context"

	"go.uber.org/zap"
)

type Config struct {
	SenderEmail     string `yaml:"SENDER_EMAIL"`
	SenderPassword  string `yaml:"SENDER_PASSWORD"`
	SMTPHost        string `yaml:"SMTP_HOST"`
	SMTPPort        int    `yaml:"SMTP_PORT"`
	SkipVerify      bool   `yaml:"SKIP_VERIFY"`
	MaxRetries      int    `yaml:"MAX_RETRIES"`
	BasicRetryPause int    `yaml:"BASIC_RETRY_PAUSE"`
}

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type TempEmailMessageWithTime struct {
	Time    string `json:"time"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type EmailMessageWithTime struct {
	Time  string
	Email EmailMessage
}

type SMTPClient struct {
	config *Config
	logger *zap.Logger
}

type EmailSender interface {
	SendEmail(ctx context.Context, email EmailMessage) error
}

func New(config *Config, logger *zap.Logger) *SMTPClient {
	return &SMTPClient{
		config: config,
		logger: logger,
	}
}
