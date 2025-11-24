package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1630 struct {
	Ca_ParametersNR_v1630                     *CA_ParametersNR_v1630      `optional`
	Ca_ParametersNRDC_v1630                   *CA_ParametersNRDC_v1630    `optional`
	Mrdc_Parameters_v1630                     *MRDC_Parameters_v1630      `optional`
	SupportedTxBandCombListPerBC_Sidelink_r16 *uper.BitString             `lb:1,ub:maxBandComb,optional`
	SupportedRxBandCombListPerBC_Sidelink_r16 *uper.BitString             `lb:1,ub:maxBandComb,optional`
	ScalingFactorTxSidelink_r16               []ScalingFactorSidelink_r16 `lb:1,ub:maxBandComb,optional`
	ScalingFactorRxSidelink_r16               []ScalingFactorSidelink_r16 `lb:1,ub:maxBandComb,optional`
}

func (ie *BandCombination_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v1630 != nil, ie.Ca_ParametersNRDC_v1630 != nil, ie.Mrdc_Parameters_v1630 != nil, ie.SupportedTxBandCombListPerBC_Sidelink_r16 != nil, ie.SupportedRxBandCombListPerBC_Sidelink_r16 != nil, len(ie.ScalingFactorTxSidelink_r16) > 0, len(ie.ScalingFactorRxSidelink_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_v1630 != nil {
		if err = ie.Ca_ParametersNR_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1630", err)
		}
	}
	if ie.Ca_ParametersNRDC_v1630 != nil {
		if err = ie.Ca_ParametersNRDC_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v1630", err)
		}
	}
	if ie.Mrdc_Parameters_v1630 != nil {
		if err = ie.Mrdc_Parameters_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_Parameters_v1630", err)
		}
	}
	if ie.SupportedTxBandCombListPerBC_Sidelink_r16 != nil {
		if err = w.WriteBitString(ie.SupportedTxBandCombListPerBC_Sidelink_r16.Bytes, uint(ie.SupportedTxBandCombListPerBC_Sidelink_r16.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode SupportedTxBandCombListPerBC_Sidelink_r16", err)
		}
	}
	if ie.SupportedRxBandCombListPerBC_Sidelink_r16 != nil {
		if err = w.WriteBitString(ie.SupportedRxBandCombListPerBC_Sidelink_r16.Bytes, uint(ie.SupportedRxBandCombListPerBC_Sidelink_r16.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode SupportedRxBandCombListPerBC_Sidelink_r16", err)
		}
	}
	if len(ie.ScalingFactorTxSidelink_r16) > 0 {
		tmp_ScalingFactorTxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		for _, i := range ie.ScalingFactorTxSidelink_r16 {
			tmp_ScalingFactorTxSidelink_r16.Value = append(tmp_ScalingFactorTxSidelink_r16.Value, &i)
		}
		if err = tmp_ScalingFactorTxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ScalingFactorTxSidelink_r16", err)
		}
	}
	if len(ie.ScalingFactorRxSidelink_r16) > 0 {
		tmp_ScalingFactorRxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		for _, i := range ie.ScalingFactorRxSidelink_r16 {
			tmp_ScalingFactorRxSidelink_r16.Value = append(tmp_ScalingFactorRxSidelink_r16.Value, &i)
		}
		if err = tmp_ScalingFactorRxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ScalingFactorRxSidelink_r16", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1630) Decode(r *uper.UperReader) error {
	var err error
	var Ca_ParametersNR_v1630Present bool
	if Ca_ParametersNR_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v1630Present bool
	if Ca_ParametersNRDC_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_Parameters_v1630Present bool
	if Mrdc_Parameters_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedTxBandCombListPerBC_Sidelink_r16Present bool
	if SupportedTxBandCombListPerBC_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedRxBandCombListPerBC_Sidelink_r16Present bool
	if SupportedRxBandCombListPerBC_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ScalingFactorTxSidelink_r16Present bool
	if ScalingFactorTxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ScalingFactorRxSidelink_r16Present bool
	if ScalingFactorRxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_v1630Present {
		ie.Ca_ParametersNR_v1630 = new(CA_ParametersNR_v1630)
		if err = ie.Ca_ParametersNR_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1630", err)
		}
	}
	if Ca_ParametersNRDC_v1630Present {
		ie.Ca_ParametersNRDC_v1630 = new(CA_ParametersNRDC_v1630)
		if err = ie.Ca_ParametersNRDC_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v1630", err)
		}
	}
	if Mrdc_Parameters_v1630Present {
		ie.Mrdc_Parameters_v1630 = new(MRDC_Parameters_v1630)
		if err = ie.Mrdc_Parameters_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_Parameters_v1630", err)
		}
	}
	if SupportedTxBandCombListPerBC_Sidelink_r16Present {
		var tmp_bs_SupportedTxBandCombListPerBC_Sidelink_r16 []byte
		var n_SupportedTxBandCombListPerBC_Sidelink_r16 uint
		if tmp_bs_SupportedTxBandCombListPerBC_Sidelink_r16, n_SupportedTxBandCombListPerBC_Sidelink_r16, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode SupportedTxBandCombListPerBC_Sidelink_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_SupportedTxBandCombListPerBC_Sidelink_r16,
			NumBits: uint64(n_SupportedTxBandCombListPerBC_Sidelink_r16),
		}
		ie.SupportedTxBandCombListPerBC_Sidelink_r16 = &tmp_bitstring
	}
	if SupportedRxBandCombListPerBC_Sidelink_r16Present {
		var tmp_bs_SupportedRxBandCombListPerBC_Sidelink_r16 []byte
		var n_SupportedRxBandCombListPerBC_Sidelink_r16 uint
		if tmp_bs_SupportedRxBandCombListPerBC_Sidelink_r16, n_SupportedRxBandCombListPerBC_Sidelink_r16, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode SupportedRxBandCombListPerBC_Sidelink_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_SupportedRxBandCombListPerBC_Sidelink_r16,
			NumBits: uint64(n_SupportedRxBandCombListPerBC_Sidelink_r16),
		}
		ie.SupportedRxBandCombListPerBC_Sidelink_r16 = &tmp_bitstring
	}
	if ScalingFactorTxSidelink_r16Present {
		tmp_ScalingFactorTxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		fn_ScalingFactorTxSidelink_r16 := func() *ScalingFactorSidelink_r16 {
			return new(ScalingFactorSidelink_r16)
		}
		if err = tmp_ScalingFactorTxSidelink_r16.Decode(r, fn_ScalingFactorTxSidelink_r16); err != nil {
			return utils.WrapError("Decode ScalingFactorTxSidelink_r16", err)
		}
		ie.ScalingFactorTxSidelink_r16 = []ScalingFactorSidelink_r16{}
		for _, i := range tmp_ScalingFactorTxSidelink_r16.Value {
			ie.ScalingFactorTxSidelink_r16 = append(ie.ScalingFactorTxSidelink_r16, *i)
		}
	}
	if ScalingFactorRxSidelink_r16Present {
		tmp_ScalingFactorRxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		fn_ScalingFactorRxSidelink_r16 := func() *ScalingFactorSidelink_r16 {
			return new(ScalingFactorSidelink_r16)
		}
		if err = tmp_ScalingFactorRxSidelink_r16.Decode(r, fn_ScalingFactorRxSidelink_r16); err != nil {
			return utils.WrapError("Decode ScalingFactorRxSidelink_r16", err)
		}
		ie.ScalingFactorRxSidelink_r16 = []ScalingFactorSidelink_r16{}
		for _, i := range tmp_ScalingFactorRxSidelink_r16.Value {
			ie.ScalingFactorRxSidelink_r16 = append(ie.ScalingFactorRxSidelink_r16, *i)
		}
	}
	return nil
}
