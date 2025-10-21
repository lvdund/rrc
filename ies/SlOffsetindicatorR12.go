package ies

import "rrc/utils"

// SL-OffsetIndicator-r12 ::= CHOICE
type SlOffsetindicatorR12 interface {
	isSlOffsetindicatorR12()
}

type SlOffsetindicatorR12SmallR12 struct {
	Value utils.INTEGER
}

func (*SlOffsetindicatorR12SmallR12) isSlOffsetindicatorR12() {}

type SlOffsetindicatorR12LargeR12 struct {
	Value utils.INTEGER
}

func (*SlOffsetindicatorR12LargeR12) isSlOffsetindicatorR12() {}
