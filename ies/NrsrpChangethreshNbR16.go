package ies

import "rrc/utils"

// NRSRP-ChangeThresh-NB-r16 ::= ENUMERATED
type NrsrpChangethreshNbR16 struct {
	Value utils.ENUMERATED
}

const (
	NrsrpChangethreshNbR16Db4    = 0
	NrsrpChangethreshNbR16Db6    = 1
	NrsrpChangethreshNbR16Db8    = 2
	NrsrpChangethreshNbR16Db10   = 3
	NrsrpChangethreshNbR16Db14   = 4
	NrsrpChangethreshNbR16Db18   = 5
	NrsrpChangethreshNbR16Db22   = 6
	NrsrpChangethreshNbR16Db26   = 7
	NrsrpChangethreshNbR16Db30   = 8
	NrsrpChangethreshNbR16Db34   = 9
	NrsrpChangethreshNbR16Spare6 = 10
	NrsrpChangethreshNbR16Spare5 = 11
	NrsrpChangethreshNbR16Spare4 = 12
	NrsrpChangethreshNbR16Spare3 = 13
	NrsrpChangethreshNbR16Spare2 = 14
	NrsrpChangethreshNbR16Spare1 = 15
)
