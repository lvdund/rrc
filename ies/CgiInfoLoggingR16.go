package ies

// CGI-Info-Logging-r16 ::= SEQUENCE
type CgiInfoLoggingR16 struct {
	PlmnIdentityR16     PlmnIdentity
	CellidentityR16     Cellidentity
	TrackingareacodeR16 *Trackingareacode
}
