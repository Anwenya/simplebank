package mail

import (
	"testing"

	"com.wlq/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGamil(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGamilSender(
		config.EmailGmailSenderName,
		config.EmailGmailSenderAddress,
		config.EmailGmailSenderPassword,
	)
	subject := "A test email"
	content := `
	<h1>hello world</h1>
	`
	to := []string{"wuliuqi0621@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}

func TestSendEmailWithTencent(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewTencentSender(
		config.EmailTencentSenderName,
		config.EmailTencentSenderAddress,
		config.EmailTencentSenderPassword,
	)
	subject := "A test email"
	content := `
	<h1>hello world</h1>
	`
	to := []string{"1737682009@qq.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
