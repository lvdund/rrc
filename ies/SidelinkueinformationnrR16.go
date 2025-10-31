package ies

import "rrc/utils"

// SidelinkUEInformationNR-r16-IEs ::= SEQUENCE
type SidelinkueinformationnrR16 struct {
	SlRxinterestedfreqlistR16 *SlInterestedfreqlistR16
	SlTxresourcereqlistR16    *SlTxresourcereqlistR16
	SlFailurelistR16          *SlFailurelistR16
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *SidelinkueinformationnrV1700
}
