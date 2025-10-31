package ies

// CSI-RS-Config-r10-csi-RS-r10 ::= CHOICE
const (
	CsiRsConfigR10CsiRsR10ChoiceNothing = iota
	CsiRsConfigR10CsiRsR10ChoiceRelease
	CsiRsConfigR10CsiRsR10ChoiceSetup
)

type CsiRsConfigR10CsiRsR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfigR10CsiRsR10Setup
}
