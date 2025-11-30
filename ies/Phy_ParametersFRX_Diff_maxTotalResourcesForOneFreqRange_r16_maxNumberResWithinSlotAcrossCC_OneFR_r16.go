package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n2   aper.Enumerated = 0
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n4   aper.Enumerated = 1
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n8   aper.Enumerated = 2
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n12  aper.Enumerated = 3
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n16  aper.Enumerated = 4
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n32  aper.Enumerated = 5
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n64  aper.Enumerated = 6
	Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16_Enum_n128 aper.Enumerated = 7
)

type Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16_maxNumberResWithinSlotAcrossCC_OneFR_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
