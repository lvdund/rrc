package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_v1730_codebookType_type1 struct {
	CodebookMode *int64 `lb:1,ub:2,optional`
}

func (ie *CodebookConfig_v1730_codebookType_type1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CodebookMode != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CodebookMode != nil {
		if err = w.WriteInteger(*ie.CodebookMode, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode CodebookMode", err)
		}
	}
	return nil
}

func (ie *CodebookConfig_v1730_codebookType_type1) Decode(r *uper.UperReader) error {
	var err error
	var CodebookModePresent bool
	if CodebookModePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CodebookModePresent {
		var tmp_int_CodebookMode int64
		if tmp_int_CodebookMode, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode CodebookMode", err)
		}
		ie.CodebookMode = &tmp_int_CodebookMode
	}
	return nil
}
