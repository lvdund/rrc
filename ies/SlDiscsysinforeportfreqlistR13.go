package ies

// SL-DiscSysInfoReportFreqList-r13 ::= SEQUENCE OF SL-DiscSysInfoReport-r13
// SIZE (1.. maxSL-DiscSysInfoReportFreq-r13)
type SlDiscsysinforeportfreqlistR13 struct {
	Value []SlDiscsysinforeportR13 `lb:1,ub:maxSLDiscsysinforeportfreqR13`
}
