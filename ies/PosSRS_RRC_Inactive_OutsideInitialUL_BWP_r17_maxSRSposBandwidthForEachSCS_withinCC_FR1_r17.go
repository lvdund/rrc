package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw5   aper.Enumerated = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw10  aper.Enumerated = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw15  aper.Enumerated = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw20  aper.Enumerated = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw25  aper.Enumerated = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw30  aper.Enumerated = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw35  aper.Enumerated = 6
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw40  aper.Enumerated = 7
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw45  aper.Enumerated = 8
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw50  aper.Enumerated = 9
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw60  aper.Enumerated = 10
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw70  aper.Enumerated = 11
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw80  aper.Enumerated = 12
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw90  aper.Enumerated = 13
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw100 aper.Enumerated = 14
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17 struct {
	Value aper.Enumerated `lb:0,ub:14,madatory`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
