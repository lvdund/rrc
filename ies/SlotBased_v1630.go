package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SlotBased_v1630 struct {
	TciMapping_r16          SlotBased_v1630_tciMapping_r16 `madatory`
	SequenceOffsetForRV_r16 int64                          `lb:0,ub:0,madatory`
}

func (ie *SlotBased_v1630) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.TciMapping_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TciMapping_r16", err)
	}
	if err = w.WriteInteger(ie.SequenceOffsetForRV_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteInteger SequenceOffsetForRV_r16", err)
	}
	return nil
}

func (ie *SlotBased_v1630) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.TciMapping_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TciMapping_r16", err)
	}
	var tmp_int_SequenceOffsetForRV_r16 int64
	if tmp_int_SequenceOffsetForRV_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadInteger SequenceOffsetForRV_r16", err)
	}
	ie.SequenceOffsetForRV_r16 = tmp_int_SequenceOffsetForRV_r16
	return nil
}
