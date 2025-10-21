package ies

import "rrc/utils"

// RCLWI-Configuration-r13 ::= CHOICE
type RclwiConfigurationR13 interface {
	isRclwiConfigurationR13()
}

type RclwiConfigurationR13Release struct {
	Value struct{}
}

func (*RclwiConfigurationR13Release) isRclwiConfigurationR13() {}

type RclwiConfigurationR13Setup struct {
	Value interface{}
}

func (*RclwiConfigurationR13Setup) isRclwiConfigurationR13() {}
