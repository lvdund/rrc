package ies

import "rrc/utils"

// RE-MappingQCLConfigToAddModList-r11 ::= SEQUENCE OF PDSCH-RE-MappingQCL-Config-r11
// SIZE (1..maxRE-MapQCL-r11)
type ReMappingqclconfigtoaddmodlistR11 struct {
	Value []PdschReMappingqclConfigR11
}
