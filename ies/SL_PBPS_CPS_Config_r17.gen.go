// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PBPS-CPS-Config-r17 (line 27578).

var sLPBPSCPSConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-AllowedResourceSelectionConfig-r17", Optional: true},
		{Name: "sl-MinNumCandidateSlotsPeriodic-r17", Optional: true},
		{Name: "sl-PBPS-OccasionReservePeriodList-r17", Optional: true},
		{Name: "sl-Additional-PBPS-Occasion-r17", Optional: true},
		{Name: "sl-CPS-WindowPeriodic-r17", Optional: true},
		{Name: "sl-MinNumCandidateSlotsAperiodic-r17", Optional: true},
		{Name: "sl-MinNumRssiMeasurementSlots-r17", Optional: true},
		{Name: "sl-DefaultCBR-RandomSelection-r17", Optional: true},
		{Name: "sl-DefaultCBR-PartialSensing-r17", Optional: true},
		{Name: "sl-CPS-WindowAperiodic-r17", Optional: true},
		{Name: "sl-PartialSensingInactiveTime-r17", Optional: true},
	},
}

const (
	SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C1 = 0
	SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C2 = 1
	SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C3 = 2
	SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C4 = 3
	SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C5 = 4
	SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C6 = 5
	SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C7 = 6
)

var sLPBPSCPSConfigR17SlAllowedResourceSelectionConfigR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C1, SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C2, SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C3, SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C4, SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C5, SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C6, SL_PBPS_CPS_Config_r17_Sl_AllowedResourceSelectionConfig_r17_C7},
}

var sLPBPSCPSConfigR17SlMinNumCandidateSlotsPeriodicR17Constraints = per.Constrained(1, 32)

var sLPBPSCPSConfigR17SlPBPSOccasionReservePeriodListR17Constraints = per.SizeRange(1, 16)

const (
	SL_PBPS_CPS_Config_r17_Sl_Additional_PBPS_Occasion_r17_Monitored = 0
)

var sLPBPSCPSConfigR17SlAdditionalPBPSOccasionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PBPS_CPS_Config_r17_Sl_Additional_PBPS_Occasion_r17_Monitored},
}

var sLPBPSCPSConfigR17SlCPSWindowPeriodicR17Constraints = per.Constrained(5, 30)

var sLPBPSCPSConfigR17SlMinNumCandidateSlotsAperiodicR17Constraints = per.Constrained(1, 32)

var sLPBPSCPSConfigR17SlMinNumRssiMeasurementSlotsR17Constraints = per.Constrained(1, 800)

var sLPBPSCPSConfigR17SlDefaultCBRRandomSelectionR17Constraints = per.Constrained(0, 100)

var sLPBPSCPSConfigR17SlDefaultCBRPartialSensingR17Constraints = per.Constrained(0, 100)

var sLPBPSCPSConfigR17SlCPSWindowAperiodicR17Constraints = per.Constrained(0, 30)

const (
	SL_PBPS_CPS_Config_r17_Sl_PartialSensingInactiveTime_r17_Enabled  = 0
	SL_PBPS_CPS_Config_r17_Sl_PartialSensingInactiveTime_r17_Disabled = 1
)

var sLPBPSCPSConfigR17SlPartialSensingInactiveTimeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PBPS_CPS_Config_r17_Sl_PartialSensingInactiveTime_r17_Enabled, SL_PBPS_CPS_Config_r17_Sl_PartialSensingInactiveTime_r17_Disabled},
}

type SL_PBPS_CPS_Config_r17 struct {
	Sl_AllowedResourceSelectionConfig_r17 *int64
	Sl_MinNumCandidateSlotsPeriodic_r17   *int64
	Sl_PBPS_OccasionReservePeriodList_r17 []int64
	Sl_Additional_PBPS_Occasion_r17       *int64
	Sl_CPS_WindowPeriodic_r17             *int64
	Sl_MinNumCandidateSlotsAperiodic_r17  *int64
	Sl_MinNumRssiMeasurementSlots_r17     *int64
	Sl_DefaultCBR_RandomSelection_r17     *int64
	Sl_DefaultCBR_PartialSensing_r17      *int64
	Sl_CPS_WindowAperiodic_r17            *int64
	Sl_PartialSensingInactiveTime_r17     *int64
}

