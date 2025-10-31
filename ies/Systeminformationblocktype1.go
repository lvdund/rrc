package ies

import "rrc/utils"

// SystemInformationBlockType1 ::= SEQUENCE
type Systeminformationblocktype1 struct {
	Cellaccessrelatedinfo *Systeminformationblocktype1Cellaccessrelatedinfo
	Cellselectioninfo     *Systeminformationblocktype1Cellselectioninfo
	PMax                  *PMax
	Freqbandindicator     Freqbandindicator
	Schedulinginfolist    Schedulinginfolist
	TddConfig             *TddConfig
	SiWindowlength        Systeminformationblocktype1SiWindowlength
	Systeminfovaluetag    utils.INTEGER `lb:0,ub:31`
	Noncriticalextension  *Systeminformationblocktype1V890
}
