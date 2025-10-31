package ies

import "rrc/utils"

// PDSCH-Config-prb-BundlingTypeDCI-1-2-r16-dynamicBundling-r16-bundleSizeSet2-r16 ::= ENUMERATED
type PdschConfigPrbBundlingtypedci12R16DynamicbundlingR16Bundlesizeset2R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPrbBundlingtypedci12R16DynamicbundlingR16Bundlesizeset2R16EnumeratedNothing = iota
	PdschConfigPrbBundlingtypedci12R16DynamicbundlingR16Bundlesizeset2R16EnumeratedN4
	PdschConfigPrbBundlingtypedci12R16DynamicbundlingR16Bundlesizeset2R16EnumeratedWideband
)
