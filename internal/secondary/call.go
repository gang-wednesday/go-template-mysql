package secondary

import (
	"io"
	"net/http"

	"github.com/sony/gobreaker"
)

var Cb *gobreaker.CircuitBreaker

func CallSecondaryApi() (string, error) {
	body, err := Cb.Execute(func() (interface{}, error) {
		resp, err := http.Get("http://localhost:8888/")
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return "", err
	}

	return string(body.([]byte)), nil
}
