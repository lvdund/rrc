package ies

// SoundingRS-UL-ConfigDedicatedUpPTsExt-r13 ::= CHOICE
const (
	SoundingrsUlConfigdedicatedupptsextR13ChoiceNothing = iota
	SoundingrsUlConfigdedicatedupptsextR13ChoiceRelease
	SoundingrsUlConfigdedicatedupptsextR13ChoiceSetup
)

type SoundingrsUlConfigdedicatedupptsextR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedupptsextR13Setup
}
