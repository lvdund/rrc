package ies

// SoundingRS-UL-ConfigDedicated ::= CHOICE
const (
	SoundingrsUlConfigdedicatedChoiceNothing = iota
	SoundingrsUlConfigdedicatedChoiceRelease
	SoundingrsUlConfigdedicatedChoiceSetup
)

type SoundingrsUlConfigdedicated struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedSetup
}
