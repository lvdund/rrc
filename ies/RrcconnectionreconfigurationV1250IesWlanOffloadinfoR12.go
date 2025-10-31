package ies

// RRCConnectionReconfiguration-v1250-IEs-wlan-OffloadInfo-r12 ::= CHOICE
const (
	RrcconnectionreconfigurationV1250IesWlanOffloadinfoR12ChoiceNothing = iota
	RrcconnectionreconfigurationV1250IesWlanOffloadinfoR12ChoiceRelease
	RrcconnectionreconfigurationV1250IesWlanOffloadinfoR12ChoiceSetup
)

type RrcconnectionreconfigurationV1250IesWlanOffloadinfoR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RrcconnectionreconfigurationV1250IesWlanOffloadinfoR12Setup
}
