package ies

import "rrc/utils"

// PDSCH-RE-MappingQCL-Config-r11-optionalSetOfFields-r11 ::= SEQUENCE
type PdschReMappingqclConfigR11OptionalsetoffieldsR11 struct {
	CrsPortscountR11           PdschReMappingqclConfigR11OptionalsetoffieldsR11CrsPortscountR11
	CrsFreqshiftR11            utils.INTEGER `lb:0,ub:5`
	MbsfnSubframeconfiglistR11 *PdschReMappingqclConfigR11OptionalsetoffieldsR11MbsfnSubframeconfiglistR11
	PdschStartR11              PdschReMappingqclConfigR11OptionalsetoffieldsR11PdschStartR11
}
