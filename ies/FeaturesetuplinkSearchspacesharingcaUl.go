package ies

import "rrc/utils"

// FeatureSetUplink-searchSpaceSharingCA-UL ::= ENUMERATED
type FeaturesetuplinkSearchspacesharingcaUl struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkSearchspacesharingcaUlEnumeratedNothing = iota
	FeaturesetuplinkSearchspacesharingcaUlEnumeratedSupported
)
