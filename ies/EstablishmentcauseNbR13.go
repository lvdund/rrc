package ies

import "rrc/utils"

// EstablishmentCause-NB-r13 ::= ENUMERATED
type EstablishmentcauseNbR13 struct {
	Value utils.ENUMERATED
}

const (
	EstablishmentcauseNbR13MtAccess                 = 0
	EstablishmentcauseNbR13MoSignalling             = 1
	EstablishmentcauseNbR13MoData                   = 2
	EstablishmentcauseNbR13MoExceptiondata          = 3
	EstablishmentcauseNbR13DelaytolerantaccessV1330 = 4
	EstablishmentcauseNbR13MtEdtV1610               = 5
	EstablishmentcauseNbR13Spare2                   = 6
	EstablishmentcauseNbR13Spare1                   = 7
)
