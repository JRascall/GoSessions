package sessions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

type sessionFileStorage struct {
	sessions      map[string]ISession
	cleanTimer    *time.Ticker
	expiryMins    int
	filePath      string
	fileExtension string
}

type sessionFile struct {
	IP        string
	Navigator string
	Data      map[string]interface{}
}

// CreateSessionFileStorage allows you create a session system using files
func CreateSessionFileStorage(expiryMins int) ISessionStorage {
	dir, _ := os.Getwd()
	filestorage := &sessionFileStorage{
		expiryMins:    expiryMins,
		cleanTimer:    time.NewTicker(time.Duration(expiryMins) * time.Minute),
		filePath:      fmt.Sprintf("%s/sessions", dir),
		sessions:      make(map[string]ISession),
		fileExtension: "sdata",
	}

	go func() {
		filestorage.Clean()
		for _ = range filestorage.cleanTimer.C {
			filestorage.Clean()
		}
	}()

	return filestorage
}

func (s *sessionFileStorage) isFileValid(path string) bool {
	dotSplit := strings.Split(path, ".")
	ext := dotSplit[len(dotSplit)-1]
	if ext == s.fileExtension {
		return true
	}
	return false
}

// Sessions returns all the sessions
func (s *sessionFileStorage) Sessions() map[string]ISession {
	return s.sessions
}

// Write crates a new session and it's file
func (s *sessionFileStorage) Write(session ISession) {
	s.sessions[session.GetSSID()] = session
	file, _ := os.Create(s.filePath + "/" + session.GetSSID() + "." + s.fileExtension)

	sessMap := make(map[string]interface{})
	Type := reflect.TypeOf(session).Elem()
	val := reflect.ValueOf(session).Elem()
	fields := val.NumField()
	for i := 0; i < fields; i++ {
		field := val.Field(i)
		if field.CanInterface() {
			sessMap[Type.Field(i).Name] = field.Interface()
		}
	}

	data, err := json.Marshal(sessMap)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(data)
	file.Close()
}

// Update allows you to update a sessions information
func (s *sessionFileStorage) Update(session ISession) {

}

// Delete allows you to delete a session by its ssid
func (s *sessionFileStorage) Delete(ssid string) {
	if strings.Contains(ssid, "."+s.fileExtension) == false {
		ssid = ssid + "." + s.fileExtension
	}

	if s.sessions[ssid] != nil {
		delete(s.sessions, ssid)
	}

	os.Remove(fmt.Sprintf("%s/%s", s.filePath, ssid))
}

// Get allows you to grab a session by it's ssid
func (s *sessionFileStorage) Get(ssid string) ISession {
	return s.sessions[ssid]
}

// Clean runs a clean up of the sessions
func (s *sessionFileStorage) Clean() {
	info, _ := ioutil.ReadDir(s.filePath)
	for _, v := range info {
		if name := v.Name(); s.isFileValid(name) {
			if timeSince := time.Now().Sub(v.ModTime()).Minutes(); int(timeSince) >= s.expiryMins {
				s.Delete(name)
			}
		}
	}
}

// Count returns the number of current sessions
func (s *sessionFileStorage) Count() int {
	return 0
}
