package ies

// SL-DiscSysInfoToReportFreqList-r13 ::= SEQUENCE OF ARFCN-ValueEUTRA-r9
// SIZE (1..maxFreq)
type SlDiscsysinfotoreportfreqlistR13 struct {
	Value []ArfcnValueeutraR9 `lb:1,ub:maxFreq`
}
