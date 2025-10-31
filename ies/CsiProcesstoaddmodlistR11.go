package ies

// CSI-ProcessToAddModList-r11 ::= SEQUENCE OF CSI-Process-r11
// SIZE (1..maxCSI-Proc-r11)
type CsiProcesstoaddmodlistR11 struct {
	Value []CsiProcessR11 `lb:1,ub:maxCSIProcR11`
}
