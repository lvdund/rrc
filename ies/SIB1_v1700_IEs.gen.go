// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB1-v1700-IEs (line 2023).

var sIB1V1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "hsdn-Cell-r17", Optional: true},
		{Name: "uac-BarringInfo-v1700", Optional: true},
		{Name: "sdt-ConfigCommon-r17", Optional: true},
		{Name: "redCap-ConfigCommon-r17", Optional: true},
		{Name: "featurePriorities-r17", Optional: true},
		{Name: "si-SchedulingInfo-v1700", Optional: true},
		{Name: "hyperSFN-r17", Optional: true},
		{Name: "eDRX-AllowedIdle-r17", Optional: true},
		{Name: "eDRX-AllowedInactive-r17", Optional: true},
		{Name: "intraFreqReselectionRedCap-r17", Optional: true},
		{Name: "cellBarredNTN-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	SIB1_v1700_IEs_Hsdn_Cell_r17_True = 0
)

var sIB1V1700IEsHsdnCellR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1700_IEs_Hsdn_Cell_r17_True},
}

var sIB1V1700IEsHyperSFNR17Constraints = per.FixedSize(10)

const (
	SIB1_v1700_IEs_EDRX_AllowedIdle_r17_True = 0
)

var sIB1V1700IEsEDRXAllowedIdleR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1700_IEs_EDRX_AllowedIdle_r17_True},
}

const (
	SIB1_v1700_IEs_EDRX_AllowedInactive_r17_True = 0
)

var sIB1V1700IEsEDRXAllowedInactiveR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1700_IEs_EDRX_AllowedInactive_r17_True},
}

const (
	SIB1_v1700_IEs_IntraFreqReselectionRedCap_r17_Allowed    = 0
	SIB1_v1700_IEs_IntraFreqReselectionRedCap_r17_NotAllowed = 1
)

var sIB1V1700IEsIntraFreqReselectionRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1700_IEs_IntraFreqReselectionRedCap_r17_Allowed, SIB1_v1700_IEs_IntraFreqReselectionRedCap_r17_NotAllowed},
}

const (
	SIB1_v1700_IEs_CellBarredNTN_r17_Barred    = 0
	SIB1_v1700_IEs_CellBarredNTN_r17_NotBarred = 1
)

var sIB1V1700IEsCellBarredNTNR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1700_IEs_CellBarredNTN_r17_Barred, SIB1_v1700_IEs_CellBarredNTN_r17_NotBarred},
}

var sIB1V1700IEsFeaturePrioritiesR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "redCapPriority-r17", Optional: true},
		{Name: "slicingPriority-r17", Optional: true},
		{Name: "msg3-Repetitions-Priority-r17", Optional: true},
		{Name: "sdt-Priority-r17", Optional: true},
	},
}

type SIB1_v1700_IEs struct {
	Hsdn_Cell_r17           *int64
	Uac_BarringInfo_v1700   *struct{ Uac_BarringInfoSetList_v1700 UAC_BarringInfoSetList_v1700 }
	Sdt_ConfigCommon_r17    *SDT_ConfigCommonSIB_r17
	RedCap_ConfigCommon_r17 *RedCap_ConfigCommonSIB_r17
	FeaturePriorities_r17   *struct {
		RedCapPriority_r17            *FeaturePriority_r17
		SlicingPriority_r17           *FeaturePriority_r17
		Msg3_Repetitions_Priority_r17 *FeaturePriority_r17
		Sdt_Priority_r17              *FeaturePriority_r17
	}
	Si_SchedulingInfo_v1700        *SI_SchedulingInfo_v1700
	HyperSFN_r17                   *per.BitString
	EDRX_AllowedIdle_r17           *int64
	EDRX_AllowedInactive_r17       *int64
	IntraFreqReselectionRedCap_r17 *int64
	CellBarredNTN_r17              *int64
	NonCriticalExtension           *SIB1_v1740_IEs
}

