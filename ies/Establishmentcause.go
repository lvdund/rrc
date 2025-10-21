package ies

import "rrc/utils"

// EstablishmentCause ::= ENUMERATED
type Establishmentcause struct {
	Value utils.ENUMERATED
}

const (
	EstablishmentcauseEmergency                = 0
	EstablishmentcauseHighpriorityaccess       = 1
	EstablishmentcauseMtAccess                 = 2
	EstablishmentcauseMoSignalling             = 3
	EstablishmentcauseMoData                   = 4
	EstablishmentcauseDelaytolerantaccessV1020 = 5
	EstablishmentcauseMoVoicecallV1280         = 6
	EstablishmentcauseSpare1                   = 7
)
