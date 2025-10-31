package ies

import "rrc/utils"

// FeatureSetDownlink-csi-RS-MeasSCellWithoutSSB ::= ENUMERATED
type FeaturesetdownlinkCsiRsMeasscellwithoutssb struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkCsiRsMeasscellwithoutssbEnumeratedNothing = iota
	FeaturesetdownlinkCsiRsMeasscellwithoutssbEnumeratedSupported
)
