package validator

import "time"

func IsExpired(exp time.Time) bool {
	return exp.Before(time.Now())
}
