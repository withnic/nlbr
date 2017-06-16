package nlbr

import (
	"fmt"
	"testing"
)

func ExampleNl2br() {
	fmt.Println(Nl2br("Hello hogehoge.\nTest hoge."))
	// Output:
	//Hello hogehoge.<br>
	//Test hoge.
}

func TestNl2br(t *testing.T) {
	act := Nl2br("Hello hogehoge.\nTest hoge.")
	want := "Hello hogehoge.<br>\nTest hoge."

	if act != want {
		t.Errorf("act %s want %s", act, want)
	}
}

func ExampleRevNl2br() {
	fmt.Println(RevNl2br("Hello hogehoge.<br>Test hoge."))
	// Output:
	//Hello hogehoge.Test hoge.
}

func TestRevNl2br(t *testing.T) {
	act := RevNl2br("Hello hogehoge.<br>Test hoge.<br>")
	want := "Hello hogehoge.Test hoge."
	if act != want {
		t.Errorf("act %s want %s", act, want)
	}
}

func ExampleBr2nl() {
	fmt.Println(Br2nl("Hello hogehoge.<br>Test hoge."))
	// Output:
	//Hello hogehoge.
	//Test hoge.
}

func TestBr2nl(t *testing.T) {
	act := Br2nl("Hello hogehoge.<br>Test hoge.<br>")
	want := "Hello hogehoge.\nTest hoge.\n"

	if act != want {
		t.Errorf("act %s want %s", act, want)
	}
}
