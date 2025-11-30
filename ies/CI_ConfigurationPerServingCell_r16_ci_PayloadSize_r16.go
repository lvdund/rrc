package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n1   aper.Enumerated = 0
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n2   aper.Enumerated = 1
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n4   aper.Enumerated = 2
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n5   aper.Enumerated = 3
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n7   aper.Enumerated = 4
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n8   aper.Enumerated = 5
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n10  aper.Enumerated = 6
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n14  aper.Enumerated = 7
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n16  aper.Enumerated = 8
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n20  aper.Enumerated = 9
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n28  aper.Enumerated = 10
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n32  aper.Enumerated = 11
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n35  aper.Enumerated = 12
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n42  aper.Enumerated = 13
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n56  aper.Enumerated = 14
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n112 aper.Enumerated = 15
)

type CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16", err)
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
