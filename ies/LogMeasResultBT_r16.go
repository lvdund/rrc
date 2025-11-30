package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasResultBT_r16 struct {
	Bt_Addr_r16 aper.BitString `lb:48,ub:48,madatory`
	Rssi_BT_r16 *int64         `lb:-128,ub:127,optional`
}

func (ie *LogMeasResultBT_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rssi_BT_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.Bt_Addr_r16.Bytes, uint(ie.Bt_Addr_r16.NumBits), &aper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("WriteBitString Bt_Addr_r16", err)
	}
	if ie.Rssi_BT_r16 != nil {
		if err = w.WriteInteger(*ie.Rssi_BT_r16, &aper.Constraint{Lb: -128, Ub: 127}, false); err != nil {
			return utils.WrapError("Encode Rssi_BT_r16", err)
		}
	}
	return nil
}

func (ie *LogMeasResultBT_r16) Decode(r *aper.AperReader) error {
	var err error
	var Rssi_BT_r16Present bool
	if Rssi_BT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_Bt_Addr_r16 []byte
	var n_Bt_Addr_r16 uint
	if tmp_bs_Bt_Addr_r16, n_Bt_Addr_r16, err = r.ReadBitString(&aper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("ReadBitString Bt_Addr_r16", err)
	}
	ie.Bt_Addr_r16 = aper.BitString{
		Bytes:   tmp_bs_Bt_Addr_r16,
		NumBits: uint64(n_Bt_Addr_r16),
	}
	if Rssi_BT_r16Present {
		var tmp_int_Rssi_BT_r16 int64
		if tmp_int_Rssi_BT_r16, err = r.ReadInteger(&aper.Constraint{Lb: -128, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode Rssi_BT_r16", err)
		}
		ie.Rssi_BT_r16 = &tmp_int_Rssi_BT_r16
	}
	return nil
}
