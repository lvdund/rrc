package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1 struct {
	SubType      *CodebookConfig_codebookType_type1_subType `lb:6,ub:6,optional`
	CodebookMode int64                                      `lb:1,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type1) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SubType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SubType != nil {
		if err = ie.SubType.Encode(w); err != nil {
			return utils.WrapError("Encode SubType", err)
		}
	}
	if err = w.WriteInteger(ie.CodebookMode, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger CodebookMode", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1) Decode(r *aper.AperReader) error {
	var err error
	var SubTypePresent bool
	if SubTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SubTypePresent {
		ie.SubType = new(CodebookConfig_codebookType_type1_subType)
		if err = ie.SubType.Decode(r); err != nil {
			return utils.WrapError("Decode SubType", err)
		}
	}
	var tmp_int_CodebookMode int64
	if tmp_int_CodebookMode, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger CodebookMode", err)
	}
	ie.CodebookMode = tmp_int_CodebookMode
	return nil
}
