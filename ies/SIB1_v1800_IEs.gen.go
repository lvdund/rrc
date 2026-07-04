// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB1-v1800-IEs (line 2050).

var sIB1V1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ncr-Support-r18", Optional: true},
		{Name: "mt-SDT-ConfigCommonSIB-r18", Optional: true},
		{Name: "musim-CapRestrictionAllowed-r18", Optional: true},
		{Name: "featurePriorities-v1800", Optional: true},
		{Name: "si-SchedulingInfo-v1800", Optional: true},
		{Name: "cellBarredATG-r18", Optional: true},
		{Name: "cellBarredNES-r18", Optional: true},
		{Name: "mobileIAB-Cell-r18", Optional: true},
		{Name: "eDRX-AllowedInactive-r18", Optional: true},
		{Name: "intraFreqReselection-eRedCap-r18", Optional: true},
		{Name: "nonServingCellMII-r18", Optional: true},
		{Name: "sdt-BeamFailureRecoveryProhibitTimer-r18", Optional: true},
		{Name: "eRedCap-ConfigCommon-r18", Optional: true},
		{Name: "cellBarredFixedVSAT-r18", Optional: true},
		{Name: "cellBarredMobileVSAT-r18", Optional: true},
		{Name: "reselectionMeasurementsNR-r18", Optional: true},
		{Name: "cellBarred2RxXR-r18", Optional: true},
		{Name: "intraFreqReselection2RxXR-r18", Optional: true},
		{Name: "barringExemptEmergencyCall-r18", Optional: true},
		{Name: "n3c-Support-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	SIB1_v1800_IEs_Ncr_Support_r18_True = 0
)

var sIB1V1800IEsNcrSupportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_Ncr_Support_r18_True},
}

const (
	SIB1_v1800_IEs_Musim_CapRestrictionAllowed_r18_True = 0
)

var sIB1V1800IEsMusimCapRestrictionAllowedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_Musim_CapRestrictionAllowed_r18_True},
}

const (
	SIB1_v1800_IEs_CellBarredATG_r18_Barred    = 0
	SIB1_v1800_IEs_CellBarredATG_r18_NotBarred = 1
)

var sIB1V1800IEsCellBarredATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_CellBarredATG_r18_Barred, SIB1_v1800_IEs_CellBarredATG_r18_NotBarred},
}

const (
	SIB1_v1800_IEs_CellBarredNES_r18_NotBarred = 0
)

var sIB1V1800IEsCellBarredNESR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_CellBarredNES_r18_NotBarred},
}

const (
	SIB1_v1800_IEs_MobileIAB_Cell_r18_True = 0
)

var sIB1V1800IEsMobileIABCellR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_MobileIAB_Cell_r18_True},
}

const (
	SIB1_v1800_IEs_EDRX_AllowedInactive_r18_True = 0
)

var sIB1V1800IEsEDRXAllowedInactiveR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_EDRX_AllowedInactive_r18_True},
}

const (
	SIB1_v1800_IEs_IntraFreqReselection_ERedCap_r18_Allowed    = 0
	SIB1_v1800_IEs_IntraFreqReselection_ERedCap_r18_NotAllowed = 1
)

var sIB1V1800IEsIntraFreqReselectionERedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_IntraFreqReselection_ERedCap_r18_Allowed, SIB1_v1800_IEs_IntraFreqReselection_ERedCap_r18_NotAllowed},
}

const (
	SIB1_v1800_IEs_NonServingCellMII_r18_True = 0
)

var sIB1V1800IEsNonServingCellMIIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_NonServingCellMII_r18_True},
}

const (
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms50   = 0
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms100  = 1
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms200  = 2
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms500  = 3
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms1000 = 4
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms1500 = 5
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms2000 = 6
	SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms3000 = 7
)

