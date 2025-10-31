package ies

// WUS-Config-r15 ::= SEQUENCE
type WusConfigR15 struct {
	MaxdurationfactorR15   WusConfigR15MaxdurationfactorR15
	NumposR15              WusConfigR15NumposR15
	FreqlocationR15        WusConfigR15FreqlocationR15
	TimeoffsetdrxR15       WusConfigR15TimeoffsetdrxR15
	TimeoffsetEdrxShortR15 WusConfigR15TimeoffsetEdrxShortR15
	TimeoffsetEdrxLongR15  *WusConfigR15TimeoffsetEdrxLongR15
}
