package ies

// SoundingRS-UL-ConfigDedicatedAperiodic-r10 ::= CHOICE
// Extensible
const (
	SoundingrsUlConfigdedicatedaperiodicR10ChoiceNothing = iota
	SoundingrsUlConfigdedicatedaperiodicR10ChoiceRelease
	SoundingrsUlConfigdedicatedaperiodicR10ChoiceSetup
)

type SoundingrsUlConfigdedicatedaperiodicR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedaperiodicR10Setup
}
