package ies

import "rrc/utils"

// ReselectionInfoRelay-r13-minHyst-r13 ::= ENUMERATED
type ReselectioninforelayR13MinhystR13 struct {
	Value utils.ENUMERATED
}

const (
	ReselectioninforelayR13MinhystR13EnumeratedNothing = iota
	ReselectioninforelayR13MinhystR13EnumeratedDb0
	ReselectioninforelayR13MinhystR13EnumeratedDb3
	ReselectioninforelayR13MinhystR13EnumeratedDb6
	ReselectioninforelayR13MinhystR13EnumeratedDb9
	ReselectioninforelayR13MinhystR13EnumeratedDb12
	ReselectioninforelayR13MinhystR13EnumeratedDbinf
)
