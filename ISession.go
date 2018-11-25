package sessions

type ISession interface {
	SSID() string
	Expiry() string
	IP() string
	Navigator() string

	Add(key string, data interface{})
	Delete(key string)
	Get(key string) interface{}
}
