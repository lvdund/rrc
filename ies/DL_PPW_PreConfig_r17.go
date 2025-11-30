package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DL_PPW_PreConfig_r17 struct {
	Dl_PPW_ID_r17                      DL_PPW_ID_r17                      `madatory`
	Dl_PPW_PeriodicityAndStartSlot_r17 DL_PPW_PeriodicityAndStartSlot_r17 `madatory`
	Length_r17                         int64                              `lb:1,ub:160,madatory`
	Type_r17                           *DL_PPW_PreConfig_r17_type_r17     `optional`
	Priority_r17                       *DL_PPW_PreConfig_r17_priority_r17 `optional`
}

func (ie *DL_PPW_PreConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Type_r17 != nil, ie.Priority_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Dl_PPW_ID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_PPW_ID_r17", err)
	}
	if err = ie.Dl_PPW_PeriodicityAndStartSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_PPW_PeriodicityAndStartSlot_r17", err)
	}
	if err = w.WriteInteger(ie.Length_r17, &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return utils.WrapError("WriteInteger Length_r17", err)
	}
	if ie.Type_r17 != nil {
		if err = ie.Type_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Type_r17", err)
		}
	}
	if ie.Priority_r17 != nil {
		if err = ie.Priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Priority_r17", err)
		}
	}
	return nil
}

func (ie *DL_PPW_PreConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Type_r17Present bool
	if Type_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Priority_r17Present bool
	if Priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Dl_PPW_ID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_PPW_ID_r17", err)
	}
	if err = ie.Dl_PPW_PeriodicityAndStartSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_PPW_PeriodicityAndStartSlot_r17", err)
	}
	var tmp_int_Length_r17 int64
	if tmp_int_Length_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return utils.WrapError("ReadInteger Length_r17", err)
	}
	ie.Length_r17 = tmp_int_Length_r17
	if Type_r17Present {
		ie.Type_r17 = new(DL_PPW_PreConfig_r17_type_r17)
		if err = ie.Type_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Type_r17", err)
		}
	}
	if Priority_r17Present {
		ie.Priority_r17 = new(DL_PPW_PreConfig_r17_priority_r17)
		if err = ie.Priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Priority_r17", err)
		}
	}
	return nil
}
