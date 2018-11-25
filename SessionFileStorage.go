package sessions

import (
	"fmt"
	"os"
	"time"
)

type SessionFileStorage struct {
	sessions   []ISession
	cleanTimer *time.Ticker
	expiryMins time.Duration
	file       *os.File
}

func createSessionFileStorage(expiryMins time.Duration) ISessionStorage {
	sessionFileStorage := &SessionFileStorage{
		expiryMins: expiryMins,
		cleanTimer: time.NewTicker(expiryMins * time.Minute),
	}

	dir, _ := os.Getwd()
	fmt.Println(dir)

	go func() {
		for c := range sessionFileStorage.cleanTimer.C {
			_ = c
			sessionFileStorage.Clean()
		}
	}()

	sessionFileStorage.Clean()
	return sessionFileStorage
}

func (s *SessionFileStorage) ExpiryMinutes() time.Duration {
	return s.expiryMins
}

func (s *SessionFileStorage) Sessions() []ISession {
	return s.sessions
}

func (s *SessionFileStorage) Write(session ISession) {

}

func (s *SessionFileStorage) Update(session ISession) {

}

func (s *SessionFileStorage) Delete(ssid string) {

}

func (s *SessionFileStorage) Get(ssid string) ISession {
	return nil
}

func (s *SessionFileStorage) Clean() {

}

func (s *SessionFileStorage) Count() int {
	return 0
}
