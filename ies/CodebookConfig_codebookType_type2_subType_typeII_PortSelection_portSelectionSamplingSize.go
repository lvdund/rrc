package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n1 aper.Enumerated = 0
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n2 aper.Enumerated = 1
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n3 aper.Enumerated = 2
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n4 aper.Enumerated = 3
)

type CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
