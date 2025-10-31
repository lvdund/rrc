package ies

// MAC-MainConfig-NB-r13-logicalChannelSR-Config-r13 ::= CHOICE
const (
	MacMainconfigNbR13LogicalchannelsrConfigR13ChoiceNothing = iota
	MacMainconfigNbR13LogicalchannelsrConfigR13ChoiceRelease
	MacMainconfigNbR13LogicalchannelsrConfigR13ChoiceSetup
)

type MacMainconfigNbR13LogicalchannelsrConfigR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *MacMainconfigNbR13LogicalchannelsrConfigR13Setup
}
