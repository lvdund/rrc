package ies

// EPDCCH-Config-r11-config-r11 ::= CHOICE
const (
	EpdcchConfigR11ConfigR11ChoiceNothing = iota
	EpdcchConfigR11ConfigR11ChoiceRelease
	EpdcchConfigR11ConfigR11ChoiceSetup
)

type EpdcchConfigR11ConfigR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *EpdcchConfigR11ConfigR11Setup
}
