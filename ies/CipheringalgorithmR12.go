package ies

import "rrc/utils"

// CipheringAlgorithm-r12 ::= utils.ENUMERATED // Extensible
type CipheringalgorithmR12 struct {
	Value utils.ENUMERATED
}

const (
	CipheringalgorithmR12EnumeratedNothing = iota
	CipheringalgorithmR12EnumeratedEea0
	CipheringalgorithmR12EnumeratedEea1
	CipheringalgorithmR12EnumeratedEea2
	CipheringalgorithmR12EnumeratedEea3_V1130
	CipheringalgorithmR12EnumeratedSpare4
	CipheringalgorithmR12EnumeratedSpare3
	CipheringalgorithmR12EnumeratedSpare2
	CipheringalgorithmR12EnumeratedSpare1
)
