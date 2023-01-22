package model_test

import (
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThrottle_NewThrottle(t *testing.T) {
	throttle := model.NewThrottle("0.0.0.1", model.ThrottleKeyTypeIP)
	require.NotEqual(t, "", throttle.HashKey)
	require.False(t, throttle.CountExpiredAt.IsZero())
	require.True(t, throttle.BlockExpiredAt.IsZero())
}

func TestThrottle_IsBlocked(t *testing.T) {
	throttle := model.NewThrottle("12345", model.ThrottleKeyTypeCustomerID)
	require.False(t, throttle.IsBlocked())
	throttle.Block()
	require.True(t, throttle.IsBlocked())
}
