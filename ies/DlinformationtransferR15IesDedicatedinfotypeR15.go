package ies

// DLInformationTransfer-r15-IEs-dedicatedInfoType-r15 ::= CHOICE
const (
	DlinformationtransferR15IesDedicatedinfotypeR15ChoiceNothing = iota
	DlinformationtransferR15IesDedicatedinfotypeR15ChoiceDedicatedinfonas
	DlinformationtransferR15IesDedicatedinfotypeR15ChoiceDedicatedinfocdma20001xrtt
	DlinformationtransferR15IesDedicatedinfotypeR15ChoiceDedicatedinfocdma2000Hrpd
)

type DlinformationtransferR15IesDedicatedinfotypeR15 struct {
	Choice                     uint64
	Dedicatedinfonas           *Dedicatedinfonas
	Dedicatedinfocdma20001xrtt *Dedicatedinfocdma2000
	Dedicatedinfocdma2000Hrpd  *Dedicatedinfocdma2000
}
