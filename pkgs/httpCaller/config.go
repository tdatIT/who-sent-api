package http

import (
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	clientTimeout             = 5 * time.Second
	dialContextTimeout        = 5 * time.Second
	clientTLSHandshakeTimeout = 5 * time.Second
	clientRetryWaitTime       = 300 * time.Millisecond
	retryCount                = 2
	debugMode                 = false
)

type config struct {
	ClientTimeout             time.Duration
	DialContextTimeout        time.Duration
	ClientTLSHandshakeTimeout time.Duration
	ClientRetryWaitTime       time.Duration
	RetryCount                int
	RetryCondition            func(r *resty.Response, err error) bool
	DebugMode                 bool
}

var defaultConfig = config{
	ClientTimeout:             clientTimeout,
	DialContextTimeout:        dialContextTimeout,
	ClientTLSHandshakeTimeout: clientTLSHandshakeTimeout,
	ClientRetryWaitTime:       clientRetryWaitTime,
	RetryCount:                retryCount,
	RetryCondition:            defaultRetryCondition,
	DebugMode:                 debugMode,
}

type Config interface {
	apply(*config)
}
