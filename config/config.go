package config

import (
	"os"
	"time"
)

var (
	DSN        = os.Getenv("DSN")
	SignKey    = os.Getenv("SIGNING-KEY")
	ExpireTime = 1000 * time.Second
)
