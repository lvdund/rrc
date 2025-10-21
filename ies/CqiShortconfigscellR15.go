package ies

import "rrc/utils"

// CQI-ShortConfigSCell-r15 ::= CHOICE
type CqiShortconfigscellR15 interface {
	isCqiShortconfigscellR15()
}

type CqiShortconfigscellR15Release struct {
	Value struct{}
}

func (*CqiShortconfigscellR15Release) isCqiShortconfigscellR15() {}

type CqiShortconfigscellR15Setup struct {
	Value interface{}
}

func (*CqiShortconfigscellR15Setup) isCqiShortconfigscellR15() {}
