// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SPS-Config (line 15232).

var sPSConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicity"},
		{Name: "nrofHARQ-Processes"},
		{Name: "n1PUCCH-AN", Optional: true},
		{Name: "mcs-Table", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	SPS_Config_Periodicity_Ms10   = 0
	SPS_Config_Periodicity_Ms20   = 1
	SPS_Config_Periodicity_Ms32   = 2
	SPS_Config_Periodicity_Ms40   = 3
	SPS_Config_Periodicity_Ms64   = 4
	SPS_Config_Periodicity_Ms80   = 5
	SPS_Config_Periodicity_Ms128  = 6
	SPS_Config_Periodicity_Ms160  = 7
	SPS_Config_Periodicity_Ms320  = 8
	SPS_Config_Periodicity_Ms640  = 9
	SPS_Config_Periodicity_Spare6 = 10
	SPS_Config_Periodicity_Spare5 = 11
	SPS_Config_Periodicity_Spare4 = 12
	SPS_Config_Periodicity_Spare3 = 13
	SPS_Config_Periodicity_Spare2 = 14
	SPS_Config_Periodicity_Spare1 = 15
)

var sPSConfigPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SPS_Config_Periodicity_Ms10, SPS_Config_Periodicity_Ms20, SPS_Config_Periodicity_Ms32, SPS_Config_Periodicity_Ms40, SPS_Config_Periodicity_Ms64, SPS_Config_Periodicity_Ms80, SPS_Config_Periodicity_Ms128, SPS_Config_Periodicity_Ms160, SPS_Config_Periodicity_Ms320, SPS_Config_Periodicity_Ms640, SPS_Config_Periodicity_Spare6, SPS_Config_Periodicity_Spare5, SPS_Config_Periodicity_Spare4, SPS_Config_Periodicity_Spare3, SPS_Config_Periodicity_Spare2, SPS_Config_Periodicity_Spare1},
}

var sPSConfigNrofHARQProcessesConstraints = per.Constrained(1, 8)

const (
	SPS_Config_Mcs_Table_Qam64LowSE = 0
)

var sPSConfigMcsTableConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SPS_Config_Mcs_Table_Qam64LowSE},
}

var sPSConfigHarqProcIDOffsetR16Constraints = per.Constrained(0, 15)

var sPSConfigPeriodicityExtR16Constraints = per.Constrained(1, 5120)

var sPSConfigHarqCodebookIDR16Constraints = per.Constrained(1, 2)

const (
	SPS_Config_Ext_Pdsch_AggregationFactor_r16_N1 = 0
	SPS_Config_Ext_Pdsch_AggregationFactor_r16_N2 = 1
	SPS_Config_Ext_Pdsch_AggregationFactor_r16_N4 = 2
	SPS_Config_Ext_Pdsch_AggregationFactor_r16_N8 = 3
)

var sPSConfigExtPdschAggregationFactorR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SPS_Config_Ext_Pdsch_AggregationFactor_r16_N1, SPS_Config_Ext_Pdsch_AggregationFactor_r16_N2, SPS_Config_Ext_Pdsch_AggregationFactor_r16_N4, SPS_Config_Ext_Pdsch_AggregationFactor_r16_N8},
}

var sPSConfigSpsHARQDeferralR17Constraints = per.Constrained(1, 32)

var sPSConfigPeriodicityExtR17Constraints = per.Constrained(1, 40960)

var sPSConfigNrofHARQProcessesV1710Constraints = per.Constrained(9, 32)

var sPSConfigHarqProcIDOffsetV1700Constraints = per.Constrained(16, 31)

type SPS_Config struct {
	Periodicity                 int64
	NrofHARQ_Processes          int64
	N1PUCCH_AN                  *PUCCH_ResourceId
	Mcs_Table                   *int64
	Sps_ConfigIndex_r16         *SPS_ConfigIndex_r16
	Harq_ProcID_Offset_r16      *int64
	PeriodicityExt_r16          *int64
	Harq_CodebookID_r16         *int64
	Pdsch_AggregationFactor_r16 *int64
	Sps_HARQ_Deferral_r17       *int64
	N1PUCCH_AN_PUCCHsSCell_r17  *PUCCH_ResourceId
	PeriodicityExt_r17          *int64
	NrofHARQ_Processes_v1710    *int64
	Harq_ProcID_Offset_v1700    *int64
}

