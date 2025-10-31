package ies

// MAC-MainConfig-phr-Config ::= CHOICE
const (
	MacMainconfigPhrConfigChoiceNothing = iota
	MacMainconfigPhrConfigChoiceRelease
	MacMainconfigPhrConfigChoiceSetup
)

type MacMainconfigPhrConfig struct {
	Choice  uint64
	Release *struct{}
	Setup   *MacMainconfigPhrConfigSetup
}
