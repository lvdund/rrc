package ies

// SoundingRS-UL-ConfigDedicatedAperiodic-v1310-setup-srs-ActivateAp-v1310 ::= CHOICE
const (
	SoundingrsUlConfigdedicatedaperiodicV1310SetupSrsActivateapV1310ChoiceNothing = iota
	SoundingrsUlConfigdedicatedaperiodicV1310SetupSrsActivateapV1310ChoiceRelease
	SoundingrsUlConfigdedicatedaperiodicV1310SetupSrsActivateapV1310ChoiceSetup
)

type SoundingrsUlConfigdedicatedaperiodicV1310SetupSrsActivateapV1310 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedaperiodicV1310SetupSrsActivateapV1310Setup
}
