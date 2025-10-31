package ies

// RMTC-Config-r13 ::= CHOICE
// Extensible
const (
	RmtcConfigR13ChoiceNothing = iota
	RmtcConfigR13ChoiceRelease
	RmtcConfigR13ChoiceSetup
)

type RmtcConfigR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RmtcConfigR13Setup
}
