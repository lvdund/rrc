package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_v1730 struct {
	CodebookType *CodebookConfig_v1730_codebookType `lb:1,ub:2,optional`
}

func (ie *CodebookConfig_v1730) Encode(w *aper.AperWriter) error {
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

func (ie *CodebookConfig_v1730) Decode(r *aper.AperReader) error {
	var err error
	var CodebookTypePresent bool
	if CodebookTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CodebookTypePresent {
		ie.CodebookType = new(CodebookConfig_v1730_codebookType)
		if err = ie.CodebookType.Decode(r); err != nil {
			return utils.WrapError("Decode CodebookType", err)
		}
	}
	return nil
}
