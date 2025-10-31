package ies

import "rrc/utils"

// MsgA-DMRS-Config-r16 ::= SEQUENCE
type MsgaDmrsConfigR16 struct {
	MsgaDmrsAdditionalpositionR16 *MsgaDmrsConfigR16MsgaDmrsAdditionalpositionR16
	MsgaMaxlengthR16              *MsgaDmrsConfigR16MsgaMaxlengthR16
	MsgaPuschDmrsCdmGroupR16      *utils.INTEGER `lb:0,ub:1`
	MsgaPuschNrofportsR16         *utils.INTEGER `lb:0,ub:1`
	MsgaScramblingid0R16          *utils.INTEGER `lb:0,ub:65535`
	MsgaScramblingid1R16          *utils.INTEGER `lb:0,ub:65535`
}
