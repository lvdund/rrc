package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination struct {
	BandList                         []BandParameters                  `lb:1,ub:maxSimultaneousBands,madatory`
	FeatureSetCombination            FeatureSetCombinationId           `madatory`
	Ca_ParametersEUTRA               *CA_ParametersEUTRA               `optional`
	Ca_ParametersNR                  *CA_ParametersNR                  `optional`
	Mrdc_Parameters                  *MRDC_Parameters                  `optional`
	SupportedBandwidthCombinationSet *uper.BitString                   `lb:1,ub:32,optional`
	PowerClass_v1530                 *BandCombination_powerClass_v1530 `optional`
}

func (ie *BandCombination) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersEUTRA != nil, ie.Ca_ParametersNR != nil, ie.Mrdc_Parameters != nil, ie.SupportedBandwidthCombinationSet != nil, ie.PowerClass_v1530 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_BandList := utils.NewSequence[*BandParameters]([]*BandParameters{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.BandList {
		tmp_BandList.Value = append(tmp_BandList.Value, &i)
	}
	if err = tmp_BandList.Encode(w); err != nil {
		return utils.WrapError("Encode BandList", err)
	}
	if err = ie.FeatureSetCombination.Encode(w); err != nil {
		return utils.WrapError("Encode FeatureSetCombination", err)
	}
	if ie.Ca_ParametersEUTRA != nil {
		if err = ie.Ca_ParametersEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersEUTRA", err)
		}
	}
	if ie.Ca_ParametersNR != nil {
		if err = ie.Ca_ParametersNR.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR", err)
		}
	}
	if ie.Mrdc_Parameters != nil {
		if err = ie.Mrdc_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_Parameters", err)
		}
	}
	if ie.SupportedBandwidthCombinationSet != nil {
		if err = w.WriteBitString(ie.SupportedBandwidthCombinationSet.Bytes, uint(ie.SupportedBandwidthCombinationSet.NumBits), &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode SupportedBandwidthCombinationSet", err)
		}
	}
	if ie.PowerClass_v1530 != nil {
		if err = ie.PowerClass_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode PowerClass_v1530", err)
		}
	}
	return nil
}

func (ie *BandCombination) Decode(r *uper.UperReader) error {
	var err error
	var Ca_ParametersEUTRAPresent bool
	if Ca_ParametersEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRPresent bool
	if Ca_ParametersNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_ParametersPresent bool
	if Mrdc_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandwidthCombinationSetPresent bool
	if SupportedBandwidthCombinationSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PowerClass_v1530Present bool
	if PowerClass_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_BandList := utils.NewSequence[*BandParameters]([]*BandParameters{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_BandList := func() *BandParameters {
		return new(BandParameters)
	}
	if err = tmp_BandList.Decode(r, fn_BandList); err != nil {
		return utils.WrapError("Decode BandList", err)
	}
	ie.BandList = []BandParameters{}
	for _, i := range tmp_BandList.Value {
		ie.BandList = append(ie.BandList, *i)
	}
	if err = ie.FeatureSetCombination.Decode(r); err != nil {
		return utils.WrapError("Decode FeatureSetCombination", err)
	}
	if Ca_ParametersEUTRAPresent {
		ie.Ca_ParametersEUTRA = new(CA_ParametersEUTRA)
		if err = ie.Ca_ParametersEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersEUTRA", err)
		}
	}
	if Ca_ParametersNRPresent {
		ie.Ca_ParametersNR = new(CA_ParametersNR)
		if err = ie.Ca_ParametersNR.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR", err)
		}
	}
	if Mrdc_ParametersPresent {
		ie.Mrdc_Parameters = new(MRDC_Parameters)
		if err = ie.Mrdc_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_Parameters", err)
		}
	}
	if SupportedBandwidthCombinationSetPresent {
		var tmp_bs_SupportedBandwidthCombinationSet []byte
		var n_SupportedBandwidthCombinationSet uint
		if tmp_bs_SupportedBandwidthCombinationSet, n_SupportedBandwidthCombinationSet, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode SupportedBandwidthCombinationSet", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_SupportedBandwidthCombinationSet,
			NumBits: uint64(n_SupportedBandwidthCombinationSet),
		}
		ie.SupportedBandwidthCombinationSet = &tmp_bitstring
	}
	if PowerClass_v1530Present {
		ie.PowerClass_v1530 = new(BandCombination_powerClass_v1530)
		if err = ie.PowerClass_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode PowerClass_v1530", err)
		}
	}
	return nil
}
