package ies

import "rrc/utils"

// MAC-MainConfig-NB-r13 ::= SEQUENCE
// Extensible
type MacMainconfigNbR13 struct {
	UlSchConfigR13                 *interface{}
	DrxConfigR13                   *DrxConfigNbR13
	TimealignmenttimerdedicatedR13 Timealignmenttimer
	LogicalchannelsrConfigR13      *interface{}
}
