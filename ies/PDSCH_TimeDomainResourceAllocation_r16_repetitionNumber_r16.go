package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n2  aper.Enumerated = 0
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n3  aper.Enumerated = 1
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n4  aper.Enumerated = 2
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n5  aper.Enumerated = 3
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n6  aper.Enumerated = 4
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n7  aper.Enumerated = 5
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n8  aper.Enumerated = 6
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16_Enum_n16 aper.Enumerated = 7
)

type PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16", err)
	}
	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
