package tests

import (
	"golang-restfull-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializeConnection("Database")
	assert.NotNil(t, connection)
	cleanup()
}
