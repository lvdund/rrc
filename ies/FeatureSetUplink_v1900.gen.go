// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1900 (line 20418).

var featureSetUplinkV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-PortGrouping-r19", Optional: true},
		{Name: "nonCodebook-CSI-RS-SRS-Enh-r19", Optional: true},
		{Name: "srs-AntennaSwitching3T6R2SP-1Periodic-r19", Optional: true},
		{Name: "srs-AntennaSwitching3T3R2SP-1Periodic-r19", Optional: true},
		{Name: "mTRP-PUSCH-TypeA-CB-3Port-r19", Optional: true},
		{Name: "mTRP-PUSCH-RepetitionTypeA-3Port-r19", Optional: true},
		{Name: "threePortsPTRS-PUSCH-r19", Optional: true},
		{Name: "ul-FullPwrMode-3Port-r19", Optional: true},
		{Name: "simultaneous-2-1-HARQ-ACK-CB-Diff-r19", Optional: true},
		{Name: "simultaneous-2-2-HARQ-ACK-CB-Diff-r19", Optional: true},
		{Name: "ul-IntraUE-MuxEnh-Diff-r19", Optional: true},
	},
}

const (
	FeatureSetUplink_v1900_Srs_PortGrouping_r19_Xt8r = 0
	FeatureSetUplink_v1900_Srs_PortGrouping_r19_Xt6r = 1
	FeatureSetUplink_v1900_Srs_PortGrouping_r19_Both = 2
)

var featureSetUplinkV1900SrsPortGroupingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1900_Srs_PortGrouping_r19_Xt8r, FeatureSetUplink_v1900_Srs_PortGrouping_r19_Xt6r, FeatureSetUplink_v1900_Srs_PortGrouping_r19_Both},
}

const (
	FeatureSetUplink_v1900_NonCodebook_CSI_RS_SRS_Enh_r19_Supported = 0
)

var featureSetUplinkV1900NonCodebookCSIRSSRSEnhR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1900_NonCodebook_CSI_RS_SRS_Enh_r19_Supported},
}

const (
	FeatureSetUplink_v1900_Srs_AntennaSwitching3T6R2SP_1Periodic_r19_Supported = 0
)

var featureSetUplinkV1900SrsAntennaSwitching3T6R2SP1PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1900_Srs_AntennaSwitching3T6R2SP_1Periodic_r19_Supported},
}

const (
	FeatureSetUplink_v1900_Srs_AntennaSwitching3T3R2SP_1Periodic_r19_Supported = 0
)

var featureSetUplinkV1900SrsAntennaSwitching3T3R2SP1PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1900_Srs_AntennaSwitching3T3R2SP_1Periodic_r19_Supported},
}

var featureSetUplinkV1900MTRPPUSCHTypeACB3PortR19Constraints = per.Constrained(1, 2)

var featureSetUplinkV1900MTRPPUSCHRepetitionTypeA3PortR19Constraints = per.Constrained(1, 3)

var featureSetUplinkV1900ThreePortsPTRSPUSCHR19Constraints = per.Constrained(1, 2)

const (
	FeatureSetUplink_v1900_Ul_FullPwrMode_3Port_r19_Supported = 0
)

var featureSetUplinkV1900UlFullPwrMode3PortR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1900_Ul_FullPwrMode_3Port_r19_Supported},
}

const (
	FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationLowPriority_r19_Sym0 = 0
	FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationLowPriority_r19_Sym1 = 1
	FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationLowPriority_r19_Sym2 = 2
)

var featureSetUplinkV1900UlIntraUEMuxEnhDiffR19PuschPreparationLowPriorityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationLowPriority_r19_Sym0, FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationLowPriority_r19_Sym1, FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationLowPriority_r19_Sym2},
}

const (
	FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationHighPriority_r19_Sym0 = 0
	FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationHighPriority_r19_Sym1 = 1
	FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationHighPriority_r19_Sym2 = 2
)

var featureSetUplinkV1900UlIntraUEMuxEnhDiffR19PuschPreparationHighPriorityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationHighPriority_r19_Sym0, FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationHighPriority_r19_Sym1, FeatureSetUplink_v1900_Ul_IntraUE_MuxEnh_Diff_r19_Pusch_PreparationHighPriority_r19_Sym2},
}