var sIB1V1800IEsSdtBeamFailureRecoveryProhibitTimerR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms50, SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms100, SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms200, SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms500, SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms1000, SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms1500, SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms2000, SIB1_v1800_IEs_Sdt_BeamFailureRecoveryProhibitTimer_r18_Ms3000},
}

const (
	SIB1_v1800_IEs_CellBarredFixedVSAT_r18_Barred    = 0
	SIB1_v1800_IEs_CellBarredFixedVSAT_r18_NotBarred = 1
)

var sIB1V1800IEsCellBarredFixedVSATR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_CellBarredFixedVSAT_r18_Barred, SIB1_v1800_IEs_CellBarredFixedVSAT_r18_NotBarred},
}

const (
	SIB1_v1800_IEs_CellBarredMobileVSAT_r18_Barred    = 0
	SIB1_v1800_IEs_CellBarredMobileVSAT_r18_NotBarred = 1
)

var sIB1V1800IEsCellBarredMobileVSATR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_CellBarredMobileVSAT_r18_Barred, SIB1_v1800_IEs_CellBarredMobileVSAT_r18_NotBarred},
}

const (
	SIB1_v1800_IEs_ReselectionMeasurementsNR_r18_True = 0
)

var sIB1V1800IEsReselectionMeasurementsNRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_ReselectionMeasurementsNR_r18_True},
}

const (
	SIB1_v1800_IEs_CellBarred2RxXR_r18_Barred = 0
)

var sIB1V1800IEsCellBarred2RxXRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_CellBarred2RxXR_r18_Barred},
}

const (
	SIB1_v1800_IEs_IntraFreqReselection2RxXR_r18_Allowed    = 0
	SIB1_v1800_IEs_IntraFreqReselection2RxXR_r18_NotAllowed = 1
)

var sIB1V1800IEsIntraFreqReselection2RxXRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_IntraFreqReselection2RxXR_r18_Allowed, SIB1_v1800_IEs_IntraFreqReselection2RxXR_r18_NotAllowed},
}

const (
	SIB1_v1800_IEs_BarringExemptEmergencyCall_r18_True = 0
)

var sIB1V1800IEsBarringExemptEmergencyCallR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_BarringExemptEmergencyCall_r18_True},
}

const (
	SIB1_v1800_IEs_N3c_Support_r18_True = 0
)

var sIB1V1800IEsN3cSupportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB1_v1800_IEs_N3c_Support_r18_True},
}

var sIB1V1800IEsFeaturePrioritiesV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "msg1-Repetitions-Priority-r18", Optional: true},
		{Name: "eRedCapPriority-r18", Optional: true},
	},
}

type SIB1_v1800_IEs struct {
	Ncr_Support_r18                 *int64
	Mt_SDT_ConfigCommonSIB_r18      *MT_SDT_ConfigCommonSIB_r18
	Musim_CapRestrictionAllowed_r18 *int64
	FeaturePriorities_v1800         *struct {
		Msg1_Repetitions_Priority_r18 *FeaturePriority_r17
		ERedCapPriority_r18           *FeaturePriority_r17
	}
	Si_SchedulingInfo_v1800                  *SI_SchedulingInfo_v1800
	CellBarredATG_r18                        *int64
	CellBarredNES_r18                        *int64
	MobileIAB_Cell_r18                       *int64
	EDRX_AllowedInactive_r18                 *int64
	IntraFreqReselection_ERedCap_r18         *int64
	NonServingCellMII_r18                    *int64
	Sdt_BeamFailureRecoveryProhibitTimer_r18 *int64
	ERedCap_ConfigCommon_r18                 *ERedCap_ConfigCommonSIB_r18
	CellBarredFixedVSAT_r18                  *int64
	CellBarredMobileVSAT_r18                 *int64
	ReselectionMeasurementsNR_r18            *int64
	CellBarred2RxXR_r18                      *int64
	IntraFreqReselection2RxXR_r18            *int64
	BarringExemptEmergencyCall_r18           *int64
	N3c_Support_r18                          *int64
	NonCriticalExtension                     *struct{}
}

