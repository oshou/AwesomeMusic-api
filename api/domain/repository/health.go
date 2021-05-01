//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
// Package repository is domain-service
package repository

// IHealthRepository is interface for Health-Check
type IHealthRepository interface {
	Ping() error
}
