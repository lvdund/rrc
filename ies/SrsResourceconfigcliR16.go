package ies

// SRS-ResourceConfigCLI-r16 ::= SEQUENCE
// Extensible
type SrsResourceconfigcliR16 struct {
	SrsResourceR16      SrsResource
	SrsScsR16           Subcarrierspacing
	RefservcellindexR16 *Servcellindex
	RefbwpR16           BwpId
}
