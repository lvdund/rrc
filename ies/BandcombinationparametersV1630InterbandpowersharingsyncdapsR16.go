package ies

import "rrc/utils"

// BandCombinationParameters-v1630-interBandPowerSharingSyncDAPS-r16 ::= ENUMERATED
type BandcombinationparametersV1630InterbandpowersharingsyncdapsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1630InterbandpowersharingsyncdapsR16EnumeratedNothing = iota
	BandcombinationparametersV1630InterbandpowersharingsyncdapsR16EnumeratedSupported
)
