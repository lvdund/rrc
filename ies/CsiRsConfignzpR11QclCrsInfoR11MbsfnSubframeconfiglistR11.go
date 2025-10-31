package ies

// CSI-RS-ConfigNZP-r11-qcl-CRS-Info-r11-mbsfn-SubframeConfigList-r11 ::= CHOICE
const (
	CsiRsConfignzpR11QclCrsInfoR11MbsfnSubframeconfiglistR11ChoiceNothing = iota
	CsiRsConfignzpR11QclCrsInfoR11MbsfnSubframeconfiglistR11ChoiceRelease
	CsiRsConfignzpR11QclCrsInfoR11MbsfnSubframeconfiglistR11ChoiceSetup
)

type CsiRsConfignzpR11QclCrsInfoR11MbsfnSubframeconfiglistR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CsiRsConfignzpR11QclCrsInfoR11MbsfnSubframeconfiglistR11Setup
}
