package ies

import "rrc/utils"

// ParametersCDMA2000-r11-parameters1XRTT-r11 ::= SEQUENCE
type Parameterscdma2000R11Parameters1xrttR11 struct {
	CsfbRegistrationparam1xrttR11     *CsfbRegistrationparam1xrtt
	CsfbRegistrationparam1xrttExtR11  *CsfbRegistrationparam1xrttV920
	Longcodestate1xrttR11             *utils.BITSTRING `lb:42,ub:42`
	Cellreselectionparameters1xrttR11 *Cellreselectionparameterscdma2000R11
	AcBarringconfig1xrttR11           *AcBarringconfig1xrttR9
	CsfbSupportfordualrxuesR11        *utils.BOOLEAN
	CsfbDualrxtxsupportR11            *bool
}
