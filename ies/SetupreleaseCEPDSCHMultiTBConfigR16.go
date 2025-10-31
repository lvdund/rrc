package ies

// SetupReleaseCEPDSCHMultiTBConfigR16 ::= SEQUENCE
type SetupreleaseCEPDSCHMultiTBConfigR16 struct {
	Choice uint64
	Release *struct{}
	Setup *CePdschMultitbConfigR16
}