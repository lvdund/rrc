package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD_Enum_supported aper.Enumerated = 0
)

type CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD", err)
	}
	return nil
}

func (ie *CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CarrierAggregationVariant_fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
