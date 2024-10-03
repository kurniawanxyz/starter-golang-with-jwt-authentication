package helper

import (
	"fmt"
	"strconv"

	"github.com/kurniawanxzy/backend-olshop/config"
	"gopkg.in/gomail.v2"
)

// EmailSender struct to hold email configuration
type EmailSender struct {
    SMTPHost     string
    SMTPPort     string
    Username     string
    Password     string
    From         string
}

// SendEmail sends an email to a specific recipient
func SendEmail(to string, subject string, body string) error {

	e := EmailSender{
		SMTPHost: config.ENV.SMTPHost,
		SMTPPort: config.ENV.SMTPPort,
		Username: config.ENV.SMTPUser,
		Password: config.ENV.SMTPPass,
		From:    config.ENV.SMTPUser,
	}

    m := gomail.NewMessage()

    // Set the sender and recipient
    m.SetHeader("From", e.From)
    m.SetHeader("To", to)
    
    // Set the subject and body
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body) // You can change "text/html" to "text/plain" if sending plain text

    port, err := strconv.Atoi(e.SMTPPort)
    if err != nil {
        fmt.Println("Invalid SMTP port:", err)
        return err
    }
    d := gomail.NewDialer(e.SMTPHost, port, e.Username, e.Password)

    // Send the email
    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Error sending email:", err)
        return err
    }

    fmt.Println("Email sent successfully to", to)
    return nil
}
