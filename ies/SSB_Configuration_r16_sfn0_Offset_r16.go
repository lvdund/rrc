package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_Configuration_r16_sfn0_Offset_r16 struct {
	Sfn_Offset_r16            int64  `lb:0,ub:1023,madatory`
	IntegerSubframeOffset_r16 *int64 `lb:0,ub:9,optional`
}

func (ie *SSB_Configuration_r16_sfn0_Offset_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.IntegerSubframeOffset_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Sfn_Offset_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger Sfn_Offset_r16", err)
	}
	if ie.IntegerSubframeOffset_r16 != nil {
		if err = w.WriteInteger(*ie.IntegerSubframeOffset_r16, &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode IntegerSubframeOffset_r16", err)
		}
	}
	return nil
}

func (ie *SSB_Configuration_r16_sfn0_Offset_r16) Decode(r *uper.UperReader) error {
	var err error
	var IntegerSubframeOffset_r16Present bool
	if IntegerSubframeOffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Sfn_Offset_r16 int64
	if tmp_int_Sfn_Offset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger Sfn_Offset_r16", err)
	}
	ie.Sfn_Offset_r16 = tmp_int_Sfn_Offset_r16
	if IntegerSubframeOffset_r16Present {
		var tmp_int_IntegerSubframeOffset_r16 int64
		if tmp_int_IntegerSubframeOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode IntegerSubframeOffset_r16", err)
		}
		ie.IntegerSubframeOffset_r16 = &tmp_int_IntegerSubframeOffset_r16
	}
	return nil
}
