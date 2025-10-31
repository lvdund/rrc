package ies

import "rrc/utils"

// BandNR-aperiodicCSI-RS-AdditionalBandwidth-r17 ::= ENUMERATED
type BandnrAperiodiccsiRsAdditionalbandwidthR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrAperiodiccsiRsAdditionalbandwidthR17EnumeratedNothing = iota
	BandnrAperiodiccsiRsAdditionalbandwidthR17EnumeratedAddbw_Set1
	BandnrAperiodiccsiRsAdditionalbandwidthR17EnumeratedAddbw_Set2
)
