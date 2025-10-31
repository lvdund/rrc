package ies

import "rrc/utils"

// SCellActivationRS-Config-r17 ::= SEQUENCE
// Extensible
type ScellactivationrsConfigR17 struct {
	ScellactivationrsIdR17 ScellactivationrsConfigidR17
	ResourcesetR17         NzpCsiRsResourcesetid
	GapbetweenburstsR17    *utils.INTEGER `lb:0,ub:31`
	QclInfoR17             TciStateid
}
