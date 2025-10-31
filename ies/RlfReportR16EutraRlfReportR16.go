package ies

import "rrc/utils"

// RLF-Report-r16-eutra-RLF-Report-r16 ::= SEQUENCE
// Extensible
type RlfReportR16EutraRlfReportR16 struct {
	FailedpcellidEutra          CgiInfoeutralogging
	MeasresultRlfReportEutraR16 utils.OCTETSTRING
}
