package ies

import "rrc/utils"

// Phy-ParametersCommon-crossSlotScheduling-r16-non-SharedSpectrumChAccess-r16 ::= ENUMERATED
type PhyParameterscommonCrossslotschedulingR16NonSharedspectrumchaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCrossslotschedulingR16NonSharedspectrumchaccessR16EnumeratedNothing = iota
	PhyParameterscommonCrossslotschedulingR16NonSharedspectrumchaccessR16EnumeratedSupported
)
