package ies

import "rrc/utils"

// UuMessageTransferSidelink-r17-IEs ::= SEQUENCE
type UumessagetransfersidelinkR17 struct {
	SlPagingdeliveryR17            *utils.OCTETSTRING
	SlSib1DeliveryR17              *utils.OCTETSTRING
	SlSysteminformationdeliveryR17 *utils.OCTETSTRING
	Latenoncriticalextension       *utils.OCTETSTRING
	Noncriticalextension           *UumessagetransfersidelinkR17IesNoncriticalextension
}
