package ies

// SetupReleasePurNrsrpChangethresholdR16 ::= SEQUENCE
type SetupreleasePurNrsrpChangethresholdR16 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PurNrsrpChangethresholdNbR16
}
