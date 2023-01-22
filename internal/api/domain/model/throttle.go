package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const (
	ThrottleKeyTypeIP         = "IP"
	ThrottleKeyTypeCustomerID = "customerID"
	ThrottleKeyTypeOperatorID = "operatorID"
	// 1分以内 -> 10回発生 -> 1日停止
	ThrottleCountExpiredDuration  = 1 * time.Minute
	ThrottleCountExpiredThreshold = 10
	ThrottleBlockExpiredDuration  = 24 * time.Hour
)

type Throttle struct {
	HashKey        string    `db:"hash_key" json:"hash_key"`
	KeyType        string    `db:"key_type" json:"key_type"`
	Key            string    `db:"key" json:"key"`
	Count          int       `db:"count" json:"count"`
	CountExpiredAt time.Time `db:"count_expired_at" json:"count_expired_at"`
	BlockExpiredAt time.Time `db:"block_expired_at" json:"block_expired_at"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

func NewThrottle(key, keyType string) *Throttle {
	t := &Throttle{Key: key, KeyType: keyType}
	t.setInitHashKey()
	t.setInitCountExpiredAt()
	t.setInitCount()

	return t
}

func (m *Throttle) TableName() string {
	return "throttles"
}

func (m *Throttle) setInitHashKey() {
	if m.HashKey != "" {
		return
	}

	b := []byte(m.Key)
	data := md5.Sum(b)
	m.HashKey = hex.EncodeToString(data[:])
}

func (m *Throttle) setInitCount() {
	m.Count = 1
}

func (m *Throttle) setInitCountExpiredAt() {
	switch m.KeyType {
	case ThrottleKeyTypeIP, ThrottleKeyTypeCustomerID:
		m.CountExpiredAt = time.Now().Add(ThrottleCountExpiredDuration)
	default:
		m.CountExpiredAt = time.Time{}
	}
}

func (m *Throttle) setInitBlockExpiredAt() {
	switch m.KeyType {
	case ThrottleKeyTypeIP, ThrottleKeyTypeCustomerID:
		m.BlockExpiredAt = time.Now().Add(ThrottleBlockExpiredDuration)
	default:
		m.BlockExpiredAt = time.Time{}
	}
}

func (m *Throttle) Block() {
	m.setInitBlockExpiredAt()
}

func (m *Throttle) IsBlocked() bool {
	if m.BlockExpiredAt.IsZero() {
		return false
	}

	if time.Now().Before(m.BlockExpiredAt) {
		return true
	}

	return false
}

func (m *Throttle) Increase() {
	if m.KeyType == ThrottleKeyTypeIP || m.KeyType == ThrottleKeyTypeCustomerID {
		if time.Now().After(m.CountExpiredAt) {
			m.setInitCountExpiredAt()
			m.Count = 0
		}

		if !m.IsBlocked() {
			m.BlockExpiredAt = time.Time{}
		}

		m.Count += 1
		if m.Count >= ThrottleCountExpiredThreshold {
			m.Block()
		}

		return
	}
	//TODO: default operator
}
