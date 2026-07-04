// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: OD-SIB1-Config-r19 (line 4693).

var oDSIB1ConfigR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierFreq-r19"},
		{Name: "validSIB26-Cells-r19"},
		{Name: "sib1-TDD-UL-DL-ConfigurationCommon-r19", Optional: true},
		{Name: "offsetToPointA-r19", Optional: true},
		{Name: "frequencyBandList-r19", Optional: true},
		{Name: "ss-PBCH-BlockPower-r19"},
		{Name: "ssb-PositionsInBurst-r19"},
		{Name: "ssb-Periodicity-r19", Optional: true},
		{Name: "od-SIB1-WindowDuration-r19"},
		{Name: "od-SIB1-WindowStartOffset-r19"},
		{Name: "k-SSB-r19"},
		{Name: "controlResourceSetZero-r19"},
		{Name: "searchSpaceZero-r19"},
		{Name: "sib1-RequestConfig-r19"},
		{Name: "sib1-RequestConfigSUL-r19", Optional: true},
	},
}

var oDSIB1ConfigR19ValidSIB26CellsR19Constraints = per.SizeRange(1, common.MaxPCI_OD_SIB1_r19)

var oDSIB1ConfigR19OffsetToPointAR19Constraints = per.Constrained(0, 2199)

var oDSIB1ConfigR19SsPBCHBlockPowerR19Constraints = per.Constrained(-60, 50)

const (
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms5    = 0
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms10   = 1
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms20   = 2
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms40   = 3
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms80   = 4
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms160  = 5
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Spare2 = 6
	OD_SIB1_Config_r19_Ssb_Periodicity_r19_Spare1 = 7
)

var oDSIB1ConfigR19SsbPeriodicityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms5, OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms10, OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms20, OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms40, OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms80, OD_SIB1_Config_r19_Ssb_Periodicity_r19_Ms160, OD_SIB1_Config_r19_Ssb_Periodicity_r19_Spare2, OD_SIB1_Config_r19_Ssb_Periodicity_r19_Spare1},
}

const (
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms20   = 0
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms40   = 1
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms80   = 2
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms160  = 3
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms320  = 4
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Spare3 = 5
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Spare2 = 6
	OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Spare1 = 7
)

var oDSIB1ConfigR19OdSIB1WindowDurationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms20, OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms40, OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms80, OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms160, OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Ms320, OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Spare3, OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Spare2, OD_SIB1_Config_r19_Od_SIB1_WindowDuration_r19_Spare1},
}

const (
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl0  = 0
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl1  = 1
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl2  = 2
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl4  = 3
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl8  = 4
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl10 = 5
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl20 = 6
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl40 = 7
	OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl80 = 8
)

var oDSIB1ConfigR19OdSIB1WindowStartOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl0, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl1, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl2, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl4, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl8, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl10, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl20, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl40, OD_SIB1_Config_r19_Od_SIB1_WindowStartOffset_r19_Sl80},
}

var oDSIB1ConfigR19KSSBR19Constraints = per.Constrained(0, 23)

var oDSIB1ConfigR19SsbPositionsInBurstR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inOneGroup"},
		{Name: "groupPresence", Optional: true},
	},
}

type OD_SIB1_Config_r19 struct {
	CarrierFreq_r19                        ARFCN_ValueNR
	ValidSIB26_Cells_r19                   []PhysCellId
	Sib1_TDD_UL_DL_ConfigurationCommon_r19 *TDD_UL_DL_ConfigCommon
	OffsetToPointA_r19                     *int64
	FrequencyBandList_r19                  *MultiFrequencyBandListNR_SIB
	Ss_PBCH_BlockPower_r19                 int64
	Ssb_PositionsInBurst_r19               struct {
		InOneGroup    per.BitString
		GroupPresence *per.BitString
	}
	Ssb_Periodicity_r19           *int64
	Od_SIB1_WindowDuration_r19    int64
	Od_SIB1_WindowStartOffset_r19 int64
	K_SSB_r19                     int64
	ControlResourceSetZero_r19    ControlResourceSetZero
	SearchSpaceZero_r19           SearchSpaceZero
	Sib1_RequestConfig_r19        SIB1_RequestConfig_r19
	Sib1_RequestConfigSUL_r19     *SIB1_RequestConfig_r19
}

