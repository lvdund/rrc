package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n2 aper.Enumerated = 0
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n4 aper.Enumerated = 1
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n6 aper.Enumerated = 2
	PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n8 aper.Enumerated = 3
)

type PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock", err)
	}
	return nil
}

func (ie *PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PUSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
