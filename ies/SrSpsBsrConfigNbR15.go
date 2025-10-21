package ies

import "rrc/utils"

// SR-SPS-BSR-Config-NB-r15 ::= CHOICE
type SrSpsBsrConfigNbR15 interface {
	isSrSpsBsrConfigNbR15()
}

type SrSpsBsrConfigNbR15Release struct {
	Value struct{}
}

func (*SrSpsBsrConfigNbR15Release) isSrSpsBsrConfigNbR15() {}

type SrSpsBsrConfigNbR15Setup struct {
	Value interface{}
}

func (*SrSpsBsrConfigNbR15Setup) isSrSpsBsrConfigNbR15() {}
