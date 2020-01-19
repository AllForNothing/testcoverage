package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMyout(t *testing.T) {
	assert.Equal(t, "good", Myout())
}
