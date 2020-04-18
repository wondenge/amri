package amri

import (
	"testing"
)

func TestMockCommand_implements(t *testing.T) {
	var _ Command = new(MockCommand)
}
