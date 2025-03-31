package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	// }

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")

	assert.NotEqual(t, 3, 4, "3 should not be equal to 4")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not excecuting")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Excecuting")
}
