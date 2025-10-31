package ies

// SoundingRS-UL-ConfigCommon ::= CHOICE
const (
	SoundingrsUlConfigcommonChoiceNothing = iota
	SoundingrsUlConfigcommonChoiceRelease
	SoundingrsUlConfigcommonChoiceSetup
)

type SoundingrsUlConfigcommon struct {
	Choice  uint64
	Release *struct{}
	Setup   *SoundingrsUlConfigcommonSetup
}
