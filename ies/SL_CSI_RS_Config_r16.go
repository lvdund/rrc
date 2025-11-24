package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_CSI_RS_Config_r16 struct {
	Sl_CSI_RS_FreqAllocation_r16 *SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16 `lb:12,ub:12,optional`
	Sl_CSI_RS_FirstSymbol_r16    *int64                                             `lb:3,ub:12,optional`
}

func (ie *SL_CSI_RS_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_CSI_RS_FreqAllocation_r16 != nil, ie.Sl_CSI_RS_FirstSymbol_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_CSI_RS_FreqAllocation_r16 != nil {
		if err = ie.Sl_CSI_RS_FreqAllocation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CSI_RS_FreqAllocation_r16", err)
		}
	}
	if ie.Sl_CSI_RS_FirstSymbol_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_CSI_RS_FirstSymbol_r16, &uper.Constraint{Lb: 3, Ub: 12}, false); err != nil {
			return utils.WrapError("Encode Sl_CSI_RS_FirstSymbol_r16", err)
		}
	}
	return nil
}

func (ie *SL_CSI_RS_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_CSI_RS_FreqAllocation_r16Present bool
	if Sl_CSI_RS_FreqAllocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CSI_RS_FirstSymbol_r16Present bool
	if Sl_CSI_RS_FirstSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_CSI_RS_FreqAllocation_r16Present {
		ie.Sl_CSI_RS_FreqAllocation_r16 = new(SL_CSI_RS_Config_r16_sl_CSI_RS_FreqAllocation_r16)
		if err = ie.Sl_CSI_RS_FreqAllocation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CSI_RS_FreqAllocation_r16", err)
		}
	}
	if Sl_CSI_RS_FirstSymbol_r16Present {
		var tmp_int_Sl_CSI_RS_FirstSymbol_r16 int64
		if tmp_int_Sl_CSI_RS_FirstSymbol_r16, err = r.ReadInteger(&uper.Constraint{Lb: 3, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode Sl_CSI_RS_FirstSymbol_r16", err)
		}
		ie.Sl_CSI_RS_FirstSymbol_r16 = &tmp_int_Sl_CSI_RS_FirstSymbol_r16
	}
	return nil
}
