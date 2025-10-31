package ies

// RA-Report-r16 ::= SEQUENCE
// Extensible
type RaReportR16 struct {
	CellidR16              RaReportR16CellidR16
	RaInformationcommonR16 *RaInformationcommonR16
	RapurposeR16           RaReportR16RapurposeR16
	SpcellidR17            *CgiInfoLoggingR16 `ext`
}
