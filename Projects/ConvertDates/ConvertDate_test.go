package cvdt

import (
	"testing"
	"time"
)

func TestPTbrToUS(t *testing.T) {
	test := ConvertDateToUs("10/06/1993T11:45:26.371Z")
	expected := "1993-06-10"
	if test != expected {
		t.Error("Expected:", expected, "Got:", test)
	}
}

func TestPTbrToUSWithoutSeconds(t *testing.T) {
	test := ConvertDateToUs("10/06/1993")
	expected := "1993-06-10"
	if test != expected {
		t.Error("Expected:", expected, "Got:", test)
	}
}

func TestUsSecondsbrToUSWithoutSeconds(t *testing.T) {
	test := ConvertDateToUs("1993-06-10T15:04:05.000Z")
	expected := "1993-06-10"
	if test != expected {
		t.Error("Expected:", expected, "Got:", test)
	}
}

func TestTimeNowToUS(t *testing.T) {
	test := ConvertDateToUs("2020-12-20 12:05:07.13640638 -0300 -03 m=+0.001360288")
	expected := "2020-12-20"
	if test != expected {
		t.Error("Expected:", expected, "Got:", test, time.Now())
	}
}
