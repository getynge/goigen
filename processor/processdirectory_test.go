package processor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessDirectory(t *testing.T) {
	methods1, pkg1, err := ProcessDirectory("testfiles", "*Example")

	require.NoError(t, err)

	methods2, pkg2, err := ProcessDirectory("testfiles", "Example")

	require.NoError(t, err)

	require.Equal(t, 6, len(methods1))

	require.Equal(t, 1, len(methods2))

	require.Equal(t, "testfiles", pkg1.Name)
	require.Equal(t, "testfiles", pkg2.Name)
}
