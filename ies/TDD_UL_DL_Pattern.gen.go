// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDD-UL-DL-Pattern (line 16074).

var tDDULDLPatternConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-UL-TransmissionPeriodicity"},
		{Name: "nrofDownlinkSlots"},
		{Name: "nrofDownlinkSymbols"},
		{Name: "nrofUplinkSlots"},
		{Name: "nrofUplinkSymbols"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms0p5   = 0
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms0p625 = 1
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms1     = 2
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms1p25  = 3
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms2     = 4
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms2p5   = 5
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms5     = 6
	TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms10    = 7
)

var tDDULDLPatternDlULTransmissionPeriodicityConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms0p5, TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms0p625, TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms1, TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms1p25, TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms2, TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms2p5, TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms5, TDD_UL_DL_Pattern_Dl_UL_TransmissionPeriodicity_Ms10},
}

var tDDULDLPatternNrofDownlinkSlotsConstraints = per.Constrained(0, common.MaxNrofSlots)

var tDDULDLPatternNrofDownlinkSymbolsConstraints = per.Constrained(0, common.MaxNrofSymbols_1)

var tDDULDLPatternNrofUplinkSlotsConstraints = per.Constrained(0, common.MaxNrofSlots)

var tDDULDLPatternNrofUplinkSymbolsConstraints = per.Constrained(0, common.MaxNrofSymbols_1)

const (
	TDD_UL_DL_Pattern_Ext_Dl_UL_TransmissionPeriodicity_v1530_Ms3 = 0
	TDD_UL_DL_Pattern_Ext_Dl_UL_TransmissionPeriodicity_v1530_Ms4 = 1
)

var tDDULDLPatternExtDlULTransmissionPeriodicityV1530Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TDD_UL_DL_Pattern_Ext_Dl_UL_TransmissionPeriodicity_v1530_Ms3, TDD_UL_DL_Pattern_Ext_Dl_UL_TransmissionPeriodicity_v1530_Ms4},
}

var tDDULDLPatternSbfdStartingSlotIndexR19Constraints = per.Constrained(0, common.MaxNrofSlots_1)

var tDDULDLPatternSbfdStartingSymbolIndexR19Constraints = per.Constrained(0, common.MaxNrofSymbols_1)

var tDDULDLPatternSbfdEndingSlotIndexR19Constraints = per.Constrained(0, common.MaxNrofSlots_1)

var tDDULDLPatternSbfdEndingSymbolIndexR19Constraints = per.Constrained(0, common.MaxNrofSymbols_1)

type TDD_UL_DL_Pattern struct {
	Dl_UL_TransmissionPeriodicity       int64
	NrofDownlinkSlots                   int64
	NrofDownlinkSymbols                 int64
	NrofUplinkSlots                     int64
	NrofUplinkSymbols                   int64
	Dl_UL_TransmissionPeriodicity_v1530 *int64
	Sbfd_StartingSlotIndex_r19          *int64
	Sbfd_StartingSymbolIndex_r19        *int64
	Sbfd_EndingSlotIndex_r19            *int64
	Sbfd_EndingSymbolIndex_r19          *int64
}

