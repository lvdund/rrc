package ies

// RLC-BearerConfig-servedRadioBearer ::= CHOICE
const (
	RlcBearerconfigServedradiobearerChoiceNothing = iota
	RlcBearerconfigServedradiobearerChoiceSrbIdentity
	RlcBearerconfigServedradiobearerChoiceDrbIdentity
)

type RlcBearerconfigServedradiobearer struct {
	Choice      uint64
	SrbIdentity *SrbIdentity
	DrbIdentity *DrbIdentity
}
