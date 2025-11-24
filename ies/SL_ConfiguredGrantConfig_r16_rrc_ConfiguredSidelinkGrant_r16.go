package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16 struct {
	Sl_TimeResourceCG_Type1_r16    *int64                                                                                      `lb:0,ub:496,optional`
	Sl_StartSubchannelCG_Type1_r16 *int64                                                                                      `lb:0,ub:26,optional`
	Sl_FreqResourceCG_Type1_r16    *int64                                                                                      `lb:0,ub:6929,optional`
	Sl_TimeOffsetCG_Type1_r16      *int64                                                                                      `lb:0,ub:7999,optional`
	Sl_N1PUCCH_AN_r16              *PUCCH_ResourceId                                                                           `optional`
	Sl_PSFCH_ToPUCCH_CG_Type1_r16  *int64                                                                                      `lb:0,ub:15,optional`
	Sl_ResourcePoolID_r16          *SL_ResourcePoolID_r16                                                                      `optional`
	Sl_TimeReferenceSFN_Type1_r16  *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16_sl_TimeReferenceSFN_Type1_r16 `optional`
}

func (ie *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_TimeResourceCG_Type1_r16 != nil, ie.Sl_StartSubchannelCG_Type1_r16 != nil, ie.Sl_FreqResourceCG_Type1_r16 != nil, ie.Sl_TimeOffsetCG_Type1_r16 != nil, ie.Sl_N1PUCCH_AN_r16 != nil, ie.Sl_PSFCH_ToPUCCH_CG_Type1_r16 != nil, ie.Sl_ResourcePoolID_r16 != nil, ie.Sl_TimeReferenceSFN_Type1_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_TimeResourceCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_TimeResourceCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode Sl_TimeResourceCG_Type1_r16", err)
		}
	}
	if ie.Sl_StartSubchannelCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_StartSubchannelCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 26}, false); err != nil {
			return utils.WrapError("Encode Sl_StartSubchannelCG_Type1_r16", err)
		}
	}
	if ie.Sl_FreqResourceCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_FreqResourceCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 6929}, false); err != nil {
			return utils.WrapError("Encode Sl_FreqResourceCG_Type1_r16", err)
		}
	}
	if ie.Sl_TimeOffsetCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_TimeOffsetCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 7999}, false); err != nil {
			return utils.WrapError("Encode Sl_TimeOffsetCG_Type1_r16", err)
		}
	}
	if ie.Sl_N1PUCCH_AN_r16 != nil {
		if err = ie.Sl_N1PUCCH_AN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_N1PUCCH_AN_r16", err)
		}
	}
	if ie.Sl_PSFCH_ToPUCCH_CG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_PSFCH_ToPUCCH_CG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Sl_PSFCH_ToPUCCH_CG_Type1_r16", err)
		}
	}
	if ie.Sl_ResourcePoolID_r16 != nil {
		if err = ie.Sl_ResourcePoolID_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ResourcePoolID_r16", err)
		}
	}
	if ie.Sl_TimeReferenceSFN_Type1_r16 != nil {
		if err = ie.Sl_TimeReferenceSFN_Type1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TimeReferenceSFN_Type1_r16", err)
		}
	}
	return nil
}

func (ie *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_TimeResourceCG_Type1_r16Present bool
	if Sl_TimeResourceCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_StartSubchannelCG_Type1_r16Present bool
	if Sl_StartSubchannelCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_FreqResourceCG_Type1_r16Present bool
	if Sl_FreqResourceCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TimeOffsetCG_Type1_r16Present bool
	if Sl_TimeOffsetCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_N1PUCCH_AN_r16Present bool
	if Sl_N1PUCCH_AN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSFCH_ToPUCCH_CG_Type1_r16Present bool
	if Sl_PSFCH_ToPUCCH_CG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ResourcePoolID_r16Present bool
	if Sl_ResourcePoolID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TimeReferenceSFN_Type1_r16Present bool
	if Sl_TimeReferenceSFN_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_TimeResourceCG_Type1_r16Present {
		var tmp_int_Sl_TimeResourceCG_Type1_r16 int64
		if tmp_int_Sl_TimeResourceCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode Sl_TimeResourceCG_Type1_r16", err)
		}
		ie.Sl_TimeResourceCG_Type1_r16 = &tmp_int_Sl_TimeResourceCG_Type1_r16
	}
	if Sl_StartSubchannelCG_Type1_r16Present {
		var tmp_int_Sl_StartSubchannelCG_Type1_r16 int64
		if tmp_int_Sl_StartSubchannelCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 26}, false); err != nil {
			return utils.WrapError("Decode Sl_StartSubchannelCG_Type1_r16", err)
		}
		ie.Sl_StartSubchannelCG_Type1_r16 = &tmp_int_Sl_StartSubchannelCG_Type1_r16
	}
	if Sl_FreqResourceCG_Type1_r16Present {
		var tmp_int_Sl_FreqResourceCG_Type1_r16 int64
		if tmp_int_Sl_FreqResourceCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6929}, false); err != nil {
			return utils.WrapError("Decode Sl_FreqResourceCG_Type1_r16", err)
		}
		ie.Sl_FreqResourceCG_Type1_r16 = &tmp_int_Sl_FreqResourceCG_Type1_r16
	}
	if Sl_TimeOffsetCG_Type1_r16Present {
		var tmp_int_Sl_TimeOffsetCG_Type1_r16 int64
		if tmp_int_Sl_TimeOffsetCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7999}, false); err != nil {
			return utils.WrapError("Decode Sl_TimeOffsetCG_Type1_r16", err)
		}
		ie.Sl_TimeOffsetCG_Type1_r16 = &tmp_int_Sl_TimeOffsetCG_Type1_r16
	}
	if Sl_N1PUCCH_AN_r16Present {
		ie.Sl_N1PUCCH_AN_r16 = new(PUCCH_ResourceId)
		if err = ie.Sl_N1PUCCH_AN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_N1PUCCH_AN_r16", err)
		}
	}
	if Sl_PSFCH_ToPUCCH_CG_Type1_r16Present {
		var tmp_int_Sl_PSFCH_ToPUCCH_CG_Type1_r16 int64
		if tmp_int_Sl_PSFCH_ToPUCCH_CG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Sl_PSFCH_ToPUCCH_CG_Type1_r16", err)
		}
		ie.Sl_PSFCH_ToPUCCH_CG_Type1_r16 = &tmp_int_Sl_PSFCH_ToPUCCH_CG_Type1_r16
	}
	if Sl_ResourcePoolID_r16Present {
		ie.Sl_ResourcePoolID_r16 = new(SL_ResourcePoolID_r16)
		if err = ie.Sl_ResourcePoolID_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ResourcePoolID_r16", err)
		}
	}
	if Sl_TimeReferenceSFN_Type1_r16Present {
		ie.Sl_TimeReferenceSFN_Type1_r16 = new(SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16_sl_TimeReferenceSFN_Type1_r16)
		if err = ie.Sl_TimeReferenceSFN_Type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TimeReferenceSFN_Type1_r16", err)
		}
	}
	return nil
}
