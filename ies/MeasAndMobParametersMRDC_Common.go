package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common struct {
	IndependentGapConfig *MeasAndMobParametersMRDC_Common_independentGapConfig `optional`
}

func (ie *MeasAndMobParametersMRDC_Common) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IndependentGapConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IndependentGapConfig != nil {
		if err = ie.IndependentGapConfig.Encode(w); err != nil {
			return utils.WrapError("Encode IndependentGapConfig", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common) Decode(r *aper.AperReader) error {
	var err error
	var IndependentGapConfigPresent bool
	if IndependentGapConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if IndependentGapConfigPresent {
		ie.IndependentGapConfig = new(MeasAndMobParametersMRDC_Common_independentGapConfig)
		if err = ie.IndependentGapConfig.Decode(r); err != nil {
			return utils.WrapError("Decode IndependentGapConfig", err)
		}
	}
	return nil
}
