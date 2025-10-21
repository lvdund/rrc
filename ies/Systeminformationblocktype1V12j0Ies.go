package ies

import "rrc/utils"

// SystemInformationBlockType1-v12j0-IEs ::= SEQUENCE
type Systeminformationblocktype1V12j0Ies struct {
	SchedulinginfolistV12j0  *SchedulinginfolistV12j0
	SchedulinginfolistextR12 *SchedulinginfolistextR12
	Noncriticalextension     *Systeminformationblocktype1V15g0Ies
}
