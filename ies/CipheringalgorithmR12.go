package ies

import "rrc/utils"

// CipheringAlgorithm-r12 ::= ENUMERATED
// Extensible
type CipheringalgorithmR12 struct {
	Value utils.ENUMERATED
}

const (
	CipheringalgorithmR12Eea0      = 0
	CipheringalgorithmR12Eea1      = 1
	CipheringalgorithmR12Eea2      = 2
	CipheringalgorithmR12Eea3V1130 = 3
	CipheringalgorithmR12Spare4    = 4
	CipheringalgorithmR12Spare3    = 5
	CipheringalgorithmR12Spare2    = 6
	CipheringalgorithmR12Spare1    = 7
)
