package ies

import "rrc/utils"

// ResumeCause ::= ENUMERATED
type Resumecause struct {
	Value utils.ENUMERATED
}

const (
	ResumecauseEmergency                = 0
	ResumecauseHighpriorityaccess       = 1
	ResumecauseMtAccess                 = 2
	ResumecauseMoSignalling             = 3
	ResumecauseMoData                   = 4
	ResumecauseDelaytolerantaccessV1020 = 5
	ResumecauseMoVoicecallV1280         = 6
	ResumecauseMtEdtV1610               = 7
)
