package ies

import "rrc/utils"

// RSRP-ChangeThresh-r16 ::= ENUMERATED
type RsrpChangethreshR16 struct {
	Value utils.ENUMERATED
}

const (
	RsrpChangethreshR16Db4    = 0
	RsrpChangethreshR16Db6    = 1
	RsrpChangethreshR16Db8    = 2
	RsrpChangethreshR16Db10   = 3
	RsrpChangethreshR16Db14   = 4
	RsrpChangethreshR16Db18   = 5
	RsrpChangethreshR16Db22   = 6
	RsrpChangethreshR16Db26   = 7
	RsrpChangethreshR16Db30   = 8
	RsrpChangethreshR16Db34   = 9
	RsrpChangethreshR16Spare6 = 10
	RsrpChangethreshR16Spare5 = 11
	RsrpChangethreshR16Spare4 = 12
	RsrpChangethreshR16Spare3 = 13
	RsrpChangethreshR16Spare2 = 14
	RsrpChangethreshR16Spare1 = 15
)
