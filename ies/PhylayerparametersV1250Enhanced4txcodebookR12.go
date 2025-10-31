package ies

import "rrc/utils"

// PhyLayerParameters-v1250-enhanced-4TxCodebook-r12 ::= ENUMERATED
type PhylayerparametersV1250Enhanced4txcodebookR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250Enhanced4txcodebookR12EnumeratedNothing = iota
	PhylayerparametersV1250Enhanced4txcodebookR12EnumeratedSupported
)
