package ies

import "rrc/utils"

// BandCombinationParameters-v1630-interBandPowerSharingAsyncDAPS-r16 ::= ENUMERATED
type BandcombinationparametersV1630InterbandpowersharingasyncdapsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandcombinationparametersV1630InterbandpowersharingasyncdapsR16EnumeratedNothing = iota
	BandcombinationparametersV1630InterbandpowersharingasyncdapsR16EnumeratedSupported
)
