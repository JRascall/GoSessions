package sessions

import (
	"time"
)

type ISessionStorage interface {
	ExpiryMinutes() time.Duration
	Sessions() []ISession

	Write(session ISession)
	Update(session ISession)
	Delete(ssid string)
	Get(ssid string) ISession
	Clean()
	Count() int
}
