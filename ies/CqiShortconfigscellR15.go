package ies

// CQI-ShortConfigSCell-r15 ::= CHOICE
const (
	CqiShortconfigscellR15ChoiceNothing = iota
	CqiShortconfigscellR15ChoiceRelease
	CqiShortconfigscellR15ChoiceSetup
)

type CqiShortconfigscellR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiShortconfigscellR15Setup
}
