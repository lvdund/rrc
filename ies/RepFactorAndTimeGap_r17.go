package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RepFactorAndTimeGap_r17 struct {
	RepetitionFactor_r17 RepFactorAndTimeGap_r17_repetitionFactor_r17 `madatory`
	TimeGap_r17          RepFactorAndTimeGap_r17_timeGap_r17          `madatory`
}

func (ie *RepFactorAndTimeGap_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.RepetitionFactor_r17.Encode(w); err != nil {
		return utils.WrapError("Encode RepetitionFactor_r17", err)
	}
	if err = ie.TimeGap_r17.Encode(w); err != nil {
		return utils.WrapError("Encode TimeGap_r17", err)
	}
	return nil
}

func (ie *RepFactorAndTimeGap_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.RepetitionFactor_r17.Decode(r); err != nil {
		return utils.WrapError("Decode RepetitionFactor_r17", err)
	}
	if err = ie.TimeGap_r17.Decode(r); err != nil {
		return utils.WrapError("Decode TimeGap_r17", err)
	}
	return nil
}
