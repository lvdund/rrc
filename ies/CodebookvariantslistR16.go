package ies

// CodebookVariantsList-r16 ::= SEQUENCE OF SupportedCSI-RS-Resource
// SIZE (1..maxNrofCSI-RS-ResourcesAlt-r16)
type CodebookvariantslistR16 struct {
	Value []SupportedcsiRsResource `lb:1,ub:maxNrofCSIRsResourcesaltR16`
}
