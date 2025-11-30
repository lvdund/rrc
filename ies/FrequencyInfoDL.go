package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyInfoDL struct {
	AbsoluteFrequencySSB    *ARFCN_ValueNR           `optional`
	FrequencyBandList       MultiFrequencyBandListNR `madatory`
	AbsoluteFrequencyPointA ARFCN_ValueNR            `madatory`
	Scs_SpecificCarrierList []SCS_SpecificCarrier    `lb:1,ub:maxSCSs,madatory`
}

func (ie *FrequencyInfoDL) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AbsoluteFrequencySSB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AbsoluteFrequencySSB != nil {
		if err = ie.AbsoluteFrequencySSB.Encode(w); err != nil {
			return utils.WrapError("Encode AbsoluteFrequencySSB", err)
		}
	}
	if err = ie.FrequencyBandList.Encode(w); err != nil {
		return utils.WrapError("Encode FrequencyBandList", err)
	}
	if err = ie.AbsoluteFrequencyPointA.Encode(w); err != nil {
		return utils.WrapError("Encode AbsoluteFrequencyPointA", err)
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

func (ie *FrequencyInfoDL) Decode(r *aper.AperReader) error {
	var err error
	var AbsoluteFrequencySSBPresent bool
	if AbsoluteFrequencySSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if AbsoluteFrequencySSBPresent {
		ie.AbsoluteFrequencySSB = new(ARFCN_ValueNR)
		if err = ie.AbsoluteFrequencySSB.Decode(r); err != nil {
			return utils.WrapError("Decode AbsoluteFrequencySSB", err)
		}
	}
	if err = ie.FrequencyBandList.Decode(r); err != nil {
		return utils.WrapError("Decode FrequencyBandList", err)
	}
	if err = ie.AbsoluteFrequencyPointA.Decode(r); err != nil {
		return utils.WrapError("Decode AbsoluteFrequencyPointA", err)
	}
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
