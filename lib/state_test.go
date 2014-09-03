package ozdata

import (
	"testing"
)

func Test_NewStateDataFailsToFindFile(t *testing.T) {
	_, err := NewStateData()
	if err != nil {
		t.Log("Could not find json")
	}
}
