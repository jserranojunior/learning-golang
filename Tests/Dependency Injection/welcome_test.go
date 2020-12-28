package sendm

import (
	"testing"
)

func TestFSendEmail(t *testing.T) {
	welcomer := CustomerWelcome{}

	test := welcomer.Hello("Bob", FSend)
	result := "Voce esta no MOCK " + "Bob"
	if test != result {
		t.Error("expected:", result, "got:", test)
	}
}

func TestSendEmail(t *testing.T) {
	welcomer := CustomerWelcome{}

	test := welcomer.Hello("Bob", Send)
	result := "Voce esta sem MOCK " + "Bob"
	if test != result {
		t.Error("expected:", result, "got:", test)
	}
}
