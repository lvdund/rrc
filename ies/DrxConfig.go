package ies

import "rrc/utils"

// DRX-Config ::= CHOICE
type DrxConfig interface {
	isDrxConfig()
}

type DrxConfigRelease struct {
	Value struct{}
}

func (*DrxConfigRelease) isDrxConfig() {}

type DrxConfigSetup struct {
	Value interface{}
}

func (*DrxConfigSetup) isDrxConfig() {}
