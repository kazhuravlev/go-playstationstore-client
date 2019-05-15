package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := New()

	require.NoError(t, err)
	assert.NotNil(t, c)
}

func TestClient_GetAllGames(t *testing.T) {
	c, _ := New()

	games, meta, err := c.GetAllGames(1)
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.NotNil(t, games)

	assert.Equal(t, 1, meta.MinPage)
	assert.Equal(t, 1, meta.CurrentPage)
	assert.True(t, meta.MaxPage > 1)
	assert.Equal(t, 30, meta.ObjectsOnPage)

	assert.Equal(t, meta.ObjectsOnPage, len(games))

	for _, game := range games {
		assert.NotEmpty(t, game.Name)
		assert.NotEmpty(t, game.ImagesSrcSet)
		assert.NotEmpty(t, game.Link)
	}
}

