package ies

// EIMTA-MainConfigServCell-r12-setup-mbsfn-SubframeConfigList-v1250 ::= CHOICE
const (
	EimtaMainconfigservcellR12SetupMbsfnSubframeconfiglistV1250ChoiceNothing = iota
	EimtaMainconfigservcellR12SetupMbsfnSubframeconfiglistV1250ChoiceRelease
	EimtaMainconfigservcellR12SetupMbsfnSubframeconfiglistV1250ChoiceSetup
)

type EimtaMainconfigservcellR12SetupMbsfnSubframeconfiglistV1250 struct {
	Choice  uint64
	Release *struct{}
	Setup   *EimtaMainconfigservcellR12SetupMbsfnSubframeconfiglistV1250Setup
}
