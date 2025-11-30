package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommon_groupBconfigured struct {
	Ra_Msg3SizeGroupA          RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA        `madatory`
	MessagePowerOffsetGroupB   RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB `madatory`
	NumberOfRA_PreamblesGroupA int64                                                       `lb:1,ub:64,madatory`
}

func (ie *RACH_ConfigCommon_groupBconfigured) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ra_Msg3SizeGroupA.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_Msg3SizeGroupA", err)
	}
	if err = ie.MessagePowerOffsetGroupB.Encode(w); err != nil {
		return utils.WrapError("Encode MessagePowerOffsetGroupB", err)
	}
	if err = w.WriteInteger(ie.NumberOfRA_PreamblesGroupA, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfRA_PreamblesGroupA", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_groupBconfigured) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ra_Msg3SizeGroupA.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_Msg3SizeGroupA", err)
	}
	if err = ie.MessagePowerOffsetGroupB.Decode(r); err != nil {
		return utils.WrapError("Decode MessagePowerOffsetGroupB", err)
	}
	var tmp_int_NumberOfRA_PreamblesGroupA int64
	if tmp_int_NumberOfRA_PreamblesGroupA, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfRA_PreamblesGroupA", err)
	}
	ie.NumberOfRA_PreamblesGroupA = tmp_int_NumberOfRA_PreamblesGroupA
	return nil
}
