package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_phaseAlphabetSize_Enum_n4 aper.Enumerated = 0
	CodebookConfig_codebookType_type2_phaseAlphabetSize_Enum_n8 aper.Enumerated = 1
)

type CodebookConfig_codebookType_type2_phaseAlphabetSize struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CodebookConfig_codebookType_type2_phaseAlphabetSize) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CodebookConfig_codebookType_type2_phaseAlphabetSize", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_phaseAlphabetSize) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CodebookConfig_codebookType_type2_phaseAlphabetSize", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
