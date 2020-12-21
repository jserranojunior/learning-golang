package sendm

import "fmt"

//EmailSender interface
type EmailSender interface {
	Send(to, subject, body string) error
}

//CustomerWelcome to welcome
type CustomerWelcome struct {
	Emailer EmailSender
}

//Welcome send email of welcome
func (welcomer *CustomerWelcome) Welcome(name, email string) error {
	body := fmt.Sprintf("Hi, %s!", name)
	subject := "Welcome"

	err := welcomer.Emailer.Send(email, subject, body)
	return err
}
