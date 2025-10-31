package ies

import "rrc/utils"

// DRB-ToAddMod-cnAssociation ::= CHOICE
const (
	DrbToaddmodCnassociationChoiceNothing = iota
	DrbToaddmodCnassociationChoiceEpsBeareridentity
	DrbToaddmodCnassociationChoiceSdapConfig
)

type DrbToaddmodCnassociation struct {
	Choice            uint64
	EpsBeareridentity *utils.INTEGER `lb:0,ub:15`
	SdapConfig        *SdapConfig
}
