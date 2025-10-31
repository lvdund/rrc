package ies

// UTRA-FDD-Parameters-r16 ::= SEQUENCE
// Extensible
type UtraFddParametersR16 struct {
	SupportedbandlistutraFddR16 []SupportedbandutraFddR16 `lb:1,ub:maxBandsUTRAFddR16`
}
