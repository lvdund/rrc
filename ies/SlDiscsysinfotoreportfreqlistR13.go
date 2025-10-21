package ies

import "rrc/utils"

// SL-DiscSysInfoToReportFreqList-r13 ::= SEQUENCE OF ARFCN-ValueEUTRA-r9
// SIZE (1..maxFreq)
type SlDiscsysinfotoreportfreqlistR13 struct {
	Value []ArfcnValueeutraR9
}
