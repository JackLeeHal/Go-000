package Week02

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommentsService(t *testing.T) {
	_, err := CommentsService("test")
	require.NoError(t, err)
}

func TestBalanceService(t *testing.T) {
	_, err := BalanceService("test")
	require.Error(t, err)
}
