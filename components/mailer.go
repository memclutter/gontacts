package components

import (
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

var (
	MailerCh chan *gomail.Message
)

func MailerInit(host string, port int, username string, password string) {
	MailerCh = make(chan *gomail.Message)

	go func() {
		d := gomail.NewDialer(host, port, username, password)

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-MailerCh:
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
				// Close the connection to the SMTP server if no email was sent in
				// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						//panic(err)
						log.Print(err)
					}
					open = false
				}
			}
		}
	}()
}
