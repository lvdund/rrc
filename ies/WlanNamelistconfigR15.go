package ies

// WLAN-NameListConfig-r15 ::= CHOICE
const (
	WlanNamelistconfigR15ChoiceNothing = iota
	WlanNamelistconfigR15ChoiceRelease
	WlanNamelistconfigR15ChoiceSetup
)

type WlanNamelistconfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *WlanNamelistR15
}
