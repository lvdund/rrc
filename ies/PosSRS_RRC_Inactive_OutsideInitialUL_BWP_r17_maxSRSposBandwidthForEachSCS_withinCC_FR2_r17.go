package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17_Enum_bw50  aper.Enumerated = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17_Enum_bw100 aper.Enumerated = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17_Enum_bw200 aper.Enumerated = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17_Enum_bw400 aper.Enumerated = 3
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17", err)
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
