package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxCC_Preference_r16 struct {
	ReducedMaxCCs_r16 *ReducedMaxCCs_r16 `optional`
}

func (ie *MaxCC_Preference_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedMaxCCs_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedMaxCCs_r16 != nil {
		if err = ie.ReducedMaxCCs_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedMaxCCs_r16", err)
		}
	}
	return nil
}

func (ie *MaxCC_Preference_r16) Decode(r *uper.UperReader) error {
	var err error
	var ReducedMaxCCs_r16Present bool
	if ReducedMaxCCs_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedMaxCCs_r16Present {
		ie.ReducedMaxCCs_r16 = new(ReducedMaxCCs_r16)
		if err = ie.ReducedMaxCCs_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedMaxCCs_r16", err)
		}
	}
	return nil
}
