package ies

// VarWLAN-Status-r13 ::= SEQUENCE
type VarwlanStatusR13 struct {
	StatusR13 WlanStatusR13
	StatusR14 *WlanStatusV1430
}
