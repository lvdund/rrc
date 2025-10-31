package ies

// BT-NameListConfig-r15 ::= CHOICE
const (
	BtNamelistconfigR15ChoiceNothing = iota
	BtNamelistconfigR15ChoiceRelease
	BtNamelistconfigR15ChoiceSetup
)

type BtNamelistconfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *BtNamelistR15
}
