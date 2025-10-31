package ies

// SecurityAlgorithmConfig ::= SEQUENCE
// Extensible
type Securityalgorithmconfig struct {
	Cipheringalgorithm     CipheringalgorithmR12
	Integrityprotalgorithm SecurityalgorithmconfigIntegrityprotalgorithm
}
