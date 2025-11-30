package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r16 struct {
	CodebookType CodebookConfig_r16_codebookType `lb:16,ub:16,madatory`
}

func (ie *CodebookConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CodebookType.Encode(w); err != nil {
		return utils.WrapError("Encode CodebookType", err)
	}
	return nil
}

func (ie *CodebookConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CodebookType.Decode(r); err != nil {
		return utils.WrapError("Decode CodebookType", err)
	}
	return nil
}
