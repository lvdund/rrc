package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig struct {
	CodebookType *CodebookConfig_codebookType `lb:6,ub:6,optional`
}

func (ie *CodebookConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CodebookType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CodebookType != nil {
		if err = ie.CodebookType.Encode(w); err != nil {
			return utils.WrapError("Encode CodebookType", err)
		}
	}
	return nil
}

func (ie *CodebookConfig) Decode(r *uper.UperReader) error {
	var err error
	var CodebookTypePresent bool
	if CodebookTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CodebookTypePresent {
		ie.CodebookType = new(CodebookConfig_codebookType)
		if err = ie.CodebookType.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookType", err)
		}
	}
	return nil
}
