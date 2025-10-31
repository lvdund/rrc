package ies

// DRX-Info ::= SEQUENCE
type DrxInfo struct {
	DrxLongcyclestartoffset DrxInfoDrxLongcyclestartoffset
	Shortdrx                *DrxInfoShortdrx
}
