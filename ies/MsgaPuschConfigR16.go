package ies

import "rrc/utils"

// MsgA-PUSCH-Config-r16 ::= SEQUENCE
type MsgaPuschConfigR16 struct {
	MsgaPuschResourcegroupaR16 *MsgaPuschResourceR16
	MsgaPuschResourcegroupbR16 *MsgaPuschResourceR16
	MsgaTransformprecoderR16   *MsgaPuschConfigR16MsgaTransformprecoderR16
	MsgaDatascramblingindexR16 *utils.INTEGER `lb:0,ub:1023`
	MsgaDeltapreambleR16       *utils.INTEGER `lb:0,ub:6`
}
