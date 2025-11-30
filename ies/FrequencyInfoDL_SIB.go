package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyInfoDL_SIB struct {
	FrequencyBandList       MultiFrequencyBandListNR_SIB `madatory`
	OffsetToPointA          int64                        `lb:0,ub:2199,madatory`
	Scs_SpecificCarrierList []SCS_SpecificCarrier        `lb:1,ub:maxSCSs,madatory`
}

func (ie *FrequencyInfoDL_SIB) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.FrequencyBandList.Encode(w); err != nil {
		return utils.WrapError("Encode FrequencyBandList", err)
	}
	if err = w.WriteInteger(ie.OffsetToPointA, &aper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("WriteInteger OffsetToPointA", err)
	}
	tmp_Scs_SpecificCarrierList := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, aper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	for _, i := range ie.Scs_SpecificCarrierList {
		tmp_Scs_SpecificCarrierList.Value = append(tmp_Scs_SpecificCarrierList.Value, &i)
	}
	if err = tmp_Scs_SpecificCarrierList.Encode(w); err != nil {
		return utils.WrapError("Encode Scs_SpecificCarrierList", err)
	}
	return nil
}

func (ie *FrequencyInfoDL_SIB) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.FrequencyBandList.Decode(r); err != nil {
		return utils.WrapError("Decode FrequencyBandList", err)
	}
	var tmp_int_OffsetToPointA int64
	if tmp_int_OffsetToPointA, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("ReadInteger OffsetToPointA", err)
	}
	ie.OffsetToPointA = tmp_int_OffsetToPointA
	tmp_Scs_SpecificCarrierList := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, aper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	fn_Scs_SpecificCarrierList := func() *SCS_SpecificCarrier {
		return new(SCS_SpecificCarrier)
	}
	if err = tmp_Scs_SpecificCarrierList.Decode(r, fn_Scs_SpecificCarrierList); err != nil {
		return utils.WrapError("Decode Scs_SpecificCarrierList", err)
	}
	ie.Scs_SpecificCarrierList = []SCS_SpecificCarrier{}
	for _, i := range tmp_Scs_SpecificCarrierList.Value {
		ie.Scs_SpecificCarrierList = append(ie.Scs_SpecificCarrierList, *i)
	}
	return nil
}
