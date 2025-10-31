package ies

import "rrc/utils"

// FeatureSetDownlink-v1730-prs-AsSpatialRelationRS-For-SRS-r17 ::= ENUMERATED
type FeaturesetdownlinkV1730PrsAsspatialrelationrsForSrsR17 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1730PrsAsspatialrelationrsForSrsR17EnumeratedNothing = iota
	FeaturesetdownlinkV1730PrsAsspatialrelationrsForSrsR17EnumeratedSupported
)
