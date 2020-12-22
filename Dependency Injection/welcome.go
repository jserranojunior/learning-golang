package sendm

//Hello send email of welcome
func (welcomer *CustomerWelcome) Hello(name string, Send SendEmail) string {
	return Send(name)
}
