package sendm

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	ISendEmail := SendEmail{"from"}

	welcomer := CustomerWelcome{
		Emailer: ISendEmail,
	}
	test := welcomer.Welcome("Bob", "bob@gmail.com")
	result := "bob@gmail.comWelcomeEsse é o body"
	if test != result {
		t.Error("expected:", result, "got:", test)
	}
}

func TestDakeSendEmail(t *testing.T) {
	fakeEmailer := FakeSendEmail{}
	fakeEmailer.On("Send", "bob@gmail.com", "Welcome", "Hi, Bob!").Return(nil)

	welcomer := CustomerWelcome{
		Emailer: fakeEmailer,
	}
	test := welcomer.Welcome("Bob", "bob@gmail.com")
	result := "bob@gmail.comWelcomeEsse é o body"
	if test != result {
		t.Error("expected:", result, "got:", test)
	}
}
