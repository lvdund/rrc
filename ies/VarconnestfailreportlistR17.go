package ies

// VarConnEstFailReportList-r17 ::= SEQUENCE
type VarconnestfailreportlistR17 struct {
	ConnestfailreportlistR17 []VarconnestfailreportR16 `lb:1,ub:maxCEFReportR17`
}
