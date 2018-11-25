package sessions

type ISessionStorage interface {
	Sessions() map[string]ISession

	Write(session ISession)
	Update(session ISession)
	Delete(ssid string)
	Get(ssid string) ISession
	Clean()
	Count() int
}
