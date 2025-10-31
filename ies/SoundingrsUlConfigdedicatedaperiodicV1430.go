package ies

// SoundingRS-UL-ConfigDedicatedAperiodic-v1430 ::= CHOICE
const (
	SoundingrsUlConfigdedicatedaperiodicV1430ChoiceNothing = iota
	SoundingrsUlConfigdedicatedaperiodicV1430ChoiceRelease
	SoundingrsUlConfigdedicatedaperiodicV1430ChoiceSetup
)

type SoundingrsUlConfigdedicatedaperiodicV1430 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedaperiodicV1430Setup
}
