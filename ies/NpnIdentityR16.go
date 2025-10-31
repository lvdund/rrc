package ies

// NPN-Identity-r16 ::= CHOICE
const (
	NpnIdentityR16ChoiceNothing = iota
	NpnIdentityR16ChoicePniNpnR16
	NpnIdentityR16ChoiceSnpnR16
)

type NpnIdentityR16 struct {
	Choice    uint64
	PniNpnR16 *NpnIdentityR16PniNpnR16
	SnpnR16   *NpnIdentityR16SnpnR16
}
