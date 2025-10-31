package ies

// SoundingRS-UL-ConfigDedicatedAperiodic-v1310 ::= CHOICE
const (
	SoundingrsUlConfigdedicatedaperiodicV1310ChoiceNothing = iota
	SoundingrsUlConfigdedicatedaperiodicV1310ChoiceRelease
	SoundingrsUlConfigdedicatedaperiodicV1310ChoiceSetup
)

type SoundingrsUlConfigdedicatedaperiodicV1310 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedaperiodicV1310Setup
}
