package ies

import "rrc/utils"

// DRX-Config-NB-r13 ::= CHOICE
type DrxConfigNbR13 interface {
	isDrxConfigNbR13()
}

type DrxConfigNbR13Release struct {
	Value struct{}
}

func (*DrxConfigNbR13Release) isDrxConfigNbR13() {}

type DrxConfigNbR13Setup struct {
	Value interface{}
}

func (*DrxConfigNbR13Setup) isDrxConfigNbR13() {}
