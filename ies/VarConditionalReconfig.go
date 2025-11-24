package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarConditionalReconfig struct {
	CondReconfigList *CondReconfigToAddModList_r16 `optional`
}

func (ie *VarConditionalReconfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CondReconfigList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CondReconfigList != nil {
		if err = ie.CondReconfigList.Encode(w); err != nil {
			return utils.WrapError("Encode CondReconfigList", err)
		}
	}
	return nil
}

func (ie *VarConditionalReconfig) Decode(r *uper.UperReader) error {
	var err error
	var CondReconfigListPresent bool
	if CondReconfigListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CondReconfigListPresent {
		ie.CondReconfigList = new(CondReconfigToAddModList_r16)
		if err = ie.CondReconfigList.Decode(r); err != nil {
			return utils.WrapError("Decode CondReconfigList", err)
		}
	}
	return nil
}
