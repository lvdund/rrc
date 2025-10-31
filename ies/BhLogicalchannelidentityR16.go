package ies

// BH-LogicalChannelIdentity-r16 ::= CHOICE
const (
	BhLogicalchannelidentityR16ChoiceNothing = iota
	BhLogicalchannelidentityR16ChoiceBhLogicalchannelidentityR16
	BhLogicalchannelidentityR16ChoiceBhLogicalchannelidentityextR16
)

type BhLogicalchannelidentityR16 struct {
	Choice                         uint64
	BhLogicalchannelidentityR16    *Logicalchannelidentity
	BhLogicalchannelidentityextR16 *BhLogicalchannelidentityExtR16
}
