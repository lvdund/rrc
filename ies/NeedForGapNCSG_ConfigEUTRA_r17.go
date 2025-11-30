package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapNCSG_ConfigEUTRA_r17 struct {
	RequestedTargetBandFilterNCSG_EUTRA_r17 []FreqBandIndicatorEUTRA `lb:1,ub:maxBandsEUTRA,optional`
}

func (ie *NeedForGapNCSG_ConfigEUTRA_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.RequestedTargetBandFilterNCSG_EUTRA_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.RequestedTargetBandFilterNCSG_EUTRA_r17) > 0 {
		tmp_RequestedTargetBandFilterNCSG_EUTRA_r17 := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, aper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		for _, i := range ie.RequestedTargetBandFilterNCSG_EUTRA_r17 {
			tmp_RequestedTargetBandFilterNCSG_EUTRA_r17.Value = append(tmp_RequestedTargetBandFilterNCSG_EUTRA_r17.Value, &i)
		}
		if err = tmp_RequestedTargetBandFilterNCSG_EUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RequestedTargetBandFilterNCSG_EUTRA_r17", err)
		}
	}
	return nil
}

func (ie *NeedForGapNCSG_ConfigEUTRA_r17) Decode(r *aper.AperReader) error {
	var err error
	var RequestedTargetBandFilterNCSG_EUTRA_r17Present bool
	if RequestedTargetBandFilterNCSG_EUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RequestedTargetBandFilterNCSG_EUTRA_r17Present {
		tmp_RequestedTargetBandFilterNCSG_EUTRA_r17 := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, aper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		fn_RequestedTargetBandFilterNCSG_EUTRA_r17 := func() *FreqBandIndicatorEUTRA {
			return new(FreqBandIndicatorEUTRA)
		}
		if err = tmp_RequestedTargetBandFilterNCSG_EUTRA_r17.Decode(r, fn_RequestedTargetBandFilterNCSG_EUTRA_r17); err != nil {
			return utils.WrapError("Decode RequestedTargetBandFilterNCSG_EUTRA_r17", err)
		}
		ie.RequestedTargetBandFilterNCSG_EUTRA_r17 = []FreqBandIndicatorEUTRA{}
		for _, i := range tmp_RequestedTargetBandFilterNCSG_EUTRA_r17.Value {
			ie.RequestedTargetBandFilterNCSG_EUTRA_r17 = append(ie.RequestedTargetBandFilterNCSG_EUTRA_r17, *i)
		}
	}
	return nil
}
