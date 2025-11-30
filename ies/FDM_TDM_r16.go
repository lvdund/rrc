package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FDM_TDM_r16 struct {
	RepetitionScheme_r16      FDM_TDM_r16_repetitionScheme_r16 `madatory`
	StartingSymbolOffsetK_r16 *int64                           `lb:0,ub:7,optional`
}

func (ie *FDM_TDM_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.StartingSymbolOffsetK_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.RepetitionScheme_r16.Encode(w); err != nil {
		return utils.WrapError("Encode RepetitionScheme_r16", err)
	}
	if ie.StartingSymbolOffsetK_r16 != nil {
		if err = w.WriteInteger(*ie.StartingSymbolOffsetK_r16, &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Encode StartingSymbolOffsetK_r16", err)
		}
	}
	return nil
}

func (ie *FDM_TDM_r16) Decode(r *aper.AperReader) error {
	var err error
	var StartingSymbolOffsetK_r16Present bool
	if StartingSymbolOffsetK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.RepetitionScheme_r16.Decode(r); err != nil {
		return utils.WrapError("Decode RepetitionScheme_r16", err)
	}
	if StartingSymbolOffsetK_r16Present {
		var tmp_int_StartingSymbolOffsetK_r16 int64
		if tmp_int_StartingSymbolOffsetK_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode StartingSymbolOffsetK_r16", err)
		}
		ie.StartingSymbolOffsetK_r16 = &tmp_int_StartingSymbolOffsetK_r16
	}
	return nil
}
