package lib

import (
	"crypto/md5"
	"encoding/base64"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func IfThenElse(condition bool, a bool, b bool) bool {
	if condition {
		return a
	}
	return b
}

func HashPassword(password, hash string) bool {
	bcryptPattern := hash[:3]
	if bcryptPattern != "$2y" {
		hashMd5 := md5.Sum([]byte(password))
		stringHashMd5 := base64.StdEncoding.EncodeToString(hashMd5[:])
		return hash == stringHashMd5
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func RandomString(n int) string {
	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func ParseDuration(param string) time.Duration {
	duration, _ := time.ParseDuration(param)

	return duration
}

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}

	return
}

func IsWeekend(t time.Time) bool {
	t = t.UTC()
	switch t.Weekday() {
	case time.Saturday:
		return true
	case time.Sunday:
		h, m, _ := t.Clock()
		if h < 12+10 {
			return true
		}
		if h == 12+10 && m <= 5 {
			return true
		}
	}
	return false
}