func (ie *TDD_UL_DL_Pattern) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tDDULDLPatternConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dl_UL_TransmissionPeriodicity_v1530 != nil
	hasExtGroup1 := ie.Sbfd_StartingSlotIndex_r19 != nil || ie.Sbfd_StartingSymbolIndex_r19 != nil || ie.Sbfd_EndingSlotIndex_r19 != nil || ie.Sbfd_EndingSymbolIndex_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. dl-UL-TransmissionPeriodicity: enumerated
	{
		if err := e.EncodeEnumerated(ie.Dl_UL_TransmissionPeriodicity, tDDULDLPatternDlULTransmissionPeriodicityConstraints); err != nil {
			return err
		}
	}

	// 3. nrofDownlinkSlots: integer
	{
		if err := e.EncodeInteger(ie.NrofDownlinkSlots, tDDULDLPatternNrofDownlinkSlotsConstraints); err != nil {
			return err
		}
	}

	// 4. nrofDownlinkSymbols: integer
	{
		if err := e.EncodeInteger(ie.NrofDownlinkSymbols, tDDULDLPatternNrofDownlinkSymbolsConstraints); err != nil {
			return err
		}
	}

	// 5. nrofUplinkSlots: integer
	{
		if err := e.EncodeInteger(ie.NrofUplinkSlots, tDDULDLPatternNrofUplinkSlotsConstraints); err != nil {
			return err
		}
	}

	// 6. nrofUplinkSymbols: integer
	{
		if err := e.EncodeInteger(ie.NrofUplinkSymbols, tDDULDLPatternNrofUplinkSymbolsConstraints); err != nil {
			return err
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
					{Name: "dl-UL-TransmissionPeriodicity-v1530", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dl_UL_TransmissionPeriodicity_v1530 != nil}); err != nil {
				return err
			}

			if ie.Dl_UL_TransmissionPeriodicity_v1530 != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_UL_TransmissionPeriodicity_v1530, tDDULDLPatternExtDlULTransmissionPeriodicityV1530Constraints); err != nil {
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
					{Name: "sbfd-StartingSlotIndex-r19", Optional: true},
					{Name: "sbfd-StartingSymbolIndex-r19", Optional: true},
					{Name: "sbfd-EndingSlotIndex-r19", Optional: true},
					{Name: "sbfd-EndingSymbolIndex-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sbfd_StartingSlotIndex_r19 != nil, ie.Sbfd_StartingSymbolIndex_r19 != nil, ie.Sbfd_EndingSlotIndex_r19 != nil, ie.Sbfd_EndingSymbolIndex_r19 != nil}); err != nil {
				return err
			}

			if ie.Sbfd_StartingSlotIndex_r19 != nil {
				if err := ex.EncodeInteger(*ie.Sbfd_StartingSlotIndex_r19, tDDULDLPatternSbfdStartingSlotIndexR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sbfd_StartingSymbolIndex_r19 != nil {
				if err := ex.EncodeInteger(*ie.Sbfd_StartingSymbolIndex_r19, tDDULDLPatternSbfdStartingSymbolIndexR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sbfd_EndingSlotIndex_r19 != nil {
				if err := ex.EncodeInteger(*ie.Sbfd_EndingSlotIndex_r19, tDDULDLPatternSbfdEndingSlotIndexR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sbfd_EndingSymbolIndex_r19 != nil {
				if err := ex.EncodeInteger(*ie.Sbfd_EndingSymbolIndex_r19, tDDULDLPatternSbfdEndingSymbolIndexR19Constraints); err != nil {
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

func (ie *TDD_UL_DL_Pattern) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tDDULDLPatternConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. dl-UL-TransmissionPeriodicity: enumerated
	{
		v0, err := d.DecodeEnumerated(tDDULDLPatternDlULTransmissionPeriodicityConstraints)
		if err != nil {
			return err
		}
		ie.Dl_UL_TransmissionPeriodicity = v0
	}

	// 3. nrofDownlinkSlots: integer
	{
		v1, err := d.DecodeInteger(tDDULDLPatternNrofDownlinkSlotsConstraints)
		if err != nil {
			return err
		}
		ie.NrofDownlinkSlots = v1
	}

	// 4. nrofDownlinkSymbols: integer
	{
		v2, err := d.DecodeInteger(tDDULDLPatternNrofDownlinkSymbolsConstraints)
		if err != nil {
			return err
		}
		ie.NrofDownlinkSymbols = v2
	}

	// 5. nrofUplinkSlots: integer
	{
		v3, err := d.DecodeInteger(tDDULDLPatternNrofUplinkSlotsConstraints)
		if err != nil {
			return err
		}
		ie.NrofUplinkSlots = v3
	}

	// 6. nrofUplinkSymbols: integer
	{
		v4, err := d.DecodeInteger(tDDULDLPatternNrofUplinkSymbolsConstraints)
		if err != nil {
			return err
		}
		ie.NrofUplinkSymbols = v4
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
				{Name: "dl-UL-TransmissionPeriodicity-v1530", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(tDDULDLPatternExtDlULTransmissionPeriodicityV1530Constraints)
			if err != nil {
				return err
			}
			ie.Dl_UL_TransmissionPeriodicity_v1530 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sbfd-StartingSlotIndex-r19", Optional: true},
				{Name: "sbfd-StartingSymbolIndex-r19", Optional: true},
				{Name: "sbfd-EndingSlotIndex-r19", Optional: true},
				{Name: "sbfd-EndingSymbolIndex-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(tDDULDLPatternSbfdStartingSlotIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_StartingSlotIndex_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(tDDULDLPatternSbfdStartingSymbolIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_StartingSymbolIndex_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(tDDULDLPatternSbfdEndingSlotIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_EndingSlotIndex_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(tDDULDLPatternSbfdEndingSymbolIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_EndingSymbolIndex_r19 = &v
		}
	}

	return nil
}
