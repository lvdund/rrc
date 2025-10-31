package ies

// SL-DRX-Config-r17 ::= SEQUENCE
// Extensible
type SlDrxConfigR17 struct {
	SlDrxConfiggcBcR17            *SlDrxConfiggcBcR17
	SlDrxConfigucToreleaselistR17 *[]SlDestinationindexR16 `lb:1,ub:maxNrofSLDestR16`
	SlDrxConfigucToaddmodlistR17  *[]SlDrxConfigucInfoR17  `lb:1,ub:maxNrofSLDestR16`
}
