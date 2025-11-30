package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSCCH_Config_r16 struct {
	Sl_TimeResourcePSCCH_r16 *SL_PSCCH_Config_r16_sl_TimeResourcePSCCH_r16 `optional`
	Sl_FreqResourcePSCCH_r16 *SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16 `optional`
	Sl_DMRS_ScrambleID_r16   *int64                                        `lb:0,ub:65535,optional`
	Sl_NumReservedBits_r16   *int64                                        `lb:2,ub:4,optional`
}

func (ie *SL_PSCCH_Config_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_TimeResourcePSCCH_r16 != nil, ie.Sl_FreqResourcePSCCH_r16 != nil, ie.Sl_DMRS_ScrambleID_r16 != nil, ie.Sl_NumReservedBits_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_TimeResourcePSCCH_r16 != nil {
		if err = ie.Sl_TimeResourcePSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TimeResourcePSCCH_r16", err)
		}
	}
	if ie.Sl_FreqResourcePSCCH_r16 != nil {
		if err = ie.Sl_FreqResourcePSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_FreqResourcePSCCH_r16", err)
		}
	}
	if ie.Sl_DMRS_ScrambleID_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_DMRS_ScrambleID_r16, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode Sl_DMRS_ScrambleID_r16", err)
		}
	}
	if ie.Sl_NumReservedBits_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_NumReservedBits_r16, &aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode Sl_NumReservedBits_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSCCH_Config_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_TimeResourcePSCCH_r16Present bool
	if Sl_TimeResourcePSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_FreqResourcePSCCH_r16Present bool
	if Sl_FreqResourcePSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_DMRS_ScrambleID_r16Present bool
	if Sl_DMRS_ScrambleID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_NumReservedBits_r16Present bool
	if Sl_NumReservedBits_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_TimeResourcePSCCH_r16Present {
		ie.Sl_TimeResourcePSCCH_r16 = new(SL_PSCCH_Config_r16_sl_TimeResourcePSCCH_r16)
		if err = ie.Sl_TimeResourcePSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TimeResourcePSCCH_r16", err)
		}
	}
	if Sl_FreqResourcePSCCH_r16Present {
		ie.Sl_FreqResourcePSCCH_r16 = new(SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16)
		if err = ie.Sl_FreqResourcePSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_FreqResourcePSCCH_r16", err)
		}
	}
	if Sl_DMRS_ScrambleID_r16Present {
		var tmp_int_Sl_DMRS_ScrambleID_r16 int64
		if tmp_int_Sl_DMRS_ScrambleID_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode Sl_DMRS_ScrambleID_r16", err)
		}
		ie.Sl_DMRS_ScrambleID_r16 = &tmp_int_Sl_DMRS_ScrambleID_r16
	}
	if Sl_NumReservedBits_r16Present {
		var tmp_int_Sl_NumReservedBits_r16 int64
		if tmp_int_Sl_NumReservedBits_r16, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sl_NumReservedBits_r16", err)
		}
		ie.Sl_NumReservedBits_r16 = &tmp_int_Sl_NumReservedBits_r16
	}
	return nil
}
