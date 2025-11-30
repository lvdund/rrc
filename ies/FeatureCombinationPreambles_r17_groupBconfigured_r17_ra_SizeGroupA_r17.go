package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b56    aper.Enumerated = 0
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b144   aper.Enumerated = 1
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b208   aper.Enumerated = 2
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b256   aper.Enumerated = 3
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b282   aper.Enumerated = 4
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b480   aper.Enumerated = 5
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b640   aper.Enumerated = 6
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b800   aper.Enumerated = 7
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b1000  aper.Enumerated = 8
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b72    aper.Enumerated = 9
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare6 aper.Enumerated = 10
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare5 aper.Enumerated = 11
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare4 aper.Enumerated = 12
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare3 aper.Enumerated = 13
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare2 aper.Enumerated = 14
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare1 aper.Enumerated = 15
)

type FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17", err)
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
