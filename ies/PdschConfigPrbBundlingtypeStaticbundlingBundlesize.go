package ies

import "rrc/utils"

// PDSCH-Config-prb-BundlingType-staticBundling-bundleSize ::= ENUMERATED
type PdschConfigPrbBundlingtypeStaticbundlingBundlesize struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPrbBundlingtypeStaticbundlingBundlesizeEnumeratedNothing = iota
	PdschConfigPrbBundlingtypeStaticbundlingBundlesizeEnumeratedN4
	PdschConfigPrbBundlingtypeStaticbundlingBundlesizeEnumeratedWideband
)
