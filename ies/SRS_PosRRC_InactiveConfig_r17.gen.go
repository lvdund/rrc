// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-PosRRC-InactiveConfig-r17 (line 1434).

var sRSPosRRCInactiveConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PosConfigNUL-r17", Optional: true},
		{Name: "srs-PosConfigSUL-r17", Optional: true},
		{Name: "bwp-NUL-r17", Optional: true},
		{Name: "bwp-SUL-r17", Optional: true},
		{Name: "inactivePosSRS-TimeAlignmentTimer-r17", Optional: true},
		{Name: "inactivePosSRS-RSRP-ChangeThreshold-r17", Optional: true},
	},
}

type SRS_PosRRC_InactiveConfig_r17 struct {
	Srs_PosConfigNUL_r17                    *SRS_PosConfig_r17
	Srs_PosConfigSUL_r17                    *SRS_PosConfig_r17
	Bwp_NUL_r17                             *BWP
	Bwp_SUL_r17                             *BWP
	InactivePosSRS_TimeAlignmentTimer_r17   *TimeAlignmentTimer
	InactivePosSRS_RSRP_ChangeThreshold_r17 *RSRP_ChangeThreshold_r17
}

func (ie *SRS_PosRRC_InactiveConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSPosRRCInactiveConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PosConfigNUL_r17 != nil, ie.Srs_PosConfigSUL_r17 != nil, ie.Bwp_NUL_r17 != nil, ie.Bwp_SUL_r17 != nil, ie.InactivePosSRS_TimeAlignmentTimer_r17 != nil, ie.InactivePosSRS_RSRP_ChangeThreshold_r17 != nil}); err != nil {
		return err
	}

	// 2. srs-PosConfigNUL-r17: ref
	{
		if ie.Srs_PosConfigNUL_r17 != nil {
			if err := ie.Srs_PosConfigNUL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. srs-PosConfigSUL-r17: ref
	{
		if ie.Srs_PosConfigSUL_r17 != nil {
			if err := ie.Srs_PosConfigSUL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. bwp-NUL-r17: ref
	{
		if ie.Bwp_NUL_r17 != nil {
			if err := ie.Bwp_NUL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. bwp-SUL-r17: ref
	{
		if ie.Bwp_SUL_r17 != nil {
			if err := ie.Bwp_SUL_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. inactivePosSRS-TimeAlignmentTimer-r17: ref
	{
		if ie.InactivePosSRS_TimeAlignmentTimer_r17 != nil {
			if err := ie.InactivePosSRS_TimeAlignmentTimer_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. inactivePosSRS-RSRP-ChangeThreshold-r17: ref
	{
		if ie.InactivePosSRS_RSRP_ChangeThreshold_r17 != nil {
			if err := ie.InactivePosSRS_RSRP_ChangeThreshold_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SRS_PosRRC_InactiveConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSPosRRCInactiveConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-PosConfigNUL-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_PosConfigNUL_r17 = new(SRS_PosConfig_r17)
			if err := ie.Srs_PosConfigNUL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. srs-PosConfigSUL-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_PosConfigSUL_r17 = new(SRS_PosConfig_r17)
			if err := ie.Srs_PosConfigSUL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. bwp-NUL-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Bwp_NUL_r17 = new(BWP)
			if err := ie.Bwp_NUL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. bwp-SUL-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Bwp_SUL_r17 = new(BWP)
			if err := ie.Bwp_SUL_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. inactivePosSRS-TimeAlignmentTimer-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.InactivePosSRS_TimeAlignmentTimer_r17 = new(TimeAlignmentTimer)
			if err := ie.InactivePosSRS_TimeAlignmentTimer_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. inactivePosSRS-RSRP-ChangeThreshold-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.InactivePosSRS_RSRP_ChangeThreshold_r17 = new(RSRP_ChangeThreshold_r17)
			if err := ie.InactivePosSRS_RSRP_ChangeThreshold_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
