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

	if !Url("www.google.com/test/test.jpg") {
		t.Error("google.com - it should be a url")
	}

	if !Url("www.google.com/test/test.jpg?v=123&v2=321") {
		t.Error("google.com - it should be a url")
	}

	if !Url("http://www.google.com/test/test.jpg?v=123&v2=321") {
		t.Error("google.com - it should be a url")
	}

	if !Url("http://www.google.com:8080/test/test.jpg?v=123&v2=321") {
		t.Error("google.com - it should be a url")
	}

	if !Url("http://www.google.com:8080") {
		t.Error("google.com - it should be a url")
	}

	if Url("2014/newimages/icon_15.jpg") {
		t.Error("2014/newimages/icon_15.jpg - it should not be a url")
	}

	if Url("google_com") {
		t.Error("google_com - it should not be a url")
	}
}
