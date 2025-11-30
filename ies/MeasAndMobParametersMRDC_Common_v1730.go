package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1730 struct {
	IndependentGapConfig_maxCC_r17 *MeasAndMobParametersMRDC_Common_v1730_independentGapConfig_maxCC_r17 `lb:1,ub:32,optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1730) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IndependentGapConfig_maxCC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IndependentGapConfig_maxCC_r17 != nil {
		if err = ie.IndependentGapConfig_maxCC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode IndependentGapConfig_maxCC_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1730) Decode(r *aper.AperReader) error {
	var err error
	var IndependentGapConfig_maxCC_r17Present bool
	if IndependentGapConfig_maxCC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if IndependentGapConfig_maxCC_r17Present {
		ie.IndependentGapConfig_maxCC_r17 = new(MeasAndMobParametersMRDC_Common_v1730_independentGapConfig_maxCC_r17)
		if err = ie.IndependentGapConfig_maxCC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode IndependentGapConfig_maxCC_r17", err)
		}
	}
	return nil
}
