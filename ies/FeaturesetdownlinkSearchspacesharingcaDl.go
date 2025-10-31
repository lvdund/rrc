package ies

import "rrc/utils"

// FeatureSetDownlink-searchSpaceSharingCA-DL ::= ENUMERATED
type FeaturesetdownlinkSearchspacesharingcaDl struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkSearchspacesharingcaDlEnumeratedNothing = iota
	FeaturesetdownlinkSearchspacesharingcaDlEnumeratedSupported
)
