package ies

// SoundingRS-UL-ConfigDedicated-v1310 ::= CHOICE
const (
	SoundingrsUlConfigdedicatedV1310ChoiceNothing = iota
	SoundingrsUlConfigdedicatedV1310ChoiceRelease
	SoundingrsUlConfigdedicatedV1310ChoiceSetup
)

type SoundingrsUlConfigdedicatedV1310 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedV1310Setup
}
