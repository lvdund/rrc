package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_Enum_semiStatic aper.Enumerated = 0
	PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_Enum_dynamic    aper.Enumerated = 1
)

type PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
