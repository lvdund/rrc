package ies

import "rrc/utils"

// UE-TxTEG-Association-r17 ::= SEQUENCE
type UeTxtegAssociationR17 struct {
	UeTxtegIdR17                      utils.INTEGER `lb:0,ub:maxNrOfTxTEGId1R17`
	NrTimestampR17                    NrTimestampR17
	AssociatedsrsPosresourceidlistR17 []SrsPosresourceidR16 `lb:1,ub:maxNrofSRSPosresourcesR16`
	ServcellidR17                     *Servcellindex
}
