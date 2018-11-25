package sessions

import (
	"os"
	"testing"
)

var fileSession ISessionStorage

func TestMain(m *testing.M) {
	fileSession = createSessionFileStorage(5)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestCreateFileStorage(t *testing.T) {
	for {

	}
}
