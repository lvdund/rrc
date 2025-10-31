package ies

// SoundingRS-UL-ConfigDedicatedAperiodic-r10-setup-srs-ActivateAp-r10 ::= CHOICE
// Extensible
const (
	SoundingrsUlConfigdedicatedaperiodicR10SetupSrsActivateapR10ChoiceNothing = iota
	SoundingrsUlConfigdedicatedaperiodicR10SetupSrsActivateapR10ChoiceRelease
	SoundingrsUlConfigdedicatedaperiodicR10SetupSrsActivateapR10ChoiceSetup
)

type SoundingrsUlConfigdedicatedaperiodicR10SetupSrsActivateapR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedaperiodicR10SetupSrsActivateapR10Setup
}
