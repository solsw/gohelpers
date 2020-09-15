// Package sendgridhelper contains SendGrid (https://github.com/sendgrid/sendgrid-go) helpers.
package sendgridhelper

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// PlainContent creates plain text mail.Content containing 'content'.
func PlainContent(content string) *mail.Content {
	return mail.NewContent("text/plain", content)
}

// HTMLContent creates HTML mail.Content containing 'content'.
func HTMLContent(content string) *mail.Content {
	return mail.NewContent("text/html", content)
}

// FileAttachment creates mail.Attachment containing binary content of file 'fileName'.
func FileAttachment(fileName string) (*mail.Attachment, error) {
	bb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	a := mail.NewAttachment()
	a.SetType("application/octet-stream")
	a.SetFilename(fileName)
	a.SetContent(base64.StdEncoding.EncodeToString(bb))
	return a, nil
}
