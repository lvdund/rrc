package ies

// SoundingRS-UL-ConfigDedicatedAdd-r16-srs-ActivateAp-r13 ::= CHOICE
const (
	SoundingrsUlConfigdedicatedaddR16SrsActivateapR13ChoiceNothing = iota
	SoundingrsUlConfigdedicatedaddR16SrsActivateapR13ChoiceRelease
	SoundingrsUlConfigdedicatedaddR16SrsActivateapR13ChoiceSetup
)

type SoundingrsUlConfigdedicatedaddR16SrsActivateapR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedaddR16SrsActivateapR13Setup
}
