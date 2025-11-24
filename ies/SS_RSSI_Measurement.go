package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SS_RSSI_Measurement struct {
	MeasurementSlots uper.BitString `lb:1,ub:80,madatory`
	EndSymbol        int64          `lb:0,ub:3,madatory`
}

func (ie *SS_RSSI_Measurement) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.MeasurementSlots.Bytes, uint(ie.MeasurementSlots.NumBits), &uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
		return utils.WrapError("WriteBitString MeasurementSlots", err)
	}
	if err = w.WriteInteger(ie.EndSymbol, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger EndSymbol", err)
	}
	return nil
}

func (ie *SS_RSSI_Measurement) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_MeasurementSlots []byte
	var n_MeasurementSlots uint
	if tmp_bs_MeasurementSlots, n_MeasurementSlots, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
		return utils.WrapError("ReadBitString MeasurementSlots", err)
	}
	ie.MeasurementSlots = uper.BitString{
		Bytes:   tmp_bs_MeasurementSlots,
		NumBits: uint64(n_MeasurementSlots),
	}
	var tmp_int_EndSymbol int64
	if tmp_int_EndSymbol, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger EndSymbol", err)
	}
	ie.EndSymbol = tmp_int_EndSymbol
	return nil
}
