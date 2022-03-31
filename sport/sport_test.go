package sport

import (
	"testing"
)

func TestFilestruct(t *testing.T) {
	var filename = "input.txt"
	test := Filestructconv(filename)
	if test != nil {
		t.Errorf("want:%v,got:%v", nil, test)
	}
}

func TestPrint(t *testing.T) {
	var Name = "Varun"
	test1 := PrintHello(Name)
	if test1 != "Varun" {
		t.Errorf("want:%v,got:%v", nil, test1)
	}
}
