package ies

// PDSCH-Config-prb-BundlingTypeDCI-1-2-r16 ::= CHOICE
const (
	PdschConfigPrbBundlingtypedci12R16ChoiceNothing = iota
	PdschConfigPrbBundlingtypedci12R16ChoiceStaticbundlingR16
	PdschConfigPrbBundlingtypedci12R16ChoiceDynamicbundlingR16
)

type PdschConfigPrbBundlingtypedci12R16 struct {
	Choice             uint64
	StaticbundlingR16  *PdschConfigPrbBundlingtypedci12R16StaticbundlingR16
	DynamicbundlingR16 *PdschConfigPrbBundlingtypedci12R16DynamicbundlingR16
}