func (ie *SIB1_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB1V1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ncr_Support_r18 != nil, ie.Mt_SDT_ConfigCommonSIB_r18 != nil, ie.Musim_CapRestrictionAllowed_r18 != nil, ie.FeaturePriorities_v1800 != nil, ie.Si_SchedulingInfo_v1800 != nil, ie.CellBarredATG_r18 != nil, ie.CellBarredNES_r18 != nil, ie.MobileIAB_Cell_r18 != nil, ie.EDRX_AllowedInactive_r18 != nil, ie.IntraFreqReselection_ERedCap_r18 != nil, ie.NonServingCellMII_r18 != nil, ie.Sdt_BeamFailureRecoveryProhibitTimer_r18 != nil, ie.ERedCap_ConfigCommon_r18 != nil, ie.CellBarredFixedVSAT_r18 != nil, ie.CellBarredMobileVSAT_r18 != nil, ie.ReselectionMeasurementsNR_r18 != nil, ie.CellBarred2RxXR_r18 != nil, ie.IntraFreqReselection2RxXR_r18 != nil, ie.BarringExemptEmergencyCall_r18 != nil, ie.N3c_Support_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. ncr-Support-r18: enumerated
	{
		if ie.Ncr_Support_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Ncr_Support_r18, sIB1V1800IEsNcrSupportR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. mt-SDT-ConfigCommonSIB-r18: ref
	{
		if ie.Mt_SDT_ConfigCommonSIB_r18 != nil {
			if err := ie.Mt_SDT_ConfigCommonSIB_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. musim-CapRestrictionAllowed-r18: enumerated
	{
		if ie.Musim_CapRestrictionAllowed_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_CapRestrictionAllowed_r18, sIB1V1800IEsMusimCapRestrictionAllowedR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. featurePriorities-v1800: sequence
	{
		if ie.FeaturePriorities_v1800 != nil {
			c := ie.FeaturePriorities_v1800
			sIB1V1800IEsFeaturePrioritiesV1800Seq := e.NewSequenceEncoder(sIB1V1800IEsFeaturePrioritiesV1800Constraints)
			if err := sIB1V1800IEsFeaturePrioritiesV1800Seq.EncodePreamble([]bool{c.Msg1_Repetitions_Priority_r18 != nil, c.ERedCapPriority_r18 != nil}); err != nil {
				return err
			}
			if c.Msg1_Repetitions_Priority_r18 != nil {
				if err := c.Msg1_Repetitions_Priority_r18.Encode(e); err != nil {
					return err
				}
			}
			if c.ERedCapPriority_r18 != nil {
				if err := c.ERedCapPriority_r18.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. si-SchedulingInfo-v1800: ref
	{
		if ie.Si_SchedulingInfo_v1800 != nil {
			if err := ie.Si_SchedulingInfo_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. cellBarredATG-r18: enumerated
	{
		if ie.CellBarredATG_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CellBarredATG_r18, sIB1V1800IEsCellBarredATGR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. cellBarredNES-r18: enumerated
	{
		if ie.CellBarredNES_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CellBarredNES_r18, sIB1V1800IEsCellBarredNESR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. mobileIAB-Cell-r18: enumerated
	{
		if ie.MobileIAB_Cell_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MobileIAB_Cell_r18, sIB1V1800IEsMobileIABCellR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. eDRX-AllowedInactive-r18: enumerated
	{
		if ie.EDRX_AllowedInactive_r18 != nil {
			if err := e.EncodeEnumerated(*ie.EDRX_AllowedInactive_r18, sIB1V1800IEsEDRXAllowedInactiveR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. intraFreqReselection-eRedCap-r18: enumerated
	{
		if ie.IntraFreqReselection_ERedCap_r18 != nil {
			if err := e.EncodeEnumerated(*ie.IntraFreqReselection_ERedCap_r18, sIB1V1800IEsIntraFreqReselectionERedCapR18Constraints); err != nil {
				return err
			}
		}
	}

	// 12. nonServingCellMII-r18: enumerated
	{
		if ie.NonServingCellMII_r18 != nil {
			if err := e.EncodeEnumerated(*ie.NonServingCellMII_r18, sIB1V1800IEsNonServingCellMIIR18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. sdt-BeamFailureRecoveryProhibitTimer-r18: enumerated
	{
		if ie.Sdt_BeamFailureRecoveryProhibitTimer_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sdt_BeamFailureRecoveryProhibitTimer_r18, sIB1V1800IEsSdtBeamFailureRecoveryProhibitTimerR18Constraints); err != nil {
				return err
			}
		}
	}

	// 14. eRedCap-ConfigCommon-r18: ref
	{
		if ie.ERedCap_ConfigCommon_r18 != nil {
			if err := ie.ERedCap_ConfigCommon_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. cellBarredFixedVSAT-r18: enumerated
	{
		if ie.CellBarredFixedVSAT_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CellBarredFixedVSAT_r18, sIB1V1800IEsCellBarredFixedVSATR18Constraints); err != nil {
				return err
			}
		}
	}

	// 16. cellBarredMobileVSAT-r18: enumerated
	{
		if ie.CellBarredMobileVSAT_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CellBarredMobileVSAT_r18, sIB1V1800IEsCellBarredMobileVSATR18Constraints); err != nil {
				return err
			}
		}
	}

	// 17. reselectionMeasurementsNR-r18: enumerated
	{
		if ie.ReselectionMeasurementsNR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ReselectionMeasurementsNR_r18, sIB1V1800IEsReselectionMeasurementsNRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 18. cellBarred2RxXR-r18: enumerated
	{
		if ie.CellBarred2RxXR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CellBarred2RxXR_r18, sIB1V1800IEsCellBarred2RxXRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 19. intraFreqReselection2RxXR-r18: enumerated
	{
		if ie.IntraFreqReselection2RxXR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.IntraFreqReselection2RxXR_r18, sIB1V1800IEsIntraFreqReselection2RxXRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 20. barringExemptEmergencyCall-r18: enumerated
	{
		if ie.BarringExemptEmergencyCall_r18 != nil {
			if err := e.EncodeEnumerated(*ie.BarringExemptEmergencyCall_r18, sIB1V1800IEsBarringExemptEmergencyCallR18Constraints); err != nil {
				return err
			}
		}
	}

	// 21. n3c-Support-r18: enumerated
	{
		if ie.N3c_Support_r18 != nil {
			if err := e.EncodeEnumerated(*ie.N3c_Support_r18, sIB1V1800IEsN3cSupportR18Constraints); err != nil {
				return err
			}
		}
	}

	// 22. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *SIB1_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB1V1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ncr-Support-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsNcrSupportR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_Support_r18 = &idx
		}
	}

	// 3. mt-SDT-ConfigCommonSIB-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Mt_SDT_ConfigCommonSIB_r18 = new(MT_SDT_ConfigCommonSIB_r18)
			if err := ie.Mt_SDT_ConfigCommonSIB_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. musim-CapRestrictionAllowed-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsMusimCapRestrictionAllowedR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_CapRestrictionAllowed_r18 = &idx
		}
	}

	// 5. featurePriorities-v1800: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.FeaturePriorities_v1800 = &struct {
				Msg1_Repetitions_Priority_r18 *FeaturePriority_r17
				ERedCapPriority_r18           *FeaturePriority_r17
			}{}
			c := ie.FeaturePriorities_v1800
			sIB1V1800IEsFeaturePrioritiesV1800Seq := d.NewSequenceDecoder(sIB1V1800IEsFeaturePrioritiesV1800Constraints)
			if err := sIB1V1800IEsFeaturePrioritiesV1800Seq.DecodePreamble(); err != nil {
				return err
			}
			if sIB1V1800IEsFeaturePrioritiesV1800Seq.IsComponentPresent(0) {
				c.Msg1_Repetitions_Priority_r18 = new(FeaturePriority_r17)
				if err := (*c.Msg1_Repetitions_Priority_r18).Decode(d); err != nil {
					return err
				}
			}
			if sIB1V1800IEsFeaturePrioritiesV1800Seq.IsComponentPresent(1) {
				c.ERedCapPriority_r18 = new(FeaturePriority_r17)
				if err := (*c.ERedCapPriority_r18).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. si-SchedulingInfo-v1800: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Si_SchedulingInfo_v1800 = new(SI_SchedulingInfo_v1800)
			if err := ie.Si_SchedulingInfo_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. cellBarredATG-r18: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsCellBarredATGR18Constraints)
			if err != nil {
				return err
			}
			ie.CellBarredATG_r18 = &idx
		}
	}

	// 8. cellBarredNES-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsCellBarredNESR18Constraints)
			if err != nil {
				return err
			}
			ie.CellBarredNES_r18 = &idx
		}
	}

	// 9. mobileIAB-Cell-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsMobileIABCellR18Constraints)
			if err != nil {
				return err
			}
			ie.MobileIAB_Cell_r18 = &idx
		}
	}

	// 10. eDRX-AllowedInactive-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsEDRXAllowedInactiveR18Constraints)
			if err != nil {
				return err
			}
			ie.EDRX_AllowedInactive_r18 = &idx
		}
	}

	// 11. intraFreqReselection-eRedCap-r18: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsIntraFreqReselectionERedCapR18Constraints)
			if err != nil {
				return err
			}
			ie.IntraFreqReselection_ERedCap_r18 = &idx
		}
	}

	// 12. nonServingCellMII-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsNonServingCellMIIR18Constraints)
			if err != nil {
				return err
			}
			ie.NonServingCellMII_r18 = &idx
		}
	}

	// 13. sdt-BeamFailureRecoveryProhibitTimer-r18: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsSdtBeamFailureRecoveryProhibitTimerR18Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_BeamFailureRecoveryProhibitTimer_r18 = &idx
		}
	}

	// 14. eRedCap-ConfigCommon-r18: ref
	{
		if seq.IsComponentPresent(12) {
			ie.ERedCap_ConfigCommon_r18 = new(ERedCap_ConfigCommonSIB_r18)
			if err := ie.ERedCap_ConfigCommon_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. cellBarredFixedVSAT-r18: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsCellBarredFixedVSATR18Constraints)
			if err != nil {
				return err
			}
			ie.CellBarredFixedVSAT_r18 = &idx
		}
	}

	// 16. cellBarredMobileVSAT-r18: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsCellBarredMobileVSATR18Constraints)
			if err != nil {
				return err
			}
			ie.CellBarredMobileVSAT_r18 = &idx
		}
	}

	// 17. reselectionMeasurementsNR-r18: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsReselectionMeasurementsNRR18Constraints)
			if err != nil {
				return err
			}
			ie.ReselectionMeasurementsNR_r18 = &idx
		}
	}

	// 18. cellBarred2RxXR-r18: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsCellBarred2RxXRR18Constraints)
			if err != nil {
				return err
			}
			ie.CellBarred2RxXR_r18 = &idx
		}
	}

	// 19. intraFreqReselection2RxXR-r18: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsIntraFreqReselection2RxXRR18Constraints)
			if err != nil {
				return err
			}
			ie.IntraFreqReselection2RxXR_r18 = &idx
		}
	}

	// 20. barringExemptEmergencyCall-r18: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsBarringExemptEmergencyCallR18Constraints)
			if err != nil {
				return err
			}
			ie.BarringExemptEmergencyCall_r18 = &idx
		}
	}

	// 21. n3c-Support-r18: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(sIB1V1800IEsN3cSupportR18Constraints)
			if err != nil {
				return err
			}
			ie.N3c_Support_r18 = &idx
		}
	}

	// 22. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(20) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
