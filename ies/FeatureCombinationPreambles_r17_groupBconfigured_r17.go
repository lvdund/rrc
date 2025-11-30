package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureCombinationPreambles_r17_groupBconfigured_r17 struct {
	Ra_SizeGroupA_r17              FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17            `madatory`
	MessagePowerOffsetGroupB_r17   FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17 `madatory`
	NumberOfRA_PreamblesGroupA_r17 int64                                                                             `lb:1,ub:64,madatory`
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ra_SizeGroupA_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_SizeGroupA_r17", err)
	}
	if err = ie.MessagePowerOffsetGroupB_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MessagePowerOffsetGroupB_r17", err)
	}
	if err = w.WriteInteger(ie.NumberOfRA_PreamblesGroupA_r17, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfRA_PreamblesGroupA_r17", err)
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ra_SizeGroupA_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_SizeGroupA_r17", err)
	}
	if err = ie.MessagePowerOffsetGroupB_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MessagePowerOffsetGroupB_r17", err)
	}
	var tmp_int_NumberOfRA_PreamblesGroupA_r17 int64
	if tmp_int_NumberOfRA_PreamblesGroupA_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfRA_PreamblesGroupA_r17", err)
	}
	ie.NumberOfRA_PreamblesGroupA_r17 = tmp_int_NumberOfRA_PreamblesGroupA_r17
	return nil
}
