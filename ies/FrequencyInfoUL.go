package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyInfoUL struct {
	FrequencyBandList          *MultiFrequencyBandListNR             `optional`
	AbsoluteFrequencyPointA    *ARFCN_ValueNR                        `optional`
	Scs_SpecificCarrierList    []SCS_SpecificCarrier                 `lb:1,ub:maxSCSs,madatory`
	AdditionalSpectrumEmission *AdditionalSpectrumEmission           `optional`
	P_Max                      *P_Max                                `optional`
	FrequencyShift7p5khz       *FrequencyInfoUL_frequencyShift7p5khz `optional`
}

func (ie *FrequencyInfoUL) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FrequencyBandList != nil, ie.AbsoluteFrequencyPointA != nil, ie.AdditionalSpectrumEmission != nil, ie.P_Max != nil, ie.FrequencyShift7p5khz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FrequencyBandList != nil {
		if err = ie.FrequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyBandList", err)
		}
	}
	if ie.AbsoluteFrequencyPointA != nil {
		if err = ie.AbsoluteFrequencyPointA.Encode(w); err != nil {
			return utils.WrapError("Encode AbsoluteFrequencyPointA", err)
		}
	}
	tmp_Scs_SpecificCarrierList := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, aper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	for _, i := range ie.Scs_SpecificCarrierList {
		tmp_Scs_SpecificCarrierList.Value = append(tmp_Scs_SpecificCarrierList.Value, &i)
	}
	if err = tmp_Scs_SpecificCarrierList.Encode(w); err != nil {
		return utils.WrapError("Encode Scs_SpecificCarrierList", err)
	}
	if ie.AdditionalSpectrumEmission != nil {
		if err = ie.AdditionalSpectrumEmission.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalSpectrumEmission", err)
		}
	}
	if ie.P_Max != nil {
		if err = ie.P_Max.Encode(w); err != nil {
			return utils.WrapError("Encode P_Max", err)
		}
	}
	if ie.FrequencyShift7p5khz != nil {
		if err = ie.FrequencyShift7p5khz.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyShift7p5khz", err)
		}
	}
	return nil
}

func (ie *FrequencyInfoUL) Decode(r *aper.AperReader) error {
	var err error
	var FrequencyBandListPresent bool
	if FrequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsoluteFrequencyPointAPresent bool
	if AbsoluteFrequencyPointAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalSpectrumEmissionPresent bool
	if AdditionalSpectrumEmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P_MaxPresent bool
	if P_MaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyShift7p5khzPresent bool
	if FrequencyShift7p5khzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if FrequencyBandListPresent {
		ie.FrequencyBandList = new(MultiFrequencyBandListNR)
		if err = ie.FrequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyBandList", err)
		}
	}
	if AbsoluteFrequencyPointAPresent {
		ie.AbsoluteFrequencyPointA = new(ARFCN_ValueNR)
		if err = ie.AbsoluteFrequencyPointA.Decode(r); err != nil {
			return utils.WrapError("Decode AbsoluteFrequencyPointA", err)
		}
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
	if AdditionalSpectrumEmissionPresent {
		ie.AdditionalSpectrumEmission = new(AdditionalSpectrumEmission)
		if err = ie.AdditionalSpectrumEmission.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalSpectrumEmission", err)
		}
	}
	if P_MaxPresent {
		ie.P_Max = new(P_Max)
		if err = ie.P_Max.Decode(r); err != nil {
			return utils.WrapError("Decode P_Max", err)
		}
	}
	if FrequencyShift7p5khzPresent {
		ie.FrequencyShift7p5khz = new(FrequencyInfoUL_frequencyShift7p5khz)
		if err = ie.FrequencyShift7p5khz.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyShift7p5khz", err)
		}
	}
	return nil
}
