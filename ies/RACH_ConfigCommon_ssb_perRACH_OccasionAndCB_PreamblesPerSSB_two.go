package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n4  aper.Enumerated = 0
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n8  aper.Enumerated = 1
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n12 aper.Enumerated = 2
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n16 aper.Enumerated = 3
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n20 aper.Enumerated = 4
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n24 aper.Enumerated = 5
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n28 aper.Enumerated = 6
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two_Enum_n32 aper.Enumerated = 7
)

type RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
