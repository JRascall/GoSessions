package main

import (
	"time"
)

type ISessionStorage interface {
	ExpiryMinutes() time.Duration
	Sessions() []ISession

	Write(session ISession)
	Update(session ISession)
	Delete(ssid string)
	LoadAll() []ISession
	Get(ssid string) ISession
	Clean()
	Count() int
}
