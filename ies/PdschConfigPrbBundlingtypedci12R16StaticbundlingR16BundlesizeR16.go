package ies

import "rrc/utils"

// PDSCH-Config-prb-BundlingTypeDCI-1-2-r16-staticBundling-r16-bundleSize-r16 ::= ENUMERATED
type PdschConfigPrbBundlingtypedci12R16StaticbundlingR16BundlesizeR16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigPrbBundlingtypedci12R16StaticbundlingR16BundlesizeR16EnumeratedNothing = iota
	PdschConfigPrbBundlingtypedci12R16StaticbundlingR16BundlesizeR16EnumeratedN4
	PdschConfigPrbBundlingtypedci12R16StaticbundlingR16BundlesizeR16EnumeratedWideband
)
