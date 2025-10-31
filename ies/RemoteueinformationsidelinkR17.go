package ies

import "rrc/utils"

// RemoteUEInformationSidelink-r17-IEs ::= SEQUENCE
type RemoteueinformationsidelinkR17 struct {
	SlRequestedsibListR17    *utils.Setuprelease[SlRequestedsibListR17]
	SlPaginginfoRemoteueR17  *utils.Setuprelease[SlPaginginfoRemoteueR17]
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RemoteueinformationsidelinkR17IesNoncriticalextension
}
