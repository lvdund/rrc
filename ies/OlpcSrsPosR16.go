package ies

// OLPC-SRS-Pos-r16 ::= SEQUENCE
type OlpcSrsPosR16 struct {
	OlpcSrsPosbasedonprsServingR16         *OlpcSrsPosR16OlpcSrsPosbasedonprsServingR16
	OlpcSrsPosbasedonssbNeighR16           *OlpcSrsPosR16OlpcSrsPosbasedonssbNeighR16
	OlpcSrsPosbasedonprsNeighR16           *OlpcSrsPosR16OlpcSrsPosbasedonprsNeighR16
	MaxnumberpathlossestimateperservingR16 *OlpcSrsPosR16MaxnumberpathlossestimateperservingR16
}
