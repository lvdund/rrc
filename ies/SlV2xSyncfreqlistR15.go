package ies

import "rrc/utils"

// SL-V2X-SyncFreqList-r15 ::= SEQUENCE OF ARFCN-ValueEUTRA-r9
// SIZE (1..maxFreqV2X-r14)
type SlV2xSyncfreqlistR15 struct {
	Value []ArfcnValueeutraR9
}
