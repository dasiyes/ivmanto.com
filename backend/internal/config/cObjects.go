package config

import "fmt"

type ServiceConfig struct {
	Name      string   `yaml:"name"`
	ProjectID string   `yaml:"project_id"`
	Port      string   `yaml:"port"`
	EmailCfg  EmailCfg `yaml:"email_cfg"`
}

type EmailCfg struct {
	// The host address of the email server
	SmtpHost string `yaml:"smtp_host"`
	// The server TLS PORT
	SmptPort string `yaml:"smtp_port"`
	// SendFrom is the account the will be used to send the email from
	SendFrom string `yaml:"send_from"`
	// SendFromAlias is the account to appear as a sender
	SendFromAlias string `yaml:"send_from_alias"`
}

func (sc *ServiceConfig) GetEmailCfg() *EmailCfg {
	return &sc.EmailCfg
}

func (ec *EmailCfg) GetHostAddress() string {
	return fmt.Sprintf("%s:%s", ec.SmtpHost, ec.SmptPort)
}
