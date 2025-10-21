package ies

import "rrc/utils"

// DMRS-Config-r11 ::= CHOICE
type DmrsConfigR11 interface {
	isDmrsConfigR11()
}

type DmrsConfigR11Release struct {
	Value struct{}
}

func (*DmrsConfigR11Release) isDmrsConfigR11() {}

type DmrsConfigR11Setup struct {
	Value interface{}
}

func (*DmrsConfigR11Setup) isDmrsConfigR11() {}
