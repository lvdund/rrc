package ies

import "rrc/utils"

// RMTC-Config-r13 ::= CHOICE
// Extensible
type RmtcConfigR13 interface {
	isRmtcConfigR13()
}

type RmtcConfigR13Release struct {
	Value struct{}
}

func (*RmtcConfigR13Release) isRmtcConfigR13() {}

type RmtcConfigR13Setup struct {
	Value interface{}
}

func (*RmtcConfigR13Setup) isRmtcConfigR13() {}