func (ie *SPS_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sPSConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sps_ConfigIndex_r16 != nil || ie.Harq_ProcID_Offset_r16 != nil || ie.PeriodicityExt_r16 != nil || ie.Harq_CodebookID_r16 != nil || ie.Pdsch_AggregationFactor_r16 != nil
	hasExtGroup1 := ie.Sps_HARQ_Deferral_r17 != nil || ie.N1PUCCH_AN_PUCCHsSCell_r17 != nil || ie.PeriodicityExt_r17 != nil || ie.NrofHARQ_Processes_v1710 != nil || ie.Harq_ProcID_Offset_v1700 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.N1PUCCH_AN != nil, ie.Mcs_Table != nil}); err != nil {
		return err
	}

	// 3. periodicity: enumerated
	{
		if err := e.EncodeEnumerated(ie.Periodicity, sPSConfigPeriodicityConstraints); err != nil {
			return err
		}
	}

	// 4. nrofHARQ-Processes: integer
	{
		if err := e.EncodeInteger(ie.NrofHARQ_Processes, sPSConfigNrofHARQProcessesConstraints); err != nil {
			return err
		}
	}

	// 5. n1PUCCH-AN: ref
	{
		if ie.N1PUCCH_AN != nil {
			if err := ie.N1PUCCH_AN.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mcs-Table: enumerated
	{
		if ie.Mcs_Table != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_Table, sPSConfigMcsTableConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sps-ConfigIndex-r16", Optional: true},
					{Name: "harq-ProcID-Offset-r16", Optional: true},
					{Name: "periodicityExt-r16", Optional: true},
					{Name: "harq-CodebookID-r16", Optional: true},
					{Name: "pdsch-AggregationFactor-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sps_ConfigIndex_r16 != nil, ie.Harq_ProcID_Offset_r16 != nil, ie.PeriodicityExt_r16 != nil, ie.Harq_CodebookID_r16 != nil, ie.Pdsch_AggregationFactor_r16 != nil}); err != nil {
				return err
			}

			if ie.Sps_ConfigIndex_r16 != nil {
				if err := ie.Sps_ConfigIndex_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Harq_ProcID_Offset_r16 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcID_Offset_r16, sPSConfigHarqProcIDOffsetR16Constraints); err != nil {
					return err
				}
			}

			if ie.PeriodicityExt_r16 != nil {
				if err := ex.EncodeInteger(*ie.PeriodicityExt_r16, sPSConfigPeriodicityExtR16Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_CodebookID_r16 != nil {
				if err := ex.EncodeInteger(*ie.Harq_CodebookID_r16, sPSConfigHarqCodebookIDR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_AggregationFactor_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_AggregationFactor_r16, sPSConfigExtPdschAggregationFactorR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sps-HARQ-Deferral-r17", Optional: true},
					{Name: "n1PUCCH-AN-PUCCHsSCell-r17", Optional: true},
					{Name: "periodicityExt-r17", Optional: true},
					{Name: "nrofHARQ-Processes-v1710", Optional: true},
					{Name: "harq-ProcID-Offset-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sps_HARQ_Deferral_r17 != nil, ie.N1PUCCH_AN_PUCCHsSCell_r17 != nil, ie.PeriodicityExt_r17 != nil, ie.NrofHARQ_Processes_v1710 != nil, ie.Harq_ProcID_Offset_v1700 != nil}); err != nil {
				return err
			}

			if ie.Sps_HARQ_Deferral_r17 != nil {
				if err := ex.EncodeInteger(*ie.Sps_HARQ_Deferral_r17, sPSConfigSpsHARQDeferralR17Constraints); err != nil {
					return err
				}
			}

			if ie.N1PUCCH_AN_PUCCHsSCell_r17 != nil {
				if err := ie.N1PUCCH_AN_PUCCHsSCell_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PeriodicityExt_r17 != nil {
				if err := ex.EncodeInteger(*ie.PeriodicityExt_r17, sPSConfigPeriodicityExtR17Constraints); err != nil {
					return err
				}
			}

			if ie.NrofHARQ_Processes_v1710 != nil {
				if err := ex.EncodeInteger(*ie.NrofHARQ_Processes_v1710, sPSConfigNrofHARQProcessesV1710Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcID_Offset_v1700 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcID_Offset_v1700, sPSConfigHarqProcIDOffsetV1700Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SPS_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sPSConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. periodicity: enumerated
	{
		v0, err := d.DecodeEnumerated(sPSConfigPeriodicityConstraints)
		if err != nil {
			return err
		}
		ie.Periodicity = v0
	}

	// 4. nrofHARQ-Processes: integer
	{
		v1, err := d.DecodeInteger(sPSConfigNrofHARQProcessesConstraints)
		if err != nil {
			return err
		}
		ie.NrofHARQ_Processes = v1
	}

	// 5. n1PUCCH-AN: ref
	{
		if seq.IsComponentPresent(2) {
			ie.N1PUCCH_AN = new(PUCCH_ResourceId)
			if err := ie.N1PUCCH_AN.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mcs-Table: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sPSConfigMcsTableConstraints)
			if err != nil {
				return err
			}
			ie.Mcs_Table = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sps-ConfigIndex-r16", Optional: true},
				{Name: "harq-ProcID-Offset-r16", Optional: true},
				{Name: "periodicityExt-r16", Optional: true},
				{Name: "harq-CodebookID-r16", Optional: true},
				{Name: "pdsch-AggregationFactor-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sps_ConfigIndex_r16 = new(SPS_ConfigIndex_r16)
			if err := ie.Sps_ConfigIndex_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(sPSConfigHarqProcIDOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcID_Offset_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(sPSConfigPeriodicityExtR16Constraints)
			if err != nil {
				return err
			}
			ie.PeriodicityExt_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(sPSConfigHarqCodebookIDR16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_CodebookID_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(sPSConfigExtPdschAggregationFactorR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_AggregationFactor_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sps-HARQ-Deferral-r17", Optional: true},
				{Name: "n1PUCCH-AN-PUCCHsSCell-r17", Optional: true},
				{Name: "periodicityExt-r17", Optional: true},
				{Name: "nrofHARQ-Processes-v1710", Optional: true},
				{Name: "harq-ProcID-Offset-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sPSConfigSpsHARQDeferralR17Constraints)
			if err != nil {
				return err
			}
			ie.Sps_HARQ_Deferral_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.N1PUCCH_AN_PUCCHsSCell_r17 = new(PUCCH_ResourceId)
			if err := ie.N1PUCCH_AN_PUCCHsSCell_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(sPSConfigPeriodicityExtR17Constraints)
			if err != nil {
				return err
			}
			ie.PeriodicityExt_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(sPSConfigNrofHARQProcessesV1710Constraints)
			if err != nil {
				return err
			}
			ie.NrofHARQ_Processes_v1710 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeInteger(sPSConfigHarqProcIDOffsetV1700Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcID_Offset_v1700 = &v
		}
	}

	return nil
}
