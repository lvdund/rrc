package ies

import "rrc/utils"

// CSI-RS-ConfigNZPToAddModList-r11 ::= SEQUENCE OF CSI-RS-ConfigNZP-r11
// SIZE (1..maxCSI-RS-NZP-r11)
type CsiRsConfignzptoaddmodlistR11 struct {
	Value []CsiRsConfignzpR11
}