func (ie *SL_PBPS_CPS_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPBPSCPSConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_AllowedResourceSelectionConfig_r17 != nil, ie.Sl_MinNumCandidateSlotsPeriodic_r17 != nil, ie.Sl_PBPS_OccasionReservePeriodList_r17 != nil, ie.Sl_Additional_PBPS_Occasion_r17 != nil, ie.Sl_CPS_WindowPeriodic_r17 != nil, ie.Sl_MinNumCandidateSlotsAperiodic_r17 != nil, ie.Sl_MinNumRssiMeasurementSlots_r17 != nil, ie.Sl_DefaultCBR_RandomSelection_r17 != nil, ie.Sl_DefaultCBR_PartialSensing_r17 != nil, ie.Sl_CPS_WindowAperiodic_r17 != nil, ie.Sl_PartialSensingInactiveTime_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-AllowedResourceSelectionConfig-r17: enumerated
	{
		if ie.Sl_AllowedResourceSelectionConfig_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_AllowedResourceSelectionConfig_r17, sLPBPSCPSConfigR17SlAllowedResourceSelectionConfigR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-MinNumCandidateSlotsPeriodic-r17: integer
	{
		if ie.Sl_MinNumCandidateSlotsPeriodic_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_MinNumCandidateSlotsPeriodic_r17, sLPBPSCPSConfigR17SlMinNumCandidateSlotsPeriodicR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-PBPS-OccasionReservePeriodList-r17: sequence-of
	{
		if ie.Sl_PBPS_OccasionReservePeriodList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPBPSCPSConfigR17SlPBPSOccasionReservePeriodListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PBPS_OccasionReservePeriodList_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_PBPS_OccasionReservePeriodList_r17 {
				if err := e.EncodeInteger(ie.Sl_PBPS_OccasionReservePeriodList_r17[i], per.Constrained(1, 16)); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-Additional-PBPS-Occasion-r17: enumerated
	{
		if ie.Sl_Additional_PBPS_Occasion_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Additional_PBPS_Occasion_r17, sLPBPSCPSConfigR17SlAdditionalPBPSOccasionR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-CPS-WindowPeriodic-r17: integer
	{
		if ie.Sl_CPS_WindowPeriodic_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_CPS_WindowPeriodic_r17, sLPBPSCPSConfigR17SlCPSWindowPeriodicR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-MinNumCandidateSlotsAperiodic-r17: integer
	{
		if ie.Sl_MinNumCandidateSlotsAperiodic_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_MinNumCandidateSlotsAperiodic_r17, sLPBPSCPSConfigR17SlMinNumCandidateSlotsAperiodicR17Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-MinNumRssiMeasurementSlots-r17: integer
	{
		if ie.Sl_MinNumRssiMeasurementSlots_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_MinNumRssiMeasurementSlots_r17, sLPBPSCPSConfigR17SlMinNumRssiMeasurementSlotsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. sl-DefaultCBR-RandomSelection-r17: integer
	{
		if ie.Sl_DefaultCBR_RandomSelection_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_DefaultCBR_RandomSelection_r17, sLPBPSCPSConfigR17SlDefaultCBRRandomSelectionR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sl-DefaultCBR-PartialSensing-r17: integer
	{
		if ie.Sl_DefaultCBR_PartialSensing_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_DefaultCBR_PartialSensing_r17, sLPBPSCPSConfigR17SlDefaultCBRPartialSensingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. sl-CPS-WindowAperiodic-r17: integer
	{
		if ie.Sl_CPS_WindowAperiodic_r17 != nil {
			if err := e.EncodeInteger(*ie.Sl_CPS_WindowAperiodic_r17, sLPBPSCPSConfigR17SlCPSWindowAperiodicR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. sl-PartialSensingInactiveTime-r17: enumerated
	{
		if ie.Sl_PartialSensingInactiveTime_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PartialSensingInactiveTime_r17, sLPBPSCPSConfigR17SlPartialSensingInactiveTimeR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PBPS_CPS_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPBPSCPSConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-AllowedResourceSelectionConfig-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLPBPSCPSConfigR17SlAllowedResourceSelectionConfigR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_AllowedResourceSelectionConfig_r17 = &idx
		}
	}

	// 4. sl-MinNumCandidateSlotsPeriodic-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLPBPSCPSConfigR17SlMinNumCandidateSlotsPeriodicR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MinNumCandidateSlotsPeriodic_r17 = &v
		}
	}

	// 5. sl-PBPS-OccasionReservePeriodList-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLPBPSCPSConfigR17SlPBPSOccasionReservePeriodListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PBPS_OccasionReservePeriodList_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				ie.Sl_PBPS_OccasionReservePeriodList_r17[i] = v
			}
		}
	}

	// 6. sl-Additional-PBPS-Occasion-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLPBPSCPSConfigR17SlAdditionalPBPSOccasionR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Additional_PBPS_Occasion_r17 = &idx
		}
	}

	// 7. sl-CPS-WindowPeriodic-r17: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLPBPSCPSConfigR17SlCPSWindowPeriodicR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CPS_WindowPeriodic_r17 = &v
		}
	}

	// 8. sl-MinNumCandidateSlotsAperiodic-r17: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(sLPBPSCPSConfigR17SlMinNumCandidateSlotsAperiodicR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MinNumCandidateSlotsAperiodic_r17 = &v
		}
	}

	// 9. sl-MinNumRssiMeasurementSlots-r17: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(sLPBPSCPSConfigR17SlMinNumRssiMeasurementSlotsR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MinNumRssiMeasurementSlots_r17 = &v
		}
	}

	// 10. sl-DefaultCBR-RandomSelection-r17: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(sLPBPSCPSConfigR17SlDefaultCBRRandomSelectionR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DefaultCBR_RandomSelection_r17 = &v
		}
	}

	// 11. sl-DefaultCBR-PartialSensing-r17: integer
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeInteger(sLPBPSCPSConfigR17SlDefaultCBRPartialSensingR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DefaultCBR_PartialSensing_r17 = &v
		}
	}

	// 12. sl-CPS-WindowAperiodic-r17: integer
	{
		if seq.IsComponentPresent(9) {
			v, err := d.DecodeInteger(sLPBPSCPSConfigR17SlCPSWindowAperiodicR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CPS_WindowAperiodic_r17 = &v
		}
	}

	// 13. sl-PartialSensingInactiveTime-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sLPBPSCPSConfigR17SlPartialSensingInactiveTimeR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PartialSensingInactiveTime_r17 = &idx
		}
	}

	return nil
}
