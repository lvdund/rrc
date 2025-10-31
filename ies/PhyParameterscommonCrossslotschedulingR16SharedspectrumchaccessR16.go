package ies

import "rrc/utils"

// Phy-ParametersCommon-crossSlotScheduling-r16-sharedSpectrumChAccess-r16 ::= ENUMERATED
type PhyParameterscommonCrossslotschedulingR16SharedspectrumchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCrossslotschedulingR16SharedspectrumchaccessR16EnumeratedNothing = iota
	PhyParameterscommonCrossslotschedulingR16SharedspectrumchaccessR16EnumeratedSupported
)
