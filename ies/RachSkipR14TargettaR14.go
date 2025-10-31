package ies

// RACH-Skip-r14-targetTA-r14 ::= CHOICE
const (
	RachSkipR14TargettaR14ChoiceNothing = iota
	RachSkipR14TargettaR14ChoiceTa0R14
	RachSkipR14TargettaR14ChoiceMcgPtagR14
	RachSkipR14TargettaR14ChoiceScgPtagR14
	RachSkipR14TargettaR14ChoiceMcgStagR14
	RachSkipR14TargettaR14ChoiceScgStagR14
)

type RachSkipR14TargettaR14 struct {
	Choice     uint64
	Ta0R14     *struct{}
	McgPtagR14 *struct{}
	ScgPtagR14 *struct{}
	McgStagR14 *StagIdR11
	ScgStagR14 *StagIdR11
}
