package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportQuantityUTRA_FDD_r16 struct {
	Cpich_RSCP bool `madatory`
	Cpich_EcN0 bool `madatory`
}

func (ie *MeasReportQuantityUTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Cpich_RSCP); err != nil {
		return utils.WrapError("WriteBoolean Cpich_RSCP", err)
	}
	if err = w.WriteBoolean(ie.Cpich_EcN0); err != nil {
		return utils.WrapError("WriteBoolean Cpich_EcN0", err)
	}
	return nil
}

func (ie *MeasReportQuantityUTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_Cpich_RSCP bool
	if tmp_bool_Cpich_RSCP, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Cpich_RSCP", err)
	}
	ie.Cpich_RSCP = tmp_bool_Cpich_RSCP
	var tmp_bool_Cpich_EcN0 bool
	if tmp_bool_Cpich_EcN0, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Cpich_EcN0", err)
	}
	ie.Cpich_EcN0 = tmp_bool_Cpich_EcN0
	return nil
}
