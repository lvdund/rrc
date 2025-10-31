package ies

// SRS-Resource-srs-TCI-State-r17 ::= CHOICE
const (
	SrsResourceSrsTciStateR17ChoiceNothing = iota
	SrsResourceSrsTciStateR17ChoiceSrsUlTciState
	SrsResourceSrsTciStateR17ChoiceSrsDlorjointtciState
)

type SrsResourceSrsTciStateR17 struct {
	Choice               uint64
	SrsUlTciState        *TciUlStateIdR17
	SrsDlorjointtciState *TciStateid
}
