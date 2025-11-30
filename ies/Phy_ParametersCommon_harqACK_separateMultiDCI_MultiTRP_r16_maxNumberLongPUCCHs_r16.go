package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16_Enum_longAndLong   aper.Enumerated = 0
	Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16_Enum_longAndShort  aper.Enumerated = 1
	Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16_Enum_shortAndShort aper.Enumerated = 2
)

type Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16", err)
	}
	return nil
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
