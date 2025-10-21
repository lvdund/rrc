package ies

import "rrc/utils"

// TPC-PDCCH-Config ::= CHOICE
type TpcPdcchConfig interface {
	isTpcPdcchConfig()
}

type TpcPdcchConfigRelease struct {
	Value struct{}
}

func (*TpcPdcchConfigRelease) isTpcPdcchConfig() {}

type TpcPdcchConfigSetup struct {
	Value interface{}
}

func (*TpcPdcchConfigSetup) isTpcPdcchConfig() {}
