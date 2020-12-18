package sun

import "testing"

func TestSun(t *testing.T){
	test := sun(2, 5)
	result := 7
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	} 
}
