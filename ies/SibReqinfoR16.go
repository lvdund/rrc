package ies

import "rrc/utils"

// SIB-ReqInfo-r16 ::= ENUMERATED
type SibReqinfoR16 struct {
	Value utils.ENUMERATED
}

const (
	SibReqinfoR16EnumeratedNothing = iota
	SibReqinfoR16EnumeratedSib12
	SibReqinfoR16EnumeratedSib13
	SibReqinfoR16EnumeratedSib14
	SibReqinfoR16EnumeratedSib20_V1700
	SibReqinfoR16EnumeratedSib21_V1700
	SibReqinfoR16EnumeratedSpare3
	SibReqinfoR16EnumeratedSpare2
	SibReqinfoR16EnumeratedSpare1
)
