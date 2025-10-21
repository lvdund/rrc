package ies

import "rrc/utils"

// SRS-TPC-PDCCH-Config-r14 ::= CHOICE
type SrsTpcPdcchConfigR14 interface {
	isSrsTpcPdcchConfigR14()
}

type SrsTpcPdcchConfigR14Release struct {
	Value struct{}
}

func (*SrsTpcPdcchConfigR14Release) isSrsTpcPdcchConfigR14() {}

type SrsTpcPdcchConfigR14Setup struct {
	Value interface{}
}

func (*SrsTpcPdcchConfigR14Setup) isSrsTpcPdcchConfigR14() {}
