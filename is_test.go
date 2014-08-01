package is

import "testing"

func TestIsEmail(t *testing.T) {
	if !Email("xxx@xxx.com") {
		t.Error("xxx@xxx.com - it should be an email")
	}

	if Email("xxx#xxx.com") {
		t.Error("xxx#xxx.com - it should not be an email")
	}
}

func TestIsNumString(t *testing.T) {
	if !NumString("1234") {
		t.Error("1234 - it should be a num string")
	}

	if NumString("@1234") {
		t.Error("1234 - it should not be a num string")
	}

	if NumString("1@234") {
		t.Error("1234 - it should not be a num string")
	}

	if NumString("1234@") {
		t.Error("1234 - it should not be a num string")
	}
}

func TestIsUrl(t *testing.T) {
	if !Url("google.com") {
		t.Error("google.com - it should be a url")
	}

	if !Url("www.google.com/test/test.jpg") {
		t.Error("www.google.com/test/test.jpg - it should be a url")
	}

	if !Url("www.google.com/test/test.jpg?v=123&v2=321") {
		t.Error("www.google.com/test/test.jpg?v=123&v2=321 - it should be a url")
	}

	if !Url("http://www.google.com/test/test.jpg?v=123&v2=321") {
		t.Error("http://www.google.com/test/test.jpg?v=123&v2=321 - it should be a url")
	}

	if !Url("http://www.google.com:8080/test/test.jpg?v=123&v2=321") {
		t.Error("http://www.google.com:8080/test/test.jpg?v=123&v2=321 - it should be a url")
	}

	if !Url("http://www.google.com:8080") {
		t.Error("http://www.google.com:8080 - it should be a url")
	}

	if !Url("http://127.0.0.1/img/logo.jpg") {
		t.Error("http://127.0.0.1/img/logo.jpg - it should be a url")
	}

	if !Url("http://127.0.0.1:8888/img/logo.jpg") {
		t.Error("http://127.0.0.1:8888/img/logo.jpg - it should be a url")
	}

	if Url("2014/newimages/icon_15.jpg") {
		t.Error("2014/newimages/icon_15.jpg - it should not be a url")
	}

	if Url("google_com") {
		t.Error("google_com - it should not be a url")
	}
}
