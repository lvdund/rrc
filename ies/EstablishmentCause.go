package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	EstablishmentCause_Enum_emergency          aper.Enumerated = 0
	EstablishmentCause_Enum_highPriorityAccess aper.Enumerated = 1
	EstablishmentCause_Enum_mt_Access          aper.Enumerated = 2
	EstablishmentCause_Enum_mo_Signalling      aper.Enumerated = 3
	EstablishmentCause_Enum_mo_Data            aper.Enumerated = 4
	EstablishmentCause_Enum_mo_VoiceCall       aper.Enumerated = 5
	EstablishmentCause_Enum_mo_VideoCall       aper.Enumerated = 6
	EstablishmentCause_Enum_mo_SMS             aper.Enumerated = 7
	EstablishmentCause_Enum_mps_PriorityAccess aper.Enumerated = 8
	EstablishmentCause_Enum_mcs_PriorityAccess aper.Enumerated = 9
	EstablishmentCause_Enum_spare6             aper.Enumerated = 10
	EstablishmentCause_Enum_spare5             aper.Enumerated = 11
	EstablishmentCause_Enum_spare4             aper.Enumerated = 12
	EstablishmentCause_Enum_spare3             aper.Enumerated = 13
	EstablishmentCause_Enum_spare2             aper.Enumerated = 14
	EstablishmentCause_Enum_spare1             aper.Enumerated = 15
)

type EstablishmentCause struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *EstablishmentCause) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode EstablishmentCause", err)
	}
	return nil
}

func (ie *EstablishmentCause) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode EstablishmentCause", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
