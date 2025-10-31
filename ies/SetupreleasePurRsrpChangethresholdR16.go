package ies

// SetupReleasePurRsrpChangethresholdR16 ::= SEQUENCE
type SetupreleasePurRsrpChangethresholdR16 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PurRsrpChangethresholdR16
}
