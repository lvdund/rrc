package ies

import "rrc/utils"

// IntegrityProtAlgorithm ::= utils.ENUMERATED // Extensible
type Integrityprotalgorithm struct {
	Value utils.ENUMERATED
}

const (
	IntegrityprotalgorithmEnumeratedNothing = iota
	IntegrityprotalgorithmEnumeratedNia0
	IntegrityprotalgorithmEnumeratedNia1
	IntegrityprotalgorithmEnumeratedNia2
	IntegrityprotalgorithmEnumeratedNia3
	IntegrityprotalgorithmEnumeratedSpare4
	IntegrityprotalgorithmEnumeratedSpare3
	IntegrityprotalgorithmEnumeratedSpare2
	IntegrityprotalgorithmEnumeratedSpare1
)
