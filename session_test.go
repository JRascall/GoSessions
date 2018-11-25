package sessions

import (
	"os"
	"testing"
)

var fileSession ISessionStorage

func TestMain(m *testing.M) {
	fileSession = createSessionFileStorage(30)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestCreateFileStorage(t *testing.T) {

}
