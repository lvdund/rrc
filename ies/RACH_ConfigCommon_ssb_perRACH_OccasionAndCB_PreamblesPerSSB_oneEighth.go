package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n4  aper.Enumerated = 0
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n8  aper.Enumerated = 1
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n12 aper.Enumerated = 2
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n16 aper.Enumerated = 3
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n20 aper.Enumerated = 4
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n24 aper.Enumerated = 5
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n28 aper.Enumerated = 6
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n32 aper.Enumerated = 7
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n36 aper.Enumerated = 8
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n40 aper.Enumerated = 9
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n44 aper.Enumerated = 10
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n48 aper.Enumerated = 11
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n52 aper.Enumerated = 12
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n56 aper.Enumerated = 13
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n60 aper.Enumerated = 14
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth_Enum_n64 aper.Enumerated = 15
)

type RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
