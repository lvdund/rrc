// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EstablishmentCause (line 1800).

const (
	EstablishmentCause_Emergency          = 0
	EstablishmentCause_HighPriorityAccess = 1
	EstablishmentCause_Mt_Access          = 2
	EstablishmentCause_Mo_Signalling      = 3
	EstablishmentCause_Mo_Data            = 4
	EstablishmentCause_Mo_VoiceCall       = 5
	EstablishmentCause_Mo_VideoCall       = 6
	EstablishmentCause_Mo_SMS             = 7
	EstablishmentCause_Mps_PriorityAccess = 8
	EstablishmentCause_Mcs_PriorityAccess = 9
	EstablishmentCause_Spare6             = 10
	EstablishmentCause_Spare5             = 11
	EstablishmentCause_Spare4             = 12
	EstablishmentCause_Spare3             = 13
	EstablishmentCause_Spare2             = 14
	EstablishmentCause_Spare1             = 15
)

var establishmentCauseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EstablishmentCause_Emergency, EstablishmentCause_HighPriorityAccess, EstablishmentCause_Mt_Access, EstablishmentCause_Mo_Signalling, EstablishmentCause_Mo_Data, EstablishmentCause_Mo_VoiceCall, EstablishmentCause_Mo_VideoCall, EstablishmentCause_Mo_SMS, EstablishmentCause_Mps_PriorityAccess, EstablishmentCause_Mcs_PriorityAccess, EstablishmentCause_Spare6, EstablishmentCause_Spare5, EstablishmentCause_Spare4, EstablishmentCause_Spare3, EstablishmentCause_Spare2, EstablishmentCause_Spare1},
}

type EstablishmentCause struct {
	Value int64
}

func (v *EstablishmentCause) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, establishmentCauseConstraints)
}

func (v *EstablishmentCause) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(establishmentCauseConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
