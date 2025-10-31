package ies

// RE-MappingQCLConfigToReleaseList-r11 ::= SEQUENCE OF PDSCH-RE-MappingQCL-ConfigId-r11
// SIZE (1..maxRE-MapQCL-r11)
type ReMappingqclconfigtoreleaselistR11 struct {
	Value []PdschReMappingqclConfigidR11 `lb:1,ub:maxREMapqclR11`
}
