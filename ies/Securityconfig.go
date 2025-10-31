package ies

// SecurityConfig ::= SEQUENCE
// Extensible
type Securityconfig struct {
	Securityalgorithmconfig *Securityalgorithmconfig
	Keytouse                *SecurityconfigKeytouse
}
