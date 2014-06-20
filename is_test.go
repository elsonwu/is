package is

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	if !Email("xxx@xxx.com") {
		t.Error("xxx@xxx.com - it should be an email")
	}

	if Email("xxx#xxx.com") {
		t.Error("xxx#xxx.com - it should not be an email")
	}
}

func TestIsUrl(t *testing.T) {
	if !Url("google.com") {
		t.Error("google.com - it should be a url")
	}

	if Url("google_com") {
		t.Error("google_com - it should not be a url")
	}
}
