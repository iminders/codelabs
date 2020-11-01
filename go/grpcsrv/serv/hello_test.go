package serv

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  actual := Hello()

  assert.Equal(t, "hello", actual)

}
