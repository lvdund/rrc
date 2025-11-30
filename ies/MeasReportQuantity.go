package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportQuantity struct {
	Rsrp bool `madatory`
	Rsrq bool `madatory`
	Sinr bool `madatory`
}

func (ie *MeasReportQuantity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Rsrp); err != nil {
		return utils.WrapError("WriteBoolean Rsrp", err)
	}
	if err = w.WriteBoolean(ie.Rsrq); err != nil {
		return utils.WrapError("WriteBoolean Rsrq", err)
	}
	if err = w.WriteBoolean(ie.Sinr); err != nil {
		return utils.WrapError("WriteBoolean Sinr", err)
	}
	return nil
}

func (ie *MeasReportQuantity) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Rsrp bool
	if tmp_bool_Rsrp, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Rsrp", err)
	}
	ie.Rsrp = tmp_bool_Rsrp
	var tmp_bool_Rsrq bool
	if tmp_bool_Rsrq, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Rsrq", err)
	}
	ie.Rsrq = tmp_bool_Rsrq
	var tmp_bool_Sinr bool
	if tmp_bool_Sinr, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Sinr", err)
	}
	ie.Sinr = tmp_bool_Sinr
	return nil
}
