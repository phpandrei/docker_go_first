package main

import (
	"testing"
	"gofirst/mypackage"
	"github.com/stretchr/testify/assert"
)

func TestGetHello(t *testing.T) {
	actual := mypackage.GetHello()
	expected := "Hello World by package!"

	if expected != actual {
    	t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}

	assert.NotNil(t, actual)	
	assert.Equal(t, expected, actual)
}
