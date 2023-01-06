package redis

import (
	"encoding/json"

	redigo "github.com/gomodule/redigo/redis"
)

func Connect() (redigo.Conn, error) {
	conn, err := redigo.Dial("tcp", "")
	return conn, err
}

func SetKeyValue(key string, data interface{}) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	mData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, string(mData))

	return err

}

func GetKeyValue(key string) (interface{}, error) {
	conn, err := Connect()
	if err != nil {
		return "", err
	}
	data, err := conn.Do("GET", key)
	if err != nil {
		return "", err
	}

	return data, nil
}
