package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-etws-CMAS-RxInConnCE-ModeA-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModeaR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610EtwsCmasRxinconnceModeaR16EnumeratedSupported
)
