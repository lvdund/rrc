package ies

// BandParameters-v14b0 ::= SEQUENCE
type BandparametersV14b0 struct {
	SrsCapabilityperbandpairlistV14b0 *[]SrsCapabilityperbandpairV14b0 `lb:1,ub:maxSimultaneousBandsR10`
}
