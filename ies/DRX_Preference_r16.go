package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Preference_r16 struct {
	PreferredDRX_InactivityTimer_r16 *DRX_Preference_r16_preferredDRX_InactivityTimer_r16 `optional`
	PreferredDRX_LongCycle_r16       *DRX_Preference_r16_preferredDRX_LongCycle_r16       `optional`
	PreferredDRX_ShortCycle_r16      *DRX_Preference_r16_preferredDRX_ShortCycle_r16      `optional`
	PreferredDRX_ShortCycleTimer_r16 *int64                                               `lb:1,ub:16,optional`
}

func (ie *DRX_Preference_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PreferredDRX_InactivityTimer_r16 != nil, ie.PreferredDRX_LongCycle_r16 != nil, ie.PreferredDRX_ShortCycle_r16 != nil, ie.PreferredDRX_ShortCycleTimer_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PreferredDRX_InactivityTimer_r16 != nil {
		if err = ie.PreferredDRX_InactivityTimer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredDRX_InactivityTimer_r16", err)
		}
	}
	if ie.PreferredDRX_LongCycle_r16 != nil {
		if err = ie.PreferredDRX_LongCycle_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredDRX_LongCycle_r16", err)
		}
	}
	if ie.PreferredDRX_ShortCycle_r16 != nil {
		if err = ie.PreferredDRX_ShortCycle_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredDRX_ShortCycle_r16", err)
		}
	}
	if ie.PreferredDRX_ShortCycleTimer_r16 != nil {
		if err = w.WriteInteger(*ie.PreferredDRX_ShortCycleTimer_r16, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode PreferredDRX_ShortCycleTimer_r16", err)
		}
	}
	return nil
}

func (ie *DRX_Preference_r16) Decode(r *aper.AperReader) error {
	var err error
	var PreferredDRX_InactivityTimer_r16Present bool
	if PreferredDRX_InactivityTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredDRX_LongCycle_r16Present bool
	if PreferredDRX_LongCycle_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredDRX_ShortCycle_r16Present bool
	if PreferredDRX_ShortCycle_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredDRX_ShortCycleTimer_r16Present bool
	if PreferredDRX_ShortCycleTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PreferredDRX_InactivityTimer_r16Present {
		ie.PreferredDRX_InactivityTimer_r16 = new(DRX_Preference_r16_preferredDRX_InactivityTimer_r16)
		if err = ie.PreferredDRX_InactivityTimer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredDRX_InactivityTimer_r16", err)
		}
	}
	if PreferredDRX_LongCycle_r16Present {
		ie.PreferredDRX_LongCycle_r16 = new(DRX_Preference_r16_preferredDRX_LongCycle_r16)
		if err = ie.PreferredDRX_LongCycle_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredDRX_LongCycle_r16", err)
		}
	}
	if PreferredDRX_ShortCycle_r16Present {
		ie.PreferredDRX_ShortCycle_r16 = new(DRX_Preference_r16_preferredDRX_ShortCycle_r16)
		if err = ie.PreferredDRX_ShortCycle_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredDRX_ShortCycle_r16", err)
		}
	}
	if PreferredDRX_ShortCycleTimer_r16Present {
		var tmp_int_PreferredDRX_ShortCycleTimer_r16 int64
		if tmp_int_PreferredDRX_ShortCycleTimer_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode PreferredDRX_ShortCycleTimer_r16", err)
		}
		ie.PreferredDRX_ShortCycleTimer_r16 = &tmp_int_PreferredDRX_ShortCycleTimer_r16
	}
	return nil
}
