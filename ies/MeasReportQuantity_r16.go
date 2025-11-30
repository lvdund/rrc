package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportQuantity_r16 struct {
	Cbr_r16 bool `madatory`
}

func (ie *MeasReportQuantity_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Cbr_r16); err != nil {
		return utils.WrapError("WriteBoolean Cbr_r16", err)
	}
	return nil
}

func (ie *MeasReportQuantity_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Cbr_r16 bool
	if tmp_bool_Cbr_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Cbr_r16", err)
	}
	ie.Cbr_r16 = tmp_bool_Cbr_r16
	return nil
}
