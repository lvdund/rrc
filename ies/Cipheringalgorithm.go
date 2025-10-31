package ies

import "rrc/utils"

// CipheringAlgorithm ::= utils.ENUMERATED // Extensible
type Cipheringalgorithm struct {
	Value utils.ENUMERATED
}

const (
	CipheringalgorithmEnumeratedNothing = iota
	CipheringalgorithmEnumeratedNea0
	CipheringalgorithmEnumeratedNea1
	CipheringalgorithmEnumeratedNea2
	CipheringalgorithmEnumeratedNea3
	CipheringalgorithmEnumeratedSpare4
	CipheringalgorithmEnumeratedSpare3
	CipheringalgorithmEnumeratedSpare2
	CipheringalgorithmEnumeratedSpare1
)
