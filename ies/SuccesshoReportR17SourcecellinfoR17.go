package ies

import "rrc/utils"

// SuccessHO-Report-r17-sourceCellInfo-r17 ::= SEQUENCE
type SuccesshoReportR17SourcecellinfoR17 struct {
	SourcepcellidR17   CgiInfoLoggingR16
	SourcecellmeasR17  *MeasresultsuccesshonrR17
	RlfInsourcedapsR17 *utils.BOOLEAN
}