type FeatureSetUplink_v1900 struct {
	Srs_PortGrouping_r19                      *int64
	NonCodebook_CSI_RS_SRS_Enh_r19            *int64
	Srs_AntennaSwitching3T6R2SP_1Periodic_r19 *int64
	Srs_AntennaSwitching3T3R2SP_1Periodic_r19 *int64
	MTRP_PUSCH_TypeA_CB_3Port_r19             *int64
	MTRP_PUSCH_RepetitionTypeA_3Port_r19      *int64
	ThreePortsPTRS_PUSCH_r19                  *int64
	Ul_FullPwrMode_3Port_r19                  *int64
	Simultaneous_2_1_HARQ_ACK_CB_Diff_r19     *SubSlot_Config_r16
	Simultaneous_2_2_HARQ_ACK_CB_Diff_r19     *SubSlot_Config_r16
	Ul_IntraUE_MuxEnh_Diff_r19                *struct {
		Pusch_PreparationLowPriority_r19  int64
		Pusch_PreparationHighPriority_r19 int64
	}
}

func (ie *FeatureSetUplink_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_PortGrouping_r19 != nil, ie.NonCodebook_CSI_RS_SRS_Enh_r19 != nil, ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19 != nil, ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19 != nil, ie.MTRP_PUSCH_TypeA_CB_3Port_r19 != nil, ie.MTRP_PUSCH_RepetitionTypeA_3Port_r19 != nil, ie.ThreePortsPTRS_PUSCH_r19 != nil, ie.Ul_FullPwrMode_3Port_r19 != nil, ie.Simultaneous_2_1_HARQ_ACK_CB_Diff_r19 != nil, ie.Simultaneous_2_2_HARQ_ACK_CB_Diff_r19 != nil, ie.Ul_IntraUE_MuxEnh_Diff_r19 != nil}); err != nil {
		return err
	}

	// 2. srs-PortGrouping-r19: enumerated
	{
		if ie.Srs_PortGrouping_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_PortGrouping_r19, featureSetUplinkV1900SrsPortGroupingR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. nonCodebook-CSI-RS-SRS-Enh-r19: enumerated
	{
		if ie.NonCodebook_CSI_RS_SRS_Enh_r19 != nil {
			if err := e.EncodeEnumerated(*ie.NonCodebook_CSI_RS_SRS_Enh_r19, featureSetUplinkV1900NonCodebookCSIRSSRSEnhR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. srs-AntennaSwitching3T6R2SP-1Periodic-r19: enumerated
	{
		if ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19, featureSetUplinkV1900SrsAntennaSwitching3T6R2SP1PeriodicR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. srs-AntennaSwitching3T3R2SP-1Periodic-r19: enumerated
	{
		if ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19, featureSetUplinkV1900SrsAntennaSwitching3T3R2SP1PeriodicR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. mTRP-PUSCH-TypeA-CB-3Port-r19: integer
	{
		if ie.MTRP_PUSCH_TypeA_CB_3Port_r19 != nil {
			if err := e.EncodeInteger(*ie.MTRP_PUSCH_TypeA_CB_3Port_r19, featureSetUplinkV1900MTRPPUSCHTypeACB3PortR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. mTRP-PUSCH-RepetitionTypeA-3Port-r19: integer
	{
		if ie.MTRP_PUSCH_RepetitionTypeA_3Port_r19 != nil {
			if err := e.EncodeInteger(*ie.MTRP_PUSCH_RepetitionTypeA_3Port_r19, featureSetUplinkV1900MTRPPUSCHRepetitionTypeA3PortR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. threePortsPTRS-PUSCH-r19: integer
	{
		if ie.ThreePortsPTRS_PUSCH_r19 != nil {
			if err := e.EncodeInteger(*ie.ThreePortsPTRS_PUSCH_r19, featureSetUplinkV1900ThreePortsPTRSPUSCHR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. ul-FullPwrMode-3Port-r19: enumerated
	{
		if ie.Ul_FullPwrMode_3Port_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_FullPwrMode_3Port_r19, featureSetUplinkV1900UlFullPwrMode3PortR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. simultaneous-2-1-HARQ-ACK-CB-Diff-r19: ref
	{
		if ie.Simultaneous_2_1_HARQ_ACK_CB_Diff_r19 != nil {
			if err := ie.Simultaneous_2_1_HARQ_ACK_CB_Diff_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. simultaneous-2-2-HARQ-ACK-CB-Diff-r19: ref
	{
		if ie.Simultaneous_2_2_HARQ_ACK_CB_Diff_r19 != nil {
			if err := ie.Simultaneous_2_2_HARQ_ACK_CB_Diff_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. ul-IntraUE-MuxEnh-Diff-r19: sequence
	{
		if ie.Ul_IntraUE_MuxEnh_Diff_r19 != nil {
			c := ie.Ul_IntraUE_MuxEnh_Diff_r19
			if err := e.EncodeEnumerated(c.Pusch_PreparationLowPriority_r19, featureSetUplinkV1900UlIntraUEMuxEnhDiffR19PuschPreparationLowPriorityR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Pusch_PreparationHighPriority_r19, featureSetUplinkV1900UlIntraUEMuxEnhDiffR19PuschPreparationHighPriorityR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-PortGrouping-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1900SrsPortGroupingR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_PortGrouping_r19 = &idx
		}
	}

	// 3. nonCodebook-CSI-RS-SRS-Enh-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1900NonCodebookCSIRSSRSEnhR19Constraints)
			if err != nil {
				return err
			}
			ie.NonCodebook_CSI_RS_SRS_Enh_r19 = &idx
		}
	}

	// 4. srs-AntennaSwitching3T6R2SP-1Periodic-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1900SrsAntennaSwitching3T6R2SP1PeriodicR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19 = &idx
		}
	}

	// 5. srs-AntennaSwitching3T3R2SP-1Periodic-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1900SrsAntennaSwitching3T3R2SP1PeriodicR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19 = &idx
		}
	}

	// 6. mTRP-PUSCH-TypeA-CB-3Port-r19: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(featureSetUplinkV1900MTRPPUSCHTypeACB3PortR19Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_TypeA_CB_3Port_r19 = &v
		}
	}

	// 7. mTRP-PUSCH-RepetitionTypeA-3Port-r19: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(featureSetUplinkV1900MTRPPUSCHRepetitionTypeA3PortR19Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_RepetitionTypeA_3Port_r19 = &v
		}
	}

	// 8. threePortsPTRS-PUSCH-r19: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(featureSetUplinkV1900ThreePortsPTRSPUSCHR19Constraints)
			if err != nil {
				return err
			}
			ie.ThreePortsPTRS_PUSCH_r19 = &v
		}
	}

	// 9. ul-FullPwrMode-3Port-r19: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1900UlFullPwrMode3PortR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FullPwrMode_3Port_r19 = &idx
		}
	}

	// 10. simultaneous-2-1-HARQ-ACK-CB-Diff-r19: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Simultaneous_2_1_HARQ_ACK_CB_Diff_r19 = new(SubSlot_Config_r16)
			if err := ie.Simultaneous_2_1_HARQ_ACK_CB_Diff_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. simultaneous-2-2-HARQ-ACK-CB-Diff-r19: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Simultaneous_2_2_HARQ_ACK_CB_Diff_r19 = new(SubSlot_Config_r16)
			if err := ie.Simultaneous_2_2_HARQ_ACK_CB_Diff_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. ul-IntraUE-MuxEnh-Diff-r19: sequence
	{
		if seq.IsComponentPresent(10) {
			ie.Ul_IntraUE_MuxEnh_Diff_r19 = &struct {
				Pusch_PreparationLowPriority_r19  int64
				Pusch_PreparationHighPriority_r19 int64
			}{}
			c := ie.Ul_IntraUE_MuxEnh_Diff_r19
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1900UlIntraUEMuxEnhDiffR19PuschPreparationLowPriorityR19Constraints)
				if err != nil {
					return err
				}
				c.Pusch_PreparationLowPriority_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1900UlIntraUEMuxEnhDiffR19PuschPreparationHighPriorityR19Constraints)
				if err != nil {
					return err
				}
				c.Pusch_PreparationHighPriority_r19 = v
			}
		}
	}

	return nil
}
