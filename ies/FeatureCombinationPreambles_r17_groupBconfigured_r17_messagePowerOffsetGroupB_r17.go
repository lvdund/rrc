package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_minusinfinity aper.Enumerated = 0
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB0           aper.Enumerated = 1
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB5           aper.Enumerated = 2
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB8           aper.Enumerated = 3
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB10          aper.Enumerated = 4
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB12          aper.Enumerated = 5
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB15          aper.Enumerated = 6
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB18          aper.Enumerated = 7
)

type FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17", err)
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
