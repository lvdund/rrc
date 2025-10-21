package ies

import "rrc/utils"

// AUL-Config-r15 ::= CHOICE
type AulConfigR15 interface {
	isAulConfigR15()
}

type AulConfigR15Release struct {
	Value struct{}
}

func (*AulConfigR15Release) isAulConfigR15() {}

type AulConfigR15Setup struct {
	Value interface{}
}

func (*AulConfigR15Setup) isAulConfigR15() {}
