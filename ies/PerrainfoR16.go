package ies

// PerRAInfo-r16 ::= CHOICE
const (
	PerrainfoR16ChoiceNothing = iota
	PerrainfoR16ChoicePerrassbinfolistR16
	PerrainfoR16ChoicePerracsiRsinfolistR16
)

type PerrainfoR16 struct {
	Choice                uint64
	PerrassbinfolistR16   *PerrassbinfoR16
	PerracsiRsinfolistR16 *PerracsiRsinfoR16
}
