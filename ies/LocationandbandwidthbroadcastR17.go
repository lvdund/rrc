package ies

import "rrc/utils"

// LocationAndBandwidthBroadcast-r17 ::= CHOICE
const (
	LocationandbandwidthbroadcastR17ChoiceNothing = iota
	LocationandbandwidthbroadcastR17ChoiceSameassib1configuredlocationandbw
	LocationandbandwidthbroadcastR17ChoiceLocationandbandwidth
)

type LocationandbandwidthbroadcastR17 struct {
	Choice                            uint64
	Sameassib1configuredlocationandbw *struct{}
	Locationandbandwidth              *utils.INTEGER `lb:0,ub:37949`
}
