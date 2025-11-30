package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentTwoCarrierInfo_r16 struct {
	ReferenceCarrierIndex_r16   ServCellIndex `madatory`
	Shift7dot5kHz_r16           bool          `madatory`
	TxDirectCurrentLocation_r16 int64         `lb:0,ub:3301,madatory`
}

func (ie *UplinkTxDirectCurrentTwoCarrierInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReferenceCarrierIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceCarrierIndex_r16", err)
	}
	if err = w.WriteBoolean(ie.Shift7dot5kHz_r16); err != nil {
		return utils.WrapError("WriteBoolean Shift7dot5kHz_r16", err)
	}
	if err = w.WriteInteger(ie.TxDirectCurrentLocation_r16, &aper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("WriteInteger TxDirectCurrentLocation_r16", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentTwoCarrierInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReferenceCarrierIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceCarrierIndex_r16", err)
	}
	var tmp_bool_Shift7dot5kHz_r16 bool
	if tmp_bool_Shift7dot5kHz_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Shift7dot5kHz_r16", err)
	}
	ie.Shift7dot5kHz_r16 = tmp_bool_Shift7dot5kHz_r16
	var tmp_int_TxDirectCurrentLocation_r16 int64
	if tmp_int_TxDirectCurrentLocation_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("ReadInteger TxDirectCurrentLocation_r16", err)
	}
	ie.TxDirectCurrentLocation_r16 = tmp_int_TxDirectCurrentLocation_r16
	return nil
}
