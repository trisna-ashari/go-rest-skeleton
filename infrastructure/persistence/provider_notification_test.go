package persistence_test

import (
	"go-rest-skeleton/infrastructure/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/gomail.v2"
)

func TestNewSMTPClient_Success(t *testing.T) {
	conf := InitConfig()
	smtpDialer, errDial := persistence.NewSMTPClient(conf.SMTPConfig)
	if errDial != nil {
		t.Fatalf("want non error, got %#v", errDial)
	}

	var smtpClient *gomail.Dialer
	assert.NoError(t, errDial)
	assert.IsType(t, smtpDialer.Client, smtpClient)
}

func TestNewSMTPClient_Failed(t *testing.T) {
	conf := InitConfig()
	conf.SMTPConfig.SMTPHost = "invalid host"
	_, errDial := persistence.NewSMTPClient(conf.SMTPConfig)

	assert.Error(t, errDial)
}

func TestNewNotificationService(t *testing.T) {

}
