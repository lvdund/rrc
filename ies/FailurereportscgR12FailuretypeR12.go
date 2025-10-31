package ies

import "rrc/utils"

// FailureReportSCG-r12-failureType-r12 ::= ENUMERATED
type FailurereportscgR12FailuretypeR12 struct {
	Value utils.ENUMERATED
}

const (
	FailurereportscgR12FailuretypeR12EnumeratedNothing = iota
	FailurereportscgR12FailuretypeR12EnumeratedT313_Expiry
	FailurereportscgR12FailuretypeR12EnumeratedRandomaccessproblem
	FailurereportscgR12FailuretypeR12EnumeratedRlc_Maxnumretx
	FailurereportscgR12FailuretypeR12EnumeratedScg_Changefailure
)
