package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n2 aper.Enumerated = 0
	PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n4 aper.Enumerated = 1
	PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n6 aper.Enumerated = 2
	PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock_Enum_n8 aper.Enumerated = 3
)

type PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock", err)
	}
	return nil
}

func (ie *PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
