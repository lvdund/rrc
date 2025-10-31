package ies

import "rrc/utils"

// SystemInformationBlockType8-parameters1XRTT ::= SEQUENCE
type Systeminformationblocktype8Parameters1xrtt struct {
	CsfbRegistrationparam1xrtt     *CsfbRegistrationparam1xrtt
	Longcodestate1xrtt             *utils.BITSTRING `lb:42,ub:42`
	Cellreselectionparameters1xrtt *Cellreselectionparameterscdma2000
}