func (ie *OD_SIB1_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(oDSIB1ConfigR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sib1_TDD_UL_DL_ConfigurationCommon_r19 != nil, ie.OffsetToPointA_r19 != nil, ie.FrequencyBandList_r19 != nil, ie.Ssb_Periodicity_r19 != nil, ie.Sib1_RequestConfigSUL_r19 != nil}); err != nil {
		return err
	}

	// 2. carrierFreq-r19: ref
	{
		if err := ie.CarrierFreq_r19.Encode(e); err != nil {
			return err
		}
	}

	// 3. validSIB26-Cells-r19: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(oDSIB1ConfigR19ValidSIB26CellsR19Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.ValidSIB26_Cells_r19))); err != nil {
			return err
		}
		for i := range ie.ValidSIB26_Cells_r19 {
			if err := ie.ValidSIB26_Cells_r19[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sib1-TDD-UL-DL-ConfigurationCommon-r19: ref
	{
		if ie.Sib1_TDD_UL_DL_ConfigurationCommon_r19 != nil {
			if err := ie.Sib1_TDD_UL_DL_ConfigurationCommon_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. offsetToPointA-r19: integer
	{
		if ie.OffsetToPointA_r19 != nil {
			if err := e.EncodeInteger(*ie.OffsetToPointA_r19, oDSIB1ConfigR19OffsetToPointAR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. frequencyBandList-r19: ref
	{
		if ie.FrequencyBandList_r19 != nil {
			if err := ie.FrequencyBandList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. ss-PBCH-BlockPower-r19: integer
	{
		if err := e.EncodeInteger(ie.Ss_PBCH_BlockPower_r19, oDSIB1ConfigR19SsPBCHBlockPowerR19Constraints); err != nil {
			return err
		}
	}

	// 8. ssb-PositionsInBurst-r19: sequence
	{
		{
			c := &ie.Ssb_PositionsInBurst_r19
			oDSIB1ConfigR19SsbPositionsInBurstR19Seq := e.NewSequenceEncoder(oDSIB1ConfigR19SsbPositionsInBurstR19Constraints)
			if err := oDSIB1ConfigR19SsbPositionsInBurstR19Seq.EncodePreamble([]bool{c.GroupPresence != nil}); err != nil {
				return err
			}
			if err := e.EncodeBitString(c.InOneGroup, per.FixedSize(8)); err != nil {
				return err
			}
			if c.GroupPresence != nil {
				if err := e.EncodeBitString((*c.GroupPresence), per.FixedSize(8)); err != nil {
					return err
				}
			}
		}
	}

	// 9. ssb-Periodicity-r19: enumerated
	{
		if ie.Ssb_Periodicity_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_Periodicity_r19, oDSIB1ConfigR19SsbPeriodicityR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. od-SIB1-WindowDuration-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Od_SIB1_WindowDuration_r19, oDSIB1ConfigR19OdSIB1WindowDurationR19Constraints); err != nil {
			return err
		}
	}

	// 11. od-SIB1-WindowStartOffset-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Od_SIB1_WindowStartOffset_r19, oDSIB1ConfigR19OdSIB1WindowStartOffsetR19Constraints); err != nil {
			return err
		}
	}

	// 12. k-SSB-r19: integer
	{
		if err := e.EncodeInteger(ie.K_SSB_r19, oDSIB1ConfigR19KSSBR19Constraints); err != nil {
			return err
		}
	}

	// 13. controlResourceSetZero-r19: ref
	{
		if err := ie.ControlResourceSetZero_r19.Encode(e); err != nil {
			return err
		}
	}

	// 14. searchSpaceZero-r19: ref
	{
		if err := ie.SearchSpaceZero_r19.Encode(e); err != nil {
			return err
		}
	}

	// 15. sib1-RequestConfig-r19: ref
	{
		if err := ie.Sib1_RequestConfig_r19.Encode(e); err != nil {
			return err
		}
	}

	// 16. sib1-RequestConfigSUL-r19: ref
	{
		if ie.Sib1_RequestConfigSUL_r19 != nil {
			if err := ie.Sib1_RequestConfigSUL_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OD_SIB1_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(oDSIB1ConfigR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. carrierFreq-r19: ref
	{
		if err := ie.CarrierFreq_r19.Decode(d); err != nil {
			return err
		}
	}

	// 3. validSIB26-Cells-r19: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(oDSIB1ConfigR19ValidSIB26CellsR19Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.ValidSIB26_Cells_r19 = make([]PhysCellId, n)
		for i := int64(0); i < n; i++ {
			if err := ie.ValidSIB26_Cells_r19[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sib1-TDD-UL-DL-ConfigurationCommon-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sib1_TDD_UL_DL_ConfigurationCommon_r19 = new(TDD_UL_DL_ConfigCommon)
			if err := ie.Sib1_TDD_UL_DL_ConfigurationCommon_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. offsetToPointA-r19: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(oDSIB1ConfigR19OffsetToPointAR19Constraints)
			if err != nil {
				return err
			}
			ie.OffsetToPointA_r19 = &v
		}
	}

	// 6. frequencyBandList-r19: ref
	{
		if seq.IsComponentPresent(4) {
			ie.FrequencyBandList_r19 = new(MultiFrequencyBandListNR_SIB)
			if err := ie.FrequencyBandList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. ss-PBCH-BlockPower-r19: integer
	{
		v5, err := d.DecodeInteger(oDSIB1ConfigR19SsPBCHBlockPowerR19Constraints)
		if err != nil {
			return err
		}
		ie.Ss_PBCH_BlockPower_r19 = v5
	}

	// 8. ssb-PositionsInBurst-r19: sequence
	{
		{
			c := &ie.Ssb_PositionsInBurst_r19
			oDSIB1ConfigR19SsbPositionsInBurstR19Seq := d.NewSequenceDecoder(oDSIB1ConfigR19SsbPositionsInBurstR19Constraints)
			if err := oDSIB1ConfigR19SsbPositionsInBurstR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				c.InOneGroup = v
			}
			if oDSIB1ConfigR19SsbPositionsInBurstR19Seq.IsComponentPresent(1) {
				c.GroupPresence = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*c.GroupPresence) = v
			}
		}
	}

	// 9. ssb-Periodicity-r19: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(oDSIB1ConfigR19SsbPeriodicityR19Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_Periodicity_r19 = &idx
		}
	}

	// 10. od-SIB1-WindowDuration-r19: enumerated
	{
		v8, err := d.DecodeEnumerated(oDSIB1ConfigR19OdSIB1WindowDurationR19Constraints)
		if err != nil {
			return err
		}
		ie.Od_SIB1_WindowDuration_r19 = v8
	}

	// 11. od-SIB1-WindowStartOffset-r19: enumerated
	{
		v9, err := d.DecodeEnumerated(oDSIB1ConfigR19OdSIB1WindowStartOffsetR19Constraints)
		if err != nil {
			return err
		}
		ie.Od_SIB1_WindowStartOffset_r19 = v9
	}

	// 12. k-SSB-r19: integer
	{
		v10, err := d.DecodeInteger(oDSIB1ConfigR19KSSBR19Constraints)
		if err != nil {
			return err
		}
		ie.K_SSB_r19 = v10
	}

	// 13. controlResourceSetZero-r19: ref
	{
		if err := ie.ControlResourceSetZero_r19.Decode(d); err != nil {
			return err
		}
	}

	// 14. searchSpaceZero-r19: ref
	{
		if err := ie.SearchSpaceZero_r19.Decode(d); err != nil {
			return err
		}
	}

	// 15. sib1-RequestConfig-r19: ref
	{
		if err := ie.Sib1_RequestConfig_r19.Decode(d); err != nil {
			return err
		}
	}

	// 16. sib1-RequestConfigSUL-r19: ref
	{
		if seq.IsComponentPresent(14) {
			ie.Sib1_RequestConfigSUL_r19 = new(SIB1_RequestConfig_r19)
			if err := ie.Sib1_RequestConfigSUL_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
