package di_test

import (
	"testing"

	"github.com/harluo/di"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	container := di.New().Instance()
	assert.NotNil(t, container)
}
