package ies

// CellGlobalIdList-r16 ::= SEQUENCE OF CGI-Info-Logging-r16
// SIZE (1..32)
type CellglobalidlistR16 struct {
	Value []CgiInfoLoggingR16 `lb:1,ub:32`
}
