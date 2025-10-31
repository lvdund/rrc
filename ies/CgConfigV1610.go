package ies

// CG-Config-v1610-IEs ::= SEQUENCE
type CgConfigV1610 struct {
	DrxInfoscg2          *DrxInfo2
	Noncriticalextension *CgConfigV1620
}
