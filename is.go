package is

import (
	"reflect"
	"regexp"
	"time"
)

var typeTime = reflect.TypeOf(time.Time{})

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

// borrow from "gopkg.in/mgo.v2/bson"
func Zero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Slice:
		return v.Len() == 0
	case reflect.Map:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct:
		if v.Type() == typeTime {
			return v.Interface().(time.Time).IsZero()
		}
		for i := v.NumField() - 1; i >= 0; i-- {
			if !Zero(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}