func (ie *SIB1_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1V1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Hsdn_Cell_r17 != nil, ie.Uac_BarringInfo_v1700 != nil, ie.Sdt_ConfigCommon_r17 != nil, ie.RedCap_ConfigCommon_r17 != nil, ie.FeaturePriorities_r17 != nil, ie.Si_SchedulingInfo_v1700 != nil, ie.HyperSFN_r17 != nil, ie.EDRX_AllowedIdle_r17 != nil, ie.EDRX_AllowedInactive_r17 != nil, ie.IntraFreqReselectionRedCap_r17 != nil, ie.CellBarredNTN_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. hsdn-Cell-r17: enumerated
	{
		if ie.Hsdn_Cell_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Hsdn_Cell_r17, sIB1V1700IEsHsdnCellR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. uac-BarringInfo-v1700: sequence
	{
		if ie.Uac_BarringInfo_v1700 != nil {
			c := ie.Uac_BarringInfo_v1700
			if err := c.Uac_BarringInfoSetList_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sdt-ConfigCommon-r17: ref
	{
		if ie.Sdt_ConfigCommon_r17 != nil {
			if err := ie.Sdt_ConfigCommon_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. redCap-ConfigCommon-r17: ref
	{
		if ie.RedCap_ConfigCommon_r17 != nil {
			if err := ie.RedCap_ConfigCommon_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. featurePriorities-r17: sequence
	{
		if ie.FeaturePriorities_r17 != nil {
			c := ie.FeaturePriorities_r17
			sIB1V1700IEsFeaturePrioritiesR17Seq := e.NewSequenceEncoder(sIB1V1700IEsFeaturePrioritiesR17Constraints)
			if err := sIB1V1700IEsFeaturePrioritiesR17Seq.EncodePreamble([]bool{c.RedCapPriority_r17 != nil, c.SlicingPriority_r17 != nil, c.Msg3_Repetitions_Priority_r17 != nil, c.Sdt_Priority_r17 != nil}); err != nil {
				return err
			}
			if c.RedCapPriority_r17 != nil {
				if err := c.RedCapPriority_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.SlicingPriority_r17 != nil {
				if err := c.SlicingPriority_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Msg3_Repetitions_Priority_r17 != nil {
				if err := c.Msg3_Repetitions_Priority_r17.Encode(e); err != nil {
					return err
				}
			}
			if c.Sdt_Priority_r17 != nil {
				if err := c.Sdt_Priority_r17.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. si-SchedulingInfo-v1700: ref
	{
		if ie.Si_SchedulingInfo_v1700 != nil {
			if err := ie.Si_SchedulingInfo_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. hyperSFN-r17: bit-string
	{
		if ie.HyperSFN_r17 != nil {
			if err := e.EncodeBitString(*ie.HyperSFN_r17, sIB1V1700IEsHyperSFNR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. eDRX-AllowedIdle-r17: enumerated
	{
		if ie.EDRX_AllowedIdle_r17 != nil {
			if err := e.EncodeEnumerated(*ie.EDRX_AllowedIdle_r17, sIB1V1700IEsEDRXAllowedIdleR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. eDRX-AllowedInactive-r17: enumerated
	{
		if ie.EDRX_AllowedInactive_r17 != nil {
			if err := e.EncodeEnumerated(*ie.EDRX_AllowedInactive_r17, sIB1V1700IEsEDRXAllowedInactiveR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. intraFreqReselectionRedCap-r17: enumerated
	{
		if ie.IntraFreqReselectionRedCap_r17 != nil {
			if err := e.EncodeEnumerated(*ie.IntraFreqReselectionRedCap_r17, sIB1V1700IEsIntraFreqReselectionRedCapR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. cellBarredNTN-r17: enumerated
	{
		if ie.CellBarredNTN_r17 != nil {
			if err := e.EncodeEnumerated(*ie.CellBarredNTN_r17, sIB1V1700IEsCellBarredNTNR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB1_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1V1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. hsdn-Cell-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sIB1V1700IEsHsdnCellR17Constraints)
			if err != nil {
				return err
			}
			ie.Hsdn_Cell_r17 = &idx
		}
	}

	// 3. uac-BarringInfo-v1700: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Uac_BarringInfo_v1700 = &struct{ Uac_BarringInfoSetList_v1700 UAC_BarringInfoSetList_v1700 }{}
			c := ie.Uac_BarringInfo_v1700
			{
				if err := c.Uac_BarringInfoSetList_v1700.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sdt-ConfigCommon-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sdt_ConfigCommon_r17 = new(SDT_ConfigCommonSIB_r17)
			if err := ie.Sdt_ConfigCommon_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. redCap-ConfigCommon-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.RedCap_ConfigCommon_r17 = new(RedCap_ConfigCommonSIB_r17)
			if err := ie.RedCap_ConfigCommon_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. featurePriorities-r17: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.FeaturePriorities_r17 = &struct {
				RedCapPriority_r17            *FeaturePriority_r17
				SlicingPriority_r17           *FeaturePriority_r17
				Msg3_Repetitions_Priority_r17 *FeaturePriority_r17
				Sdt_Priority_r17              *FeaturePriority_r17
			}{}
			c := ie.FeaturePriorities_r17
			sIB1V1700IEsFeaturePrioritiesR17Seq := d.NewSequenceDecoder(sIB1V1700IEsFeaturePrioritiesR17Constraints)
			if err := sIB1V1700IEsFeaturePrioritiesR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if sIB1V1700IEsFeaturePrioritiesR17Seq.IsComponentPresent(0) {
				c.RedCapPriority_r17 = new(FeaturePriority_r17)
				if err := (*c.RedCapPriority_r17).Decode(d); err != nil {
					return err
				}
			}
			if sIB1V1700IEsFeaturePrioritiesR17Seq.IsComponentPresent(1) {
				c.SlicingPriority_r17 = new(FeaturePriority_r17)
				if err := (*c.SlicingPriority_r17).Decode(d); err != nil {
					return err
				}
			}
			if sIB1V1700IEsFeaturePrioritiesR17Seq.IsComponentPresent(2) {
				c.Msg3_Repetitions_Priority_r17 = new(FeaturePriority_r17)
				if err := (*c.Msg3_Repetitions_Priority_r17).Decode(d); err != nil {
					return err
				}
			}
			if sIB1V1700IEsFeaturePrioritiesR17Seq.IsComponentPresent(3) {
				c.Sdt_Priority_r17 = new(FeaturePriority_r17)
				if err := (*c.Sdt_Priority_r17).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. si-SchedulingInfo-v1700: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Si_SchedulingInfo_v1700 = new(SI_SchedulingInfo_v1700)
			if err := ie.Si_SchedulingInfo_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. hyperSFN-r17: bit-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeBitString(sIB1V1700IEsHyperSFNR17Constraints)
			if err != nil {
				return err
			}
			ie.HyperSFN_r17 = &v
		}
	}

	// 9. eDRX-AllowedIdle-r17: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sIB1V1700IEsEDRXAllowedIdleR17Constraints)
			if err != nil {
				return err
			}
			ie.EDRX_AllowedIdle_r17 = &idx
		}
	}

	// 10. eDRX-AllowedInactive-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(sIB1V1700IEsEDRXAllowedInactiveR17Constraints)
			if err != nil {
				return err
			}
			ie.EDRX_AllowedInactive_r17 = &idx
		}
	}

	// 11. intraFreqReselectionRedCap-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sIB1V1700IEsIntraFreqReselectionRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.IntraFreqReselectionRedCap_r17 = &idx
		}
	}

	// 12. cellBarredNTN-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sIB1V1700IEsCellBarredNTNR17Constraints)
			if err != nil {
				return err
			}
			ie.CellBarredNTN_r17 = &idx
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(11) {
			ie.NonCriticalExtension = new(SIB1_v1740_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
