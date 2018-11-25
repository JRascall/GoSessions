package sessions

type ISession interface {
	GetSSID() string
	GetIP() string
	GetNavigator() string

	Add(key string, data interface{})
	Delete(key string)
	Retrive(key string) interface{}
}
