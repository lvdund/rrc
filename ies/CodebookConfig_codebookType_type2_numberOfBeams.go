package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_numberOfBeams_Enum_two   aper.Enumerated = 0
	CodebookConfig_codebookType_type2_numberOfBeams_Enum_three aper.Enumerated = 1
	CodebookConfig_codebookType_type2_numberOfBeams_Enum_four  aper.Enumerated = 2
)

type CodebookConfig_codebookType_type2_numberOfBeams struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type2_numberOfBeams) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CodebookConfig_codebookType_type2_numberOfBeams", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_numberOfBeams) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CodebookConfig_codebookType_type2_numberOfBeams", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
