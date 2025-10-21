package ies

import "rrc/utils"

// TPC-PDCCH-ConfigSCell-r13 ::= CHOICE
type TpcPdcchConfigscellR13 interface {
	isTpcPdcchConfigscellR13()
}

type TpcPdcchConfigscellR13Release struct {
	Value struct{}
}

func (*TpcPdcchConfigscellR13Release) isTpcPdcchConfigscellR13() {}

type TpcPdcchConfigscellR13Setup struct {
	Value interface{}
}

func (*TpcPdcchConfigscellR13Setup) isTpcPdcchConfigscellR13() {}
