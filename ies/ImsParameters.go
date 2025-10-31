package ies

// IMS-Parameters ::= SEQUENCE
// Extensible
type ImsParameters struct {
	ImsParameterscommon  *ImsParameterscommon
	ImsParametersfrxDiff *ImsParametersfrxDiff
}
