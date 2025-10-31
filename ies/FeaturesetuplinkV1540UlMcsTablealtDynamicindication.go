package ies

import "rrc/utils"

// FeatureSetUplink-v1540-ul-MCS-TableAlt-DynamicIndication ::= ENUMERATED
type FeaturesetuplinkV1540UlMcsTablealtDynamicindication struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1540UlMcsTablealtDynamicindicationEnumeratedNothing = iota
	FeaturesetuplinkV1540UlMcsTablealtDynamicindicationEnumeratedSupported
)
