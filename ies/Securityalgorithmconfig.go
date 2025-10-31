package ies

// SecurityAlgorithmConfig ::= SEQUENCE
// Extensible
type Securityalgorithmconfig struct {
	Cipheringalgorithm     Cipheringalgorithm
	Integrityprotalgorithm *Integrityprotalgorithm
}
