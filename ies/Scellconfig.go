package ies

import "rrc/utils"

// SCellConfig ::= SEQUENCE
// Extensible
type Scellconfig struct {
	Scellindex                      Scellindex
	Scellconfigcommon               *Servingcellconfigcommon
	Scellconfigdedicated            *Servingcellconfig
	Smtc                            *SsbMtc                            `ext`
	ScellstateR16                   *ScellconfigScellstateR16          `ext`
	SecondarydrxGroupconfigR16      *utils.BOOLEAN                     `ext`
	PreconfgapstatusR17             *utils.BITSTRING                   `lb:maxNrofGapIdR17,ub:maxNrofGapIdR17,ext`
	GoodservingcellevaluationbfdR17 *GoodservingcellevaluationR17      `ext`
	Scellsib20R17                   *utils.Setuprelease[Scellsib20R17] `ext`
}
