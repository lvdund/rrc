package ies

// BWP-UplinkDedicated-ul-TCI-StateList-r17 ::= CHOICE
const (
	BwpUplinkdedicatedUlTciStatelistR17ChoiceNothing = iota
	BwpUplinkdedicatedUlTciStatelistR17ChoiceExplicitlist
	BwpUplinkdedicatedUlTciStatelistR17ChoiceUnifiedtciStaterefR17
)

type BwpUplinkdedicatedUlTciStatelistR17 struct {
	Choice                uint64
	Explicitlist          *BwpUplinkdedicatedUlTciStatelistR17Explicitlist
	UnifiedtciStaterefR17 *ServingcellandbwpIdR17
}
