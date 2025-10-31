package ies

import "rrc/utils"

// FailureReportSCG ::= SEQUENCE
// Extensible
type Failurereportscg struct {
	Failuretype          FailurereportscgFailuretype
	Measresultfreqlist   *Measresultfreqlist
	MeasresultscgFailure *utils.OCTETSTRING
	LocationinfoR16      *LocationinfoR16                     `ext`
	FailuretypeV1610     *FailurereportscgFailuretypeV1610    `ext`
	PreviouspscellidR17  *FailurereportscgPreviouspscellidR17 `ext`
	FailedpscellidR17    *FailurereportscgFailedpscellidR17   `ext`
	TimescgfailureR17    *utils.INTEGER                       `lb:0,ub:1023,ext`
	PerrainfolistR17     *PerrainfolistR16                    `ext`
}
