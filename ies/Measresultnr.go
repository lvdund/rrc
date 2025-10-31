package ies

import "rrc/utils"

// MeasResultNR ::= SEQUENCE
// Extensible
type Measresultnr struct {
	Physcellid        *Physcellid
	Measresult        *MeasresultnrMeasresult
	CgiInfo           *CgiInfonr                     `ext`
	ChocandidateR17   *utils.BOOLEAN                 `ext`
	ChoconfigR17      *[]CondtriggerconfigR16        `lb:1,ub:2,ext`
	TriggeredeventR17 *MeasresultnrTriggeredeventR17 `ext`
}
