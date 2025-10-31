package ies

// CSI-ProcessToReleaseList-r11 ::= SEQUENCE OF CSI-ProcessId-r11
// SIZE (1..maxCSI-Proc-r11)
type CsiProcesstoreleaselistR11 struct {
	Value []CsiProcessidR11 `lb:1,ub:maxCSIProcR11`
}
