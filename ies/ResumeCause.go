package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ResumeCause_Enum_emergency          aper.Enumerated = 0
	ResumeCause_Enum_highPriorityAccess aper.Enumerated = 1
	ResumeCause_Enum_mt_Access          aper.Enumerated = 2
	ResumeCause_Enum_mo_Signalling      aper.Enumerated = 3
	ResumeCause_Enum_mo_Data            aper.Enumerated = 4
	ResumeCause_Enum_mo_VoiceCall       aper.Enumerated = 5
	ResumeCause_Enum_mo_VideoCall       aper.Enumerated = 6
	ResumeCause_Enum_mo_SMS             aper.Enumerated = 7
	ResumeCause_Enum_rna_Update         aper.Enumerated = 8
	ResumeCause_Enum_mps_PriorityAccess aper.Enumerated = 9
	ResumeCause_Enum_mcs_PriorityAccess aper.Enumerated = 10
	ResumeCause_Enum_spare1             aper.Enumerated = 11
	ResumeCause_Enum_spare2             aper.Enumerated = 12
	ResumeCause_Enum_spare3             aper.Enumerated = 13
	ResumeCause_Enum_spare4             aper.Enumerated = 14
	ResumeCause_Enum_spare5             aper.Enumerated = 15
)

type ResumeCause struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *ResumeCause) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode ResumeCause", err)
	}
	return nil
}

func (ie *ResumeCause) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode ResumeCause", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
