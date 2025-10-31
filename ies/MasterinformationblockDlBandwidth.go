package ies

import "rrc/utils"

// MasterInformationBlock-dl-Bandwidth ::= ENUMERATED
type MasterinformationblockDlBandwidth struct {
	Value utils.ENUMERATED
}

const (
	MasterinformationblockDlBandwidthEnumeratedNothing = iota
	MasterinformationblockDlBandwidthEnumeratedN6
	MasterinformationblockDlBandwidthEnumeratedN15
	MasterinformationblockDlBandwidthEnumeratedN25
	MasterinformationblockDlBandwidthEnumeratedN50
	MasterinformationblockDlBandwidthEnumeratedN75
	MasterinformationblockDlBandwidthEnumeratedN100
)
