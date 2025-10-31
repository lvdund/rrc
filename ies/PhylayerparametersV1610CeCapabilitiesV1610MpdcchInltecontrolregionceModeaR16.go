package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-mpdcch-InLteControlRegionCE-ModeA-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610MpdcchInltecontrolregionceModeaR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610MpdcchInltecontrolregionceModeaR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610MpdcchInltecontrolregionceModeaR16EnumeratedSupported
)
