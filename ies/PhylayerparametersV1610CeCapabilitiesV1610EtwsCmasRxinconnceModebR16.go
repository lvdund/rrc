package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-etws-CMAS-RxInConnCE-ModeB-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModebR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModebR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModebR16EnumeratedSupported
)
