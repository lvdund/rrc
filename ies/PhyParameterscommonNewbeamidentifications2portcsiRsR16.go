package ies

import "rrc/utils"

// Phy-ParametersCommon-newBeamIdentifications2PortCSI-RS-r16 ::= ENUMERATED
type PhyParameterscommonNewbeamidentifications2portcsiRsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonNewbeamidentifications2portcsiRsR16EnumeratedNothing = iota
	PhyParameterscommonNewbeamidentifications2portcsiRsR16EnumeratedSupported
)
