package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type WLAN_RTT_r16 struct {
	RttValue_r16    int64                     `lb:0,ub:16777215,madatory`
	RttUnits_r16    WLAN_RTT_r16_rttUnits_r16 `madatory`
	RttAccuracy_r16 *int64                    `lb:0,ub:255,optional`
}

func (ie *WLAN_RTT_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RttAccuracy_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.RttValue_r16, &aper.Constraint{Lb: 0, Ub: 16777215}, false); err != nil {
		return utils.WrapError("WriteInteger RttValue_r16", err)
	}
	if err = ie.RttUnits_r16.Encode(w); err != nil {
		return utils.WrapError("Encode RttUnits_r16", err)
	}
	if ie.RttAccuracy_r16 != nil {
		if err = w.WriteInteger(*ie.RttAccuracy_r16, &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Encode RttAccuracy_r16", err)
		}
	}
	return nil
}

func (ie *WLAN_RTT_r16) Decode(r *aper.AperReader) error {
	var err error
	var RttAccuracy_r16Present bool
	if RttAccuracy_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_RttValue_r16 int64
	if tmp_int_RttValue_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16777215}, false); err != nil {
		return utils.WrapError("ReadInteger RttValue_r16", err)
	}
	ie.RttValue_r16 = tmp_int_RttValue_r16
	if err = ie.RttUnits_r16.Decode(r); err != nil {
		return utils.WrapError("Decode RttUnits_r16", err)
	}
	if RttAccuracy_r16Present {
		var tmp_int_RttAccuracy_r16 int64
		if tmp_int_RttAccuracy_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode RttAccuracy_r16", err)
		}
		ie.RttAccuracy_r16 = &tmp_int_RttAccuracy_r16
	}
	return nil
}
