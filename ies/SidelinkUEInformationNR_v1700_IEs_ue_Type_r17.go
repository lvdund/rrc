package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SidelinkUEInformationNR_v1700_IEs_ue_Type_r17_Enum_relayUE  aper.Enumerated = 0
	SidelinkUEInformationNR_v1700_IEs_ue_Type_r17_Enum_remoteUE aper.Enumerated = 1
)

type SidelinkUEInformationNR_v1700_IEs_ue_Type_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SidelinkUEInformationNR_v1700_IEs_ue_Type_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SidelinkUEInformationNR_v1700_IEs_ue_Type_r17", err)
	}
	return nil
}

func (ie *SidelinkUEInformationNR_v1700_IEs_ue_Type_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SidelinkUEInformationNR_v1700_IEs_ue_Type_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
