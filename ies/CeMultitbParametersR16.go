package ies

import "rrc/utils"

// CE-MultiTB-Parameters-r16 ::= SEQUENCE
type CeMultitbParametersR16 struct {
	PdschMultitbCeModeaR16       *utils.ENUMERATED
	PdschMultitbCeModebR16       *utils.ENUMERATED
	PuschMultitbCeModeaR16       *utils.ENUMERATED
	PuschMultitbCeModebR16       *utils.ENUMERATED
	CeMultitb64qamR16            *utils.ENUMERATED
	CeMultitbEarlyterminationR16 *utils.ENUMERATED
	CeMultitbFrequencyhoppingR16 *utils.ENUMERATED
	CeMultitbHarqAckbundlingR16  *utils.ENUMERATED
	CeMultitbInterleavingR16     *utils.ENUMERATED
	CeMultitbSubprbR16           *utils.ENUMERATED
}
