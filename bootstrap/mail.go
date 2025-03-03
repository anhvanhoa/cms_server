package bootstrap

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/wneessen/go-mail"
)

type ConfigMail struct {
	Host     string
	Port     int
	UserName string
	Password string
	Mailer   string
	Email    string
	Name     string
	TSL      *tls.Config
}

type MailProvider interface {
	SendMail(to []string, subject, body string, data map[string]any) error
	// SendMailWithAttachment(to []string, subject, body, attachment string) error
}

type mailProvider struct {
	mail     *mail.Msg
	provider *mail.Client
	config   *ConfigMail
}

func (m *mailProvider) SendMail(to []string, subject, body string, data map[string]any) error {
	m.mail.SetGenHeader("Content-Type", "text/html")
	m.mail.SetGenHeader("charset", "utf-8")
	m.mail.SetGenHeader("X-Mailer", m.config.Mailer)
	m.mail.SetGenHeader("Date", time.Now().Format(time.RFC1123Z))
	m.mail.Subject(subject)

	if content, err := mustache.ParseString(body); err != nil {
		return err
	} else if body, err = content.Render(data); err != nil {
		return err
	} else {
		m.mail.AddAlternativeString(mail.TypeTextHTML, body)
	}
	m.mail.FromFormat(m.config.Name, m.config.Email)
	m.mail.To(to...)

	if err := m.provider.Send(m.mail); err != nil {
		return err
	}
	return nil
}

func NewMailProvider(cf *ConfigMail) (MailProvider, error) {
	provider, err := mail.NewClient(
		cf.Host, mail.WithPort(cf.Port),
		mail.WithTLSConfig(cf.TSL),
		mail.WithUsername(cf.UserName),
		mail.WithPassword(cf.Password),
	)
	if err != nil {
		log.Fatalf("Failed to create mail provider: %v", err)
		return nil, err
	}
	return &mailProvider{
		mail:     mail.NewMsg(),
		provider: provider,
		config:   cf,
	}, nil
}
