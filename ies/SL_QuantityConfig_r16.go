package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_QuantityConfig_r16 struct {
	Sl_FilterCoefficientDMRS_r16 FilterCoefficient `madatory`
}

func (ie *SL_QuantityConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Sl_FilterCoefficientDMRS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_FilterCoefficientDMRS_r16", err)
	}
	return nil
}

func (ie *SL_QuantityConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Sl_FilterCoefficientDMRS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_FilterCoefficientDMRS_r16", err)
	}
	return nil
}
