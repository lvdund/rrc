package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_UplinkConfig_transformPrecoderEnabled struct {
	SampleDensity                 []int64                                                                   `lb:5,ub:5,e_lb:0,e_ub:0,madatory`
	TimeDensityTransformPrecoding *PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding `optional`
}

func (ie *PTRS_UplinkConfig_transformPrecoderEnabled) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.TimeDensityTransformPrecoding != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_SampleDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 5, Ub: 5}, false)
	for _, i := range ie.SampleDensity {
		tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 0}, false)
		tmp_SampleDensity.Value = append(tmp_SampleDensity.Value, &tmp_ie)
	}
	if err = tmp_SampleDensity.Encode(w); err != nil {
		return utils.WrapError("Encode SampleDensity", err)
	}
	if ie.TimeDensityTransformPrecoding != nil {
		if err = ie.TimeDensityTransformPrecoding.Encode(w); err != nil {
			return utils.WrapError("Encode TimeDensityTransformPrecoding", err)
		}
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderEnabled) Decode(r *aper.AperReader) error {
	var err error
	var TimeDensityTransformPrecodingPresent bool
	if TimeDensityTransformPrecodingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_SampleDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 5, Ub: 5}, false)
	fn_SampleDensity := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 0}, false)
		return &ie
	}
	if err = tmp_SampleDensity.Decode(r, fn_SampleDensity); err != nil {
		return utils.WrapError("Decode SampleDensity", err)
	}
	ie.SampleDensity = []int64{}
	for _, i := range tmp_SampleDensity.Value {
		ie.SampleDensity = append(ie.SampleDensity, int64(i.Value))
	}
	if TimeDensityTransformPrecodingPresent {
		ie.TimeDensityTransformPrecoding = new(PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding)
		if err = ie.TimeDensityTransformPrecoding.Decode(r); err != nil {
			return utils.WrapError("Decode TimeDensityTransformPrecoding", err)
		}
	}
	return nil
}
