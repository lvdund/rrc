package ies

// SON-Parameters-NB-r16 ::= SEQUENCE
type SonParametersNbR16 struct {
	AnrReportR16  *SonParametersNbR16AnrReportR16
	RachReportR16 *SonParametersNbR16RachReportR16
}
