package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DelayBudgetReport_type1_Enum_msMinus1280 aper.Enumerated = 0
	DelayBudgetReport_type1_Enum_msMinus640  aper.Enumerated = 1
	DelayBudgetReport_type1_Enum_msMinus320  aper.Enumerated = 2
	DelayBudgetReport_type1_Enum_msMinus160  aper.Enumerated = 3
	DelayBudgetReport_type1_Enum_msMinus80   aper.Enumerated = 4
	DelayBudgetReport_type1_Enum_msMinus60   aper.Enumerated = 5
	DelayBudgetReport_type1_Enum_msMinus40   aper.Enumerated = 6
	DelayBudgetReport_type1_Enum_msMinus20   aper.Enumerated = 7
	DelayBudgetReport_type1_Enum_ms0         aper.Enumerated = 8
	DelayBudgetReport_type1_Enum_ms20        aper.Enumerated = 9
	DelayBudgetReport_type1_Enum_ms40        aper.Enumerated = 10
	DelayBudgetReport_type1_Enum_ms60        aper.Enumerated = 11
	DelayBudgetReport_type1_Enum_ms80        aper.Enumerated = 12
	DelayBudgetReport_type1_Enum_ms160       aper.Enumerated = 13
	DelayBudgetReport_type1_Enum_ms320       aper.Enumerated = 14
	DelayBudgetReport_type1_Enum_ms640       aper.Enumerated = 15
	DelayBudgetReport_type1_Enum_ms1280      aper.Enumerated = 16
)

type DelayBudgetReport_type1 struct {
	Value aper.Enumerated `lb:0,ub:16,madatory`
}

func (ie *DelayBudgetReport_type1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode DelayBudgetReport_type1", err)
	}
	return nil
}

func (ie *DelayBudgetReport_type1) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode DelayBudgetReport_type1", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
