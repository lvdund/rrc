package ies

// PUCCH-ConfigDedicated-v1020-pucch-Format-r10-channelSelection-r10-n1PUCCH-AN-CS-r10-setup ::= SEQUENCE
type PucchConfigdedicatedV1020PucchFormatR10ChannelselectionR10N1pucchAnCsR10Setup struct {
	N1pucchAnCsListR10 []N1pucchAnCsR10 `lb:1,ub:2`
}
