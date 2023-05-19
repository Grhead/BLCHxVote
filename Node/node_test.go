package Node

import (
	"VOX2/Transport/Network"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	var testPack *Network.Package
	testPack.Option = GetBlockConst
	testPack.Data = "motor"

	blocks, err := GetBlocks(testPack)
	if err != nil {
		return
	}
	// assert for not nil (good when you expect something)
	if assert.NotNil(t, blocks) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, 123, 123, "they should be equal")

	}
}
