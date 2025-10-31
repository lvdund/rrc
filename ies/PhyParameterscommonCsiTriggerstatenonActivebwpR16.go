package ies

import "rrc/utils"

// Phy-ParametersCommon-csi-TriggerStateNon-ActiveBWP-r16 ::= ENUMERATED
type PhyParameterscommonCsiTriggerstatenonActivebwpR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCsiTriggerstatenonActivebwpR16EnumeratedNothing = iota
	PhyParameterscommonCsiTriggerstatenonActivebwpR16EnumeratedSupported
)
