package ies

// PUCCH-CSI-Resource ::= SEQUENCE
type PucchCsiResource struct {
	Uplinkbandwidthpartid BwpId
	PucchResource         PucchResourceid
}
