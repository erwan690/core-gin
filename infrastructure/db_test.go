package infrastructure

import (
	"testing"

	"core-gin/lib"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	db := NewDatabase(lib.GetLogger(), &lib.Env{})
	assert.NotNil(t, db)
	assert.NotNil(t, db.DB)
}
