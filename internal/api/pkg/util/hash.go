package util

import (
    "crypto/sha256"
    "encoding/hex"
)

func HashFromString(s string) string {
    hashedBytes := sha256.Sum256([]byte(s))

    return hex.EncodeToString(hashedBytes[:])
}
