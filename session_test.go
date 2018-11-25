package sessions

import "testing"

func TestCreatesession(t *testing.T) {
	fileSession := createSessionFileStorage(30)

	if fileSession != nil {
		t.Error("Failed!")
	}
}
