package main

/*
import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

var sessionStorage ISessionStorage

func Sessions() gin.HandlerFunc {
	sessionStorage = createSessionFileStorage(database, 30)
	sessionStorage.LoadAll()

	return func(c *gin.Context) {
		session := GetCurrentSession(c)

		ip := c.Request.Header.Get("REMOTE_ADDR")
		userAgent := c.Request.Header.Get("User-Agent")

		//Facebook embed creating sessions!!!
		if strings.Contains(userAgent, "facebook") {
			return
		}

		//Make sure not to create session for the scraper!
		//TODO: Move all exceptions to a function
		if strings.Contains(c.Request.RequestURI, "scrape") {
			c.Next()
			return
		}

		if session == nil {
			session = CreateSession(c, ip, userAgent)
		} else {
			timeParsed, _ := time.Parse(MySQLTimeFormat, session.Expiry)
			duration := time.Now().Sub(timeParsed)
			minutes := int64(duration.Minutes())
			if minutes >= 0 {
				sessionStorage.Delete(session.SSID)
				session = CreateSession(c, ip, userAgent)
			} else {
				session.Expiry = time.Now().Add(time.Hour).Format(MySQLTimeFormat)
				sessionStorage.Update(session)
			}
		}

		c.Set("session", session)
		c.Next()
	}
}

func GetCurrentSession(c *gin.Context) *Session {
	ssid, err := c.Cookie("ssid")
	if err != nil {
		return nil
	}
	return sessionStorage.Get(ssid)
}

func CreateSession(c *gin.Context, ip string, navigator string) *Session {
	ssid := uuid.NewV4().String()
	session := &Session{SSID: ssid, Data: map[string]interface{}{}, Expiry: time.Now().Add(1 * time.Hour).Format(MySQLTimeFormat), IP: ip, Navigator: navigator}
	c.Header("Set-Cookie", "ssid="+ssid+"; HttpOnly; Path=/")

	session.Add("Auth", false)
	sessionStorage.Write(session)
	return session
}
*/
