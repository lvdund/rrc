package ies

// SetupReleaseCePuschMultitbConfigR16 ::= SEQUENCE
type SetupreleaseCePuschMultitbConfigR16 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CePuschMultitbConfigR16
}
