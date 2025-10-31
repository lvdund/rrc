package ies

import "rrc/utils"

// PhyLayerParameters-v1610-addSRS-r16-addSRS-AntennaSwitching-r16 ::= ENUMERATED
type PhylayerparametersV1610AddsrsR16AddsrsAntennaswitchingR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610AddsrsR16AddsrsAntennaswitchingR16EnumeratedNothing = iota
	PhylayerparametersV1610AddsrsR16AddsrsAntennaswitchingR16EnumeratedUsebasic
)
