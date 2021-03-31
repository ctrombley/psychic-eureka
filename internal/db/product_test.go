package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadProducts(t *testing.T) {
	p, err := LoadProducts(context.Background(), 12353, 11214)
	require.NoError(t, err)
	require.NotNil(t, p)
	require.Equal(t, 1, len(p))
	require.Equal(t, "DNA Fire OG", p[0].Name)
}

func TestFindLocationCount(t *testing.T) {
	c, err := FindLocationCountByCoords(context.Background(), 9654, -118.2436849, 34.0522342)
	require.NoError(t, err)
	require.NotNil(t, c)
	require.Equal(t, 11, c)

	c, err = FindLocationCountByZip(context.Background(), 9654, "90012")
	require.NoError(t, err)
	require.NotNil(t, c)
	require.Equal(t, 5, c)
}
