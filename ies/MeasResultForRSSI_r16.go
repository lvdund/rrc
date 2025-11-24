package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultForRSSI_r16 struct {
	Rssi_Result_r16      RSSI_Range_r16 `madatory`
	ChannelOccupancy_r16 int64          `lb:0,ub:100,madatory`
}

func (ie *MeasResultForRSSI_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Rssi_Result_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rssi_Result_r16", err)
	}
	if err = w.WriteInteger(ie.ChannelOccupancy_r16, &uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
		return utils.WrapError("WriteInteger ChannelOccupancy_r16", err)
	}
	return nil
}

func (ie *MeasResultForRSSI_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Rssi_Result_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rssi_Result_r16", err)
	}
	var tmp_int_ChannelOccupancy_r16 int64
	if tmp_int_ChannelOccupancy_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
		return utils.WrapError("ReadInteger ChannelOccupancy_r16", err)
	}
	ie.ChannelOccupancy_r16 = tmp_int_ChannelOccupancy_r16
	return nil
}
