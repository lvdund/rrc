package ies

// PDSCH-Config-prb-BundlingType ::= CHOICE
const (
	PdschConfigPrbBundlingtypeChoiceNothing = iota
	PdschConfigPrbBundlingtypeChoiceStaticbundling
	PdschConfigPrbBundlingtypeChoiceDynamicbundling
)

type PdschConfigPrbBundlingtype struct {
	Choice          uint64
	Staticbundling  *PdschConfigPrbBundlingtypeStaticbundling
	Dynamicbundling *PdschConfigPrbBundlingtypeDynamicbundling
}
