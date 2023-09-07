package mail

//! I did not create a test Gmail account for this test so it will fail
// TODO: create a test Gmail account and use it to test this function

import (
	"testing"

	"github.com/abdulrahman-02/G-Bank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("../")
	require.NoError(t, err)

	sender := NewGmailSender(
		config.EmailSenderName,
		config.EmailSenderAddress,
		config.EmailSenderPassword,
	)

	subject := "Test Subject"
	content := `
	<h1>Test Content</h1>
	<p>This is a test email</p>
	`
	to := []string{"abderrahmanebenaissa740@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
