package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16 struct {
	SlotOffset_r16 *int64 `lb:1,ub:32,optional`
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SlotOffset_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SlotOffset_r16 != nil {
		if err = w.WriteInteger(*ie.SlotOffset_r16, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode SlotOffset_r16", err)
		}
	}
	return nil
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16) Decode(r *aper.AperReader) error {
	var err error
	var SlotOffset_r16Present bool
	if SlotOffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SlotOffset_r16Present {
		var tmp_int_SlotOffset_r16 int64
		if tmp_int_SlotOffset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode SlotOffset_r16", err)
		}
		ie.SlotOffset_r16 = &tmp_int_SlotOffset_r16
	}
	return nil
}
