package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n2   aper.Enumerated = 0
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n4   aper.Enumerated = 1
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n8   aper.Enumerated = 2
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n12  aper.Enumerated = 3
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n16  aper.Enumerated = 4
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n32  aper.Enumerated = 5
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n40  aper.Enumerated = 6
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n48  aper.Enumerated = 7
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n64  aper.Enumerated = 8
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n72  aper.Enumerated = 9
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n80  aper.Enumerated = 10
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n96  aper.Enumerated = 11
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n128 aper.Enumerated = 12
	Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16_Enum_n256 aper.Enumerated = 13
)

type Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16 struct {
	Value aper.Enumerated `lb:0,ub:13,madatory`
}

func (ie *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16", err)
	}
	return nil
}

func (ie *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16_maxNumberResAcrossCC_AcrossFR_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
