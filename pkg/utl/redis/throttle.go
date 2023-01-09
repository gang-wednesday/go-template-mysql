package redis

import (
	"math"
	"time"
)

func StartVisit(key string, exp time.Duration) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	ttl := math.Ceil(exp.Seconds())
	_, err = conn.Do("SETEX", key, ttl, 1)

	return err

}

func IncVisites(key string) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	_, err = conn.Do("INCR", key)
	return err
}
