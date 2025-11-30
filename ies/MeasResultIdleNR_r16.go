package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultIdleNR_r16 struct {
	MeasResultServingCell_r16           *MeasResultIdleNR_r16_measResultServingCell_r16 `optional`
	MeasResultsPerCarrierListIdleNR_r16 []MeasResultsPerCarrierIdleNR_r16               `lb:1,ub:maxFreqIdle_r16,optional`
}

func (ie *MeasResultIdleNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultServingCell_r16 != nil, len(ie.MeasResultsPerCarrierListIdleNR_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasResultServingCell_r16 != nil {
		if err = ie.MeasResultServingCell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultServingCell_r16", err)
		}
	}
	if len(ie.MeasResultsPerCarrierListIdleNR_r16) > 0 {
		tmp_MeasResultsPerCarrierListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleNR_r16]([]*MeasResultsPerCarrierIdleNR_r16{}, aper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		for _, i := range ie.MeasResultsPerCarrierListIdleNR_r16 {
			tmp_MeasResultsPerCarrierListIdleNR_r16.Value = append(tmp_MeasResultsPerCarrierListIdleNR_r16.Value, &i)
		}
		if err = tmp_MeasResultsPerCarrierListIdleNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultsPerCarrierListIdleNR_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultIdleNR_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasResultServingCell_r16Present bool
	if MeasResultServingCell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultsPerCarrierListIdleNR_r16Present bool
	if MeasResultsPerCarrierListIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasResultServingCell_r16Present {
		ie.MeasResultServingCell_r16 = new(MeasResultIdleNR_r16_measResultServingCell_r16)
		if err = ie.MeasResultServingCell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultServingCell_r16", err)
		}
	}
	if MeasResultsPerCarrierListIdleNR_r16Present {
		tmp_MeasResultsPerCarrierListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleNR_r16]([]*MeasResultsPerCarrierIdleNR_r16{}, aper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		fn_MeasResultsPerCarrierListIdleNR_r16 := func() *MeasResultsPerCarrierIdleNR_r16 {
			return new(MeasResultsPerCarrierIdleNR_r16)
		}
		if err = tmp_MeasResultsPerCarrierListIdleNR_r16.Decode(r, fn_MeasResultsPerCarrierListIdleNR_r16); err != nil {
			return utils.WrapError("Decode MeasResultsPerCarrierListIdleNR_r16", err)
		}
		ie.MeasResultsPerCarrierListIdleNR_r16 = []MeasResultsPerCarrierIdleNR_r16{}
		for _, i := range tmp_MeasResultsPerCarrierListIdleNR_r16.Value {
			ie.MeasResultsPerCarrierListIdleNR_r16 = append(ie.MeasResultsPerCarrierListIdleNR_r16, *i)
		}
	}
	return nil
}
