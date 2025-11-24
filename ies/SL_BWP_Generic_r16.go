package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_BWP_Generic_r16 struct {
	Sl_BWP_r16                     *BWP                                     `optional`
	Sl_LengthSymbols_r16           *SL_BWP_Generic_r16_sl_LengthSymbols_r16 `optional`
	Sl_StartSymbol_r16             *SL_BWP_Generic_r16_sl_StartSymbol_r16   `optional`
	Sl_PSBCH_Config_r16            *SL_PSBCH_Config_r16                     `optional,setuprelease`
	Sl_TxDirectCurrentLocation_r16 *int64                                   `lb:0,ub:3301,optional`
}

func (ie *SL_BWP_Generic_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_BWP_r16 != nil, ie.Sl_LengthSymbols_r16 != nil, ie.Sl_StartSymbol_r16 != nil, ie.Sl_PSBCH_Config_r16 != nil, ie.Sl_TxDirectCurrentLocation_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_BWP_r16 != nil {
		if err = ie.Sl_BWP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_BWP_r16", err)
		}
	}
	if ie.Sl_LengthSymbols_r16 != nil {
		if err = ie.Sl_LengthSymbols_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_LengthSymbols_r16", err)
		}
	}
	if ie.Sl_StartSymbol_r16 != nil {
		if err = ie.Sl_StartSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_StartSymbol_r16", err)
		}
	}
	if ie.Sl_PSBCH_Config_r16 != nil {
		tmp_Sl_PSBCH_Config_r16 := utils.SetupRelease[*SL_PSBCH_Config_r16]{
			Setup: ie.Sl_PSBCH_Config_r16,
		}
		if err = tmp_Sl_PSBCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PSBCH_Config_r16", err)
		}
	}
	if ie.Sl_TxDirectCurrentLocation_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_TxDirectCurrentLocation_r16, &uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
			return utils.WrapError("Encode Sl_TxDirectCurrentLocation_r16", err)
		}
	}
	return nil
}

func (ie *SL_BWP_Generic_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_BWP_r16Present bool
	if Sl_BWP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_LengthSymbols_r16Present bool
	if Sl_LengthSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_StartSymbol_r16Present bool
	if Sl_StartSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PSBCH_Config_r16Present bool
	if Sl_PSBCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_TxDirectCurrentLocation_r16Present bool
	if Sl_TxDirectCurrentLocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_BWP_r16Present {
		ie.Sl_BWP_r16 = new(BWP)
		if err = ie.Sl_BWP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_BWP_r16", err)
		}
	}
	if Sl_LengthSymbols_r16Present {
		ie.Sl_LengthSymbols_r16 = new(SL_BWP_Generic_r16_sl_LengthSymbols_r16)
		if err = ie.Sl_LengthSymbols_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_LengthSymbols_r16", err)
		}
	}
	if Sl_StartSymbol_r16Present {
		ie.Sl_StartSymbol_r16 = new(SL_BWP_Generic_r16_sl_StartSymbol_r16)
		if err = ie.Sl_StartSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_StartSymbol_r16", err)
		}
	}
	if Sl_PSBCH_Config_r16Present {
		tmp_Sl_PSBCH_Config_r16 := utils.SetupRelease[*SL_PSBCH_Config_r16]{}
		if err = tmp_Sl_PSBCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PSBCH_Config_r16", err)
		}
		ie.Sl_PSBCH_Config_r16 = tmp_Sl_PSBCH_Config_r16.Setup
	}
	if Sl_TxDirectCurrentLocation_r16Present {
		var tmp_int_Sl_TxDirectCurrentLocation_r16 int64
		if tmp_int_Sl_TxDirectCurrentLocation_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
			return utils.WrapError("Decode Sl_TxDirectCurrentLocation_r16", err)
		}
		ie.Sl_TxDirectCurrentLocation_r16 = &tmp_int_Sl_TxDirectCurrentLocation_r16
	}
	return nil
}
