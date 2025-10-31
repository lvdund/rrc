package ies

// RCLWI-Config-r13-command ::= CHOICE
const (
	RclwiConfigR13CommandChoiceNothing = iota
	RclwiConfigR13CommandChoiceSteertowlanR13
	RclwiConfigR13CommandChoiceSteertolteR13
)

type RclwiConfigR13Command struct {
	Choice         uint64
	SteertowlanR13 *RclwiConfigR13CommandSteertowlanR13
	SteertolteR13  *struct{}
}
