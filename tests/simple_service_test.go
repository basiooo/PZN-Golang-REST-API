package tests

import (
	"golang-restfull-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}
func TestSimpleServiceSUccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
