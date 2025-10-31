package ies

// SoundingRS-UL-ConfigDedicatedAperiodicUpPTsExt-r13 ::= CHOICE
const (
	SoundingrsUlConfigdedicatedaperiodicupptsextR13ChoiceNothing = iota
	SoundingrsUlConfigdedicatedaperiodicupptsextR13ChoiceRelease
	SoundingrsUlConfigdedicatedaperiodicupptsextR13ChoiceSetup
)

type SoundingrsUlConfigdedicatedaperiodicupptsextR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigdedicatedaperiodicupptsextR13Setup
}
