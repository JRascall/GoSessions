package sessions

type Session struct {
	SSID      string
	IP        string
	Navigator string
	test      int
}

func createSession(ssid string) ISession {
	return &Session{
		SSID:      ssid,
		IP:        "localhost",
		Navigator: "browser",
	}
}

func (s *Session) GetSSID() string {
	return s.SSID
}

func (s *Session) GetIP() string {
	return s.IP
}

func (s *Session) GetNavigator() string {
	return s.Navigator
}

func (s *Session) Add(key string, data interface{}) {

}

func (s *Session) Delete(key string) {

}

func (s *Session) Retrive(key string) interface{} {
	return nil
}
