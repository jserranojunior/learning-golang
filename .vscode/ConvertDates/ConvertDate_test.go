package cvdt

import "testing"

func TestPTbrToUS(t *testing.T) {
	test := ConvertDate.PtToUs("10/06/1993")
	expected := "2020-06-10"
	if test != expected {
		t.Error("Expected:", expected, "Got:", text)
	}

}
