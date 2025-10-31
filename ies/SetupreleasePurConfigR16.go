package ies

// SetupReleasePurConfigR16 ::= SEQUENCE
type SetupreleasePurConfigR16 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PurConfigR16
}
