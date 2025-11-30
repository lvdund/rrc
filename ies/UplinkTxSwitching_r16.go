package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxSwitching_r16 struct {
	UplinkTxSwitchingPeriodLocation_r16 bool                                               `madatory`
	UplinkTxSwitchingCarrier_r16        UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16 `madatory`
}

func (ie *UplinkTxSwitching_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.UplinkTxSwitchingPeriodLocation_r16); err != nil {
		return utils.WrapError("WriteBoolean UplinkTxSwitchingPeriodLocation_r16", err)
	}
	if err = ie.UplinkTxSwitchingCarrier_r16.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkTxSwitchingCarrier_r16", err)
	}
	return nil
}

func (ie *UplinkTxSwitching_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_bool_UplinkTxSwitchingPeriodLocation_r16 bool
	if tmp_bool_UplinkTxSwitchingPeriodLocation_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean UplinkTxSwitchingPeriodLocation_r16", err)
	}
	ie.UplinkTxSwitchingPeriodLocation_r16 = tmp_bool_UplinkTxSwitchingPeriodLocation_r16
	if err = ie.UplinkTxSwitchingCarrier_r16.Decode(r); err != nil {
		return utils.WrapError("Decode UplinkTxSwitchingCarrier_r16", err)
	}
	return nil
}
