package ies

import "rrc/utils"

// PDSCH-Config-prb-BundlingType-dynamicBundling-bundleSizeSet2 ::= ENUMERATED
type PdschConfigPrbBundlingtypeDynamicbundlingBundlesizeset2 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPrbBundlingtypeDynamicbundlingBundlesizeset2EnumeratedNothing = iota
	PdschConfigPrbBundlingtypeDynamicbundlingBundlesizeset2EnumeratedN4
	PdschConfigPrbBundlingtypeDynamicbundlingBundlesizeset2EnumeratedWideband
)
