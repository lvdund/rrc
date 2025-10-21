package ies

import "rrc/utils"

// PUR-Parameters-r16 ::= SEQUENCE
type PurParametersR16 struct {
	PurCp5gcCeModeaR16     *utils.ENUMERATED
	PurCp5gcCeModebR16     *utils.ENUMERATED
	PurUp5gcCeModeaR16     *utils.ENUMERATED
	PurUp5gcCeModebR16     *utils.ENUMERATED
	PurCpEpcCeModeaR16     *utils.ENUMERATED
	PurCpEpcCeModebR16     *utils.ENUMERATED
	PurUpEpcCeModeaR16     *utils.ENUMERATED
	PurUpEpcCeModebR16     *utils.ENUMERATED
	PurCpL1ackR16          *utils.ENUMERATED
	PurFrequencyhoppingR16 *utils.ENUMERATED
	PurPuschNbMaxtbsR16    *utils.ENUMERATED
	PurRsrpValidationR16   *utils.ENUMERATED
	PurSubprbCeModeaR16    *utils.ENUMERATED
	PurSubprbCeModebR16    *utils.ENUMERATED
}
