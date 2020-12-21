package sendm

//SendEmail from string
type SendEmail struct {
	From string
}

//Send email
func (sender *SendEmail) Send(to, subject, body string) error {
	emailer := &SendEmail{
		From: "hi@welcome.com",
	}
	welcomer := &CustomerWelcome{
		Emailer: emailer,
	}
	err := welcomer.Welcome("Bob", "bob@smith.com")
	return (err)
}
