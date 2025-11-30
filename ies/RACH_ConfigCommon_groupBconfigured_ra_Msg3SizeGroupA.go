package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b56    aper.Enumerated = 0
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b144   aper.Enumerated = 1
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b208   aper.Enumerated = 2
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b256   aper.Enumerated = 3
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b282   aper.Enumerated = 4
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b480   aper.Enumerated = 5
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b640   aper.Enumerated = 6
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b800   aper.Enumerated = 7
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b1000  aper.Enumerated = 8
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b72    aper.Enumerated = 9
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare6 aper.Enumerated = 10
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare5 aper.Enumerated = 11
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare4 aper.Enumerated = 12
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare3 aper.Enumerated = 13
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare2 aper.Enumerated = 14
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare1 aper.Enumerated = 15
)

type RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
