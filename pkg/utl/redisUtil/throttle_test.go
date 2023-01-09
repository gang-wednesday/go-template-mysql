package redisutil

import (
	"context"
	"errors"
	"fmt"
	"go-template/testutls"
	"testing"
	"time"

	redismock "github.com/go-redis/redismock/v8"
)

func TestSetCounter(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
		err     error
	}{
		{
			name:    "succesfull redis set counter operation",
			wantErr: false,
			err:     nil,
		},
		{
			name:    "failed redis set counter operation",
			wantErr: true,
			err:     errors.New("connection error"),
		},
	}
	db, mock := redismock.NewClientMock()
	defer db.Close()
	for _, tt := range cases {
		var err error
		if tt.name == "succesfull redis set counter operation" {
			duration := time.Minute.Nanoseconds() * windowTime
			mock.Regexp().ExpectSet(string(testutls.MockIpAddress[0]), `[0-9]+`, time.Duration(duration)).SetVal("31")

		} else {
			duration := time.Minute.Nanoseconds() * windowTime
			mock.Regexp().
				ExpectSet(string(testutls.MockIpAddress[0]), `[0-9]+`, time.Duration(duration)).
				SetErr(fmt.Errorf("connection error"))
		}
		err = SetCounter(context.Background(), db, string(testutls.MockIpAddress[0]), 30)
		if err == nil && tt.wantErr == true {
			t.Fatal("expected an error")
		}
		if err != nil && tt.wantErr == false {
			t.Fatal(err)
		}

	}
}

func TestGetCounter(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
		err     error
	}{
		{
			name:    "successfull get counter execution",
			wantErr: false,
			err:     nil,
		}, {
			name:    "get rate limit exceeded error",
			wantErr: true,
			err:     errors.New("request limit breached!!!"),
		},
	}
	db, mock := redismock.NewClientMock()
	defer db.Close()
	for _, tt := range cases {
		if tt.name == "successfull get counter execution" {
			mock.ExpectGet(string(testutls.MockIpAddress[0])).SetVal("10")
		} else {
			mock.ExpectGet(string(testutls.MockIpAddress[0])).SetVal("200")
		}
		_, err := GetCounter(context.Background(), db, string(testutls.MockIpAddress[0]))
		if err != nil {
			if tt.wantErr == false || tt.err.Error() != err.Error() {
				t.Fatal(tt.wantErr)
			}
		}
		if err == nil && tt.wantErr != false {
			t.Fatal("expected an error")
		}
	}
}
