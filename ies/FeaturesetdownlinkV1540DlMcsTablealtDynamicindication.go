package ies

import "rrc/utils"

// FeatureSetDownlink-v1540-dl-MCS-TableAlt-DynamicIndication ::= ENUMERATED
type FeaturesetdownlinkV1540DlMcsTablealtDynamicindication struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkV1540DlMcsTablealtDynamicindicationEnumeratedNothing = iota
	FeaturesetdownlinkV1540DlMcsTablealtDynamicindicationEnumeratedSupported
)
