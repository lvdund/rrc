package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n1  aper.Enumerated = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n2  aper.Enumerated = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n3  aper.Enumerated = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n4  aper.Enumerated = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n5  aper.Enumerated = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n6  aper.Enumerated = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n8  aper.Enumerated = 6
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n10 aper.Enumerated = 7
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n12 aper.Enumerated = 8
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17_Enum_n14 aper.Enumerated = 9
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17 struct {
	Value aper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17", err)
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
