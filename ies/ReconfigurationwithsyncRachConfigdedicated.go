package ies

// ReconfigurationWithSync-rach-ConfigDedicated ::= CHOICE
const (
	ReconfigurationwithsyncRachConfigdedicatedChoiceNothing = iota
	ReconfigurationwithsyncRachConfigdedicatedChoiceUplink
	ReconfigurationwithsyncRachConfigdedicatedChoiceSupplementaryuplink
)

type ReconfigurationwithsyncRachConfigdedicated struct {
	Choice              uint64
	Uplink              *RachConfigdedicated
	Supplementaryuplink *RachConfigdedicated
}
