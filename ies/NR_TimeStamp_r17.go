package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_TimeStamp_r17 struct {
	Nr_SFN_r17  int64                        `lb:0,ub:1023,madatory`
	Nr_Slot_r17 NR_TimeStamp_r17_nr_Slot_r17 `lb:0,ub:9,madatory`
}

func (ie *NR_TimeStamp_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Nr_SFN_r17, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger Nr_SFN_r17", err)
	}
	if err = ie.Nr_Slot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Nr_Slot_r17", err)
	}
	return nil
}

func (ie *NR_TimeStamp_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Nr_SFN_r17 int64
	if tmp_int_Nr_SFN_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger Nr_SFN_r17", err)
	}
	ie.Nr_SFN_r17 = tmp_int_Nr_SFN_r17
	if err = ie.Nr_Slot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Nr_Slot_r17", err)
	}
	return nil
}
