//go:generate mockgen -destination=mock/$GOPACKAGE.gen.go -package=$GOPACKAGE github.com/basslove/daradara/internal/api/domain/$GOPACKAGE Throttler,CustomerAuthenticator,OperatorAuthenticator

package service
