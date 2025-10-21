package ies

import "rrc/utils"

// SoundingRS-UL-ConfigCommon ::= CHOICE
type SoundingrsUlConfigcommon interface {
	isSoundingrsUlConfigcommon()
}

type SoundingrsUlConfigcommonRelease struct {
	Value struct{}
}

func (*SoundingrsUlConfigcommonRelease) isSoundingrsUlConfigcommon() {}

type SoundingrsUlConfigcommonSetup struct {
	Value interface{}
}

func (*SoundingrsUlConfigcommonSetup) isSoundingrsUlConfigcommon() {}
