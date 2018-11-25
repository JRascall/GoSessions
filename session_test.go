package sessions

import (
	"os"
	"testing"

	"github.com/google/uuid"
)

var fileSession ISessionStorage

func TestMain(m *testing.M) {
	fileSession = CreateSessionFileStorage(1)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestFileStorageSessionWrite(t *testing.T) {
	uid, _ := uuid.NewUUID()
	sess := createSession(uid.String())
	fileSession.Write(sess)

	for {
	}
}
