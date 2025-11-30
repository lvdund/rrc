package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_PRS_MeasurementInfo_r16 struct {
	Dl_PRS_PointA_r16                  ARFCN_ValueNR                                                 `madatory`
	Nr_MeasPRS_RepetitionAndOffset_r16 NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16 `lb:0,ub:19,madatory`
	Nr_MeasPRS_length_r16              NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16              `madatory,ext`
}

func (ie *NR_PRS_MeasurementInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Dl_PRS_PointA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_PRS_PointA_r16", err)
	}
	if err = ie.Nr_MeasPRS_RepetitionAndOffset_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Nr_MeasPRS_RepetitionAndOffset_r16", err)
	}
	return nil
}

func (ie *NR_PRS_MeasurementInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Dl_PRS_PointA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_PRS_PointA_r16", err)
	}
	if err = ie.Nr_MeasPRS_RepetitionAndOffset_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Nr_MeasPRS_RepetitionAndOffset_r16", err)
	}
	return nil
}
