package is

import "regexp"

const (
	EmailPattern = `^[a-zA-Z0-9!#$%&\'*+\\/=?^_{|}~-]+(?:\.[a-zA-Z0-9!#$%&\'*+\\/=?^_{|}~-]+)*@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`
	UrlPattern   = `((http|https|ftp){1}\:\/\/)?([a-zA-Z0-9]+[a-zA-Z0-9\-\.]+[a-zA-Z0-9]+\.(net|cn|co|hk|tw|com|edu|gov|us|int|mil|org|int|mil|vg|uk|idv|tk|se|nz|nu|nl|ms|jp|jobs|it|ind|gen|firm|in|gs|fr|fm|eu|es|de|bz|be|at|am|ag|mx|asia|ws|xxx|tv|cc|ca|mobi|me|biz|arpa|info|name|pro|aero|coop|museum|ly|eg|mk)|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3})){1}(:[a-zA-Z0-9]*)?\/?([a-zA-Z0-9\-\._\?\'\/\\\+&amp;%\$#\=~])*`
)

func Email(v string) bool {
	ok, err := regexp.MatchString(EmailPattern, v)
	return ok && err == nil
}

func Url(v string) bool {
	ok, err := regexp.MatchString(UrlPattern, v)
	return ok && err == nil
}
