package test

import (
	"testing"

	"github.com/surajgoraicse/utils"
)

func TestIsEven(t *testing.T) {

	got := utils.IsEven(2)
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
	got = utils.IsEven(3)
	want = false
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
func TestHashString(t *testing.T) {
	got := utils.HashString("hello world")
	want := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	got = utils.HashString("hello")
	want = "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
	got = utils.HashString("")
	want = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
