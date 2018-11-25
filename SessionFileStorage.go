package sessions

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type SessionFileStorage struct {
	sessions      map[string]ISession
	cleanTimer    *time.Ticker
	expiryMins    int
	filePath      string
	fileExtension string
}

func createSessionFileStorage(expiryMins int) ISessionStorage {
	dir, _ := os.Getwd()
	sessionFileStorage := &SessionFileStorage{
		expiryMins:    expiryMins,
		cleanTimer:    time.NewTicker(time.Duration(expiryMins) * time.Minute),
		filePath:      fmt.Sprintf("%s/sessions", dir),
		sessions:      make(map[string]ISession),
		fileExtension: "sdata",
	}

	go func() {
		sessionFileStorage.Clean()
		for _ = range sessionFileStorage.cleanTimer.C {
			sessionFileStorage.Clean()
		}
	}()

	return sessionFileStorage
}

func (s *SessionFileStorage) isFileValid(path string) bool {
	dotSplit := strings.Split(path, ".")
	ext := dotSplit[len(dotSplit)-1]
	if ext == s.fileExtension {
		return true
	}
	return false
}

func (s *SessionFileStorage) Sessions() map[string]ISession {
	return s.sessions
}

func (s *SessionFileStorage) Write(session ISession) {

}

func (s *SessionFileStorage) Update(session ISession) {

}

func (s *SessionFileStorage) Delete(ssid string) {
	if strings.Contains(ssid, "."+s.fileExtension) == false {
		ssid = ssid + "." + s.fileExtension
	}

	if s.sessions[ssid] != nil {
		delete(s.sessions, ssid)
	}

	os.Remove(fmt.Sprintf("%s/%s", s.filePath, ssid))
}

func (s *SessionFileStorage) Get(ssid string) ISession {
	return s.sessions[ssid]
}

func (s *SessionFileStorage) Clean() {
	info, _ := ioutil.ReadDir(s.filePath)
	for _, v := range info {
		if name := v.Name(); s.isFileValid(name) {
			if timeSince := time.Now().Sub(v.ModTime()).Minutes(); int(timeSince) >= s.expiryMins {
				s.Delete(name)
			}
		}
	}
}

func (s *SessionFileStorage) Count() int {
	return 0
}
