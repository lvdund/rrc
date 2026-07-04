// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PeriodicalReportConfig (line 13836).

var periodicalReportConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rsType"},
		{Name: "reportInterval"},
		{Name: "reportAmount"},
		{Name: "reportQuantityCell"},
		{Name: "maxReportCells"},
		{Name: "reportQuantityRS-Indexes", Optional: true},
		{Name: "maxNrofRS-IndexesToReport", Optional: true},
		{Name: "includeBeamMeasurements"},
		{Name: "useAllowedCellList"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	PeriodicalReportConfig_ReportAmount_r1       = 0
	PeriodicalReportConfig_ReportAmount_r2       = 1
	PeriodicalReportConfig_ReportAmount_r4       = 2
	PeriodicalReportConfig_ReportAmount_r8       = 3
	PeriodicalReportConfig_ReportAmount_r16      = 4
	PeriodicalReportConfig_ReportAmount_r32      = 5
	PeriodicalReportConfig_ReportAmount_r64      = 6
	PeriodicalReportConfig_ReportAmount_Infinity = 7
)

var periodicalReportConfigReportAmountConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicalReportConfig_ReportAmount_r1, PeriodicalReportConfig_ReportAmount_r2, PeriodicalReportConfig_ReportAmount_r4, PeriodicalReportConfig_ReportAmount_r8, PeriodicalReportConfig_ReportAmount_r16, PeriodicalReportConfig_ReportAmount_r32, PeriodicalReportConfig_ReportAmount_r64, PeriodicalReportConfig_ReportAmount_Infinity},
}

var periodicalReportConfigMaxReportCellsConstraints = per.Constrained(1, common.MaxCellReport)

var periodicalReportConfigMaxNrofRSIndexesToReportConstraints = per.Constrained(1, common.MaxNrofIndexesToReport)

const (
	PeriodicalReportConfig_Ext_IncludeCommonLocationInfo_r16_True = 0
)

var periodicalReportConfigExtIncludeCommonLocationInfoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicalReportConfig_Ext_IncludeCommonLocationInfo_r16_True},
}

var periodicalReportConfigExtIncludeBTMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfig_Ext_IncludeBT_Meas_r16_Release = 0
	PeriodicalReportConfig_Ext_IncludeBT_Meas_r16_Setup   = 1
)

var periodicalReportConfigExtIncludeWLANMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfig_Ext_IncludeWLAN_Meas_r16_Release = 0
	PeriodicalReportConfig_Ext_IncludeWLAN_Meas_r16_Setup   = 1
)

var periodicalReportConfigExtIncludeSensorMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfig_Ext_IncludeSensor_Meas_r16_Release = 0
	PeriodicalReportConfig_Ext_IncludeSensor_Meas_r16_Setup   = 1
)

var periodicalReportConfigExtUlDelayValueConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfig_Ext_Ul_DelayValueConfig_r16_Release = 0
	PeriodicalReportConfig_Ext_Ul_DelayValueConfig_r16_Setup   = 1
)

const (
	PeriodicalReportConfig_Ext_ReportAddNeighMeas_r16_Setup = 0
)

var periodicalReportConfigExtReportAddNeighMeasR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicalReportConfig_Ext_ReportAddNeighMeas_r16_Setup},
}

var periodicalReportConfigExtUlExcessDelayConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfig_Ext_Ul_ExcessDelayConfig_r17_Release = 0
	PeriodicalReportConfig_Ext_Ul_ExcessDelayConfig_r17_Setup   = 1
)

const (
	PeriodicalReportConfig_Ext_CoarseLocationRequest_r17_True = 0
)

var periodicalReportConfigExtCoarseLocationRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicalReportConfig_Ext_CoarseLocationRequest_r17_True},
}

type PeriodicalReportConfig struct {
	RsType                        NR_RS_Type
	ReportInterval                ReportInterval
	ReportAmount                  int64
	ReportQuantityCell            MeasReportQuantity
	MaxReportCells                int64
	ReportQuantityRS_Indexes      *MeasReportQuantity
	MaxNrofRS_IndexesToReport     *int64
	IncludeBeamMeasurements       bool
	UseAllowedCellList            bool
	MeasRSSI_ReportConfig_r16     *MeasRSSI_ReportConfig_r16
	IncludeCommonLocationInfo_r16 *int64
	IncludeBT_Meas_r16            *struct {
		Choice  int
		Release *struct{}
		Setup   *BT_NameList_r16
	}
	IncludeWLAN_Meas_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *WLAN_NameList_r16
	}
	IncludeSensor_Meas_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *Sensor_NameList_r16
	}
	Ul_DelayValueConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_DelayValueConfig_r16
	}
	ReportAddNeighMeas_r16   *int64
	Ul_ExcessDelayConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_ExcessDelayConfig_r17
	}
	CoarseLocationRequest_r17 *int64
	ReportQuantityRelay_r17   *SL_MeasReportQuantity_r16
}

func (ie *PeriodicalReportConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(periodicalReportConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MeasRSSI_ReportConfig_r16 != nil || ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil || ie.Ul_DelayValueConfig_r16 != nil || ie.ReportAddNeighMeas_r16 != nil
	hasExtGroup1 := ie.Ul_ExcessDelayConfig_r17 != nil || ie.CoarseLocationRequest_r17 != nil || ie.ReportQuantityRelay_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReportQuantityRS_Indexes != nil, ie.MaxNrofRS_IndexesToReport != nil}); err != nil {
		return err
	}

	// 3. rsType: ref
	{
		if err := ie.RsType.Encode(e); err != nil {
			return err
		}
	}

	// 4. reportInterval: ref
	{
		if err := ie.ReportInterval.Encode(e); err != nil {
			return err
		}
	}

	// 5. reportAmount: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount, periodicalReportConfigReportAmountConstraints); err != nil {
			return err
		}
	}

	// 6. reportQuantityCell: ref
	{
		if err := ie.ReportQuantityCell.Encode(e); err != nil {
			return err
		}
	}

	// 7. maxReportCells: integer
	{
		if err := e.EncodeInteger(ie.MaxReportCells, periodicalReportConfigMaxReportCellsConstraints); err != nil {
			return err
		}
	}

	// 8. reportQuantityRS-Indexes: ref
	{
		if ie.ReportQuantityRS_Indexes != nil {
			if err := ie.ReportQuantityRS_Indexes.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. maxNrofRS-IndexesToReport: integer
	{
		if ie.MaxNrofRS_IndexesToReport != nil {
			if err := e.EncodeInteger(*ie.MaxNrofRS_IndexesToReport, periodicalReportConfigMaxNrofRSIndexesToReportConstraints); err != nil {
				return err
			}
		}
	}

	// 10. includeBeamMeasurements: boolean
	{
		if err := e.EncodeBoolean(ie.IncludeBeamMeasurements); err != nil {
			return err
		}
	}

	// 11. useAllowedCellList: boolean
	{
		if err := e.EncodeBoolean(ie.UseAllowedCellList); err != nil {
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
					{Name: "measRSSI-ReportConfig-r16", Optional: true},
					{Name: "includeCommonLocationInfo-r16", Optional: true},
					{Name: "includeBT-Meas-r16", Optional: true},
					{Name: "includeWLAN-Meas-r16", Optional: true},
					{Name: "includeSensor-Meas-r16", Optional: true},
					{Name: "ul-DelayValueConfig-r16", Optional: true},
					{Name: "reportAddNeighMeas-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasRSSI_ReportConfig_r16 != nil, ie.IncludeCommonLocationInfo_r16 != nil, ie.IncludeBT_Meas_r16 != nil, ie.IncludeWLAN_Meas_r16 != nil, ie.IncludeSensor_Meas_r16 != nil, ie.Ul_DelayValueConfig_r16 != nil, ie.ReportAddNeighMeas_r16 != nil}); err != nil {
				return err
			}

			if ie.MeasRSSI_ReportConfig_r16 != nil {
				if err := ie.MeasRSSI_ReportConfig_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.IncludeCommonLocationInfo_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IncludeCommonLocationInfo_r16, periodicalReportConfigExtIncludeCommonLocationInfoR16Constraints); err != nil {
					return err
				}
			}

			if ie.IncludeBT_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigExtIncludeBTMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeBT_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeBT_Meas_r16).Choice {
				case PeriodicalReportConfig_Ext_IncludeBT_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfig_Ext_IncludeBT_Meas_r16_Setup:
					if err := (*ie.IncludeBT_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeWLAN_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigExtIncludeWLANMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeWLAN_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeWLAN_Meas_r16).Choice {
				case PeriodicalReportConfig_Ext_IncludeWLAN_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfig_Ext_IncludeWLAN_Meas_r16_Setup:
					if err := (*ie.IncludeWLAN_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeSensor_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigExtIncludeSensorMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeSensor_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeSensor_Meas_r16).Choice {
				case PeriodicalReportConfig_Ext_IncludeSensor_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfig_Ext_IncludeSensor_Meas_r16_Setup:
					if err := (*ie.IncludeSensor_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ul_DelayValueConfig_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigExtUlDelayValueConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_DelayValueConfig_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_DelayValueConfig_r16).Choice {
				case PeriodicalReportConfig_Ext_Ul_DelayValueConfig_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfig_Ext_Ul_DelayValueConfig_r16_Setup:
					if err := (*ie.Ul_DelayValueConfig_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ReportAddNeighMeas_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ReportAddNeighMeas_r16, periodicalReportConfigExtReportAddNeighMeasR16Constraints); err != nil {
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
					{Name: "ul-ExcessDelayConfig-r17", Optional: true},
					{Name: "coarseLocationRequest-r17", Optional: true},
					{Name: "reportQuantityRelay-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ul_ExcessDelayConfig_r17 != nil, ie.CoarseLocationRequest_r17 != nil, ie.ReportQuantityRelay_r17 != nil}); err != nil {
				return err
			}

			if ie.Ul_ExcessDelayConfig_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigExtUlExcessDelayConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_ExcessDelayConfig_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_ExcessDelayConfig_r17).Choice {
				case PeriodicalReportConfig_Ext_Ul_ExcessDelayConfig_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfig_Ext_Ul_ExcessDelayConfig_r17_Setup:
					if err := (*ie.Ul_ExcessDelayConfig_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.CoarseLocationRequest_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.CoarseLocationRequest_r17, periodicalReportConfigExtCoarseLocationRequestR17Constraints); err != nil {
					return err
				}
			}

			if ie.ReportQuantityRelay_r17 != nil {
				if err := ie.ReportQuantityRelay_r17.Encode(ex); err != nil {
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

func (ie *PeriodicalReportConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(periodicalReportConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rsType: ref
	{
		if err := ie.RsType.Decode(d); err != nil {
			return err
		}
	}

	// 4. reportInterval: ref
	{
		if err := ie.ReportInterval.Decode(d); err != nil {
			return err
		}
	}

	// 5. reportAmount: enumerated
	{
		v2, err := d.DecodeEnumerated(periodicalReportConfigReportAmountConstraints)
		if err != nil {
			return err
		}
		ie.ReportAmount = v2
	}

	// 6. reportQuantityCell: ref
	{
		if err := ie.ReportQuantityCell.Decode(d); err != nil {
			return err
		}
	}

	// 7. maxReportCells: integer
	{
		v4, err := d.DecodeInteger(periodicalReportConfigMaxReportCellsConstraints)
		if err != nil {
			return err
		}
		ie.MaxReportCells = v4
	}

	// 8. reportQuantityRS-Indexes: ref
	{
		if seq.IsComponentPresent(5) {
			ie.ReportQuantityRS_Indexes = new(MeasReportQuantity)
			if err := ie.ReportQuantityRS_Indexes.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. maxNrofRS-IndexesToReport: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(periodicalReportConfigMaxNrofRSIndexesToReportConstraints)
			if err != nil {
				return err
			}
			ie.MaxNrofRS_IndexesToReport = &v
		}
	}

	// 10. includeBeamMeasurements: boolean
	{
		v7, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.IncludeBeamMeasurements = v7
	}

	// 11. useAllowedCellList: boolean
	{
		v8, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.UseAllowedCellList = v8
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
				{Name: "measRSSI-ReportConfig-r16", Optional: true},
				{Name: "includeCommonLocationInfo-r16", Optional: true},
				{Name: "includeBT-Meas-r16", Optional: true},
				{Name: "includeWLAN-Meas-r16", Optional: true},
				{Name: "includeSensor-Meas-r16", Optional: true},
				{Name: "ul-DelayValueConfig-r16", Optional: true},
				{Name: "reportAddNeighMeas-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MeasRSSI_ReportConfig_r16 = new(MeasRSSI_ReportConfig_r16)
			if err := ie.MeasRSSI_ReportConfig_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(periodicalReportConfigExtIncludeCommonLocationInfoR16Constraints)
			if err != nil {
				return err
			}
			ie.IncludeCommonLocationInfo_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.IncludeBT_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BT_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigExtIncludeBTMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeBT_Meas_r16).Choice = int(index)
			switch index {
			case PeriodicalReportConfig_Ext_IncludeBT_Meas_r16_Release:
				(*ie.IncludeBT_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfig_Ext_IncludeBT_Meas_r16_Setup:
				(*ie.IncludeBT_Meas_r16).Setup = new(BT_NameList_r16)
				if err := (*ie.IncludeBT_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.IncludeWLAN_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *WLAN_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigExtIncludeWLANMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeWLAN_Meas_r16).Choice = int(index)
			switch index {
			case PeriodicalReportConfig_Ext_IncludeWLAN_Meas_r16_Release:
				(*ie.IncludeWLAN_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfig_Ext_IncludeWLAN_Meas_r16_Setup:
				(*ie.IncludeWLAN_Meas_r16).Setup = new(WLAN_NameList_r16)
				if err := (*ie.IncludeWLAN_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.IncludeSensor_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Sensor_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigExtIncludeSensorMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeSensor_Meas_r16).Choice = int(index)
			switch index {
			case PeriodicalReportConfig_Ext_IncludeSensor_Meas_r16_Release:
				(*ie.IncludeSensor_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfig_Ext_IncludeSensor_Meas_r16_Setup:
				(*ie.IncludeSensor_Meas_r16).Setup = new(Sensor_NameList_r16)
				if err := (*ie.IncludeSensor_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Ul_DelayValueConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_DelayValueConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigExtUlDelayValueConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_DelayValueConfig_r16).Choice = int(index)
			switch index {
			case PeriodicalReportConfig_Ext_Ul_DelayValueConfig_r16_Release:
				(*ie.Ul_DelayValueConfig_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfig_Ext_Ul_DelayValueConfig_r16_Setup:
				(*ie.Ul_DelayValueConfig_r16).Setup = new(UL_DelayValueConfig_r16)
				if err := (*ie.Ul_DelayValueConfig_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(periodicalReportConfigExtReportAddNeighMeasR16Constraints)
			if err != nil {
				return err
			}
			ie.ReportAddNeighMeas_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ul-ExcessDelayConfig-r17", Optional: true},
				{Name: "coarseLocationRequest-r17", Optional: true},
				{Name: "reportQuantityRelay-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ul_ExcessDelayConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_ExcessDelayConfig_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigExtUlExcessDelayConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_ExcessDelayConfig_r17).Choice = int(index)
			switch index {
			case PeriodicalReportConfig_Ext_Ul_ExcessDelayConfig_r17_Release:
				(*ie.Ul_ExcessDelayConfig_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfig_Ext_Ul_ExcessDelayConfig_r17_Setup:
				(*ie.Ul_ExcessDelayConfig_r17).Setup = new(UL_ExcessDelayConfig_r17)
				if err := (*ie.Ul_ExcessDelayConfig_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(periodicalReportConfigExtCoarseLocationRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.CoarseLocationRequest_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.ReportQuantityRelay_r17 = new(SL_MeasReportQuantity_r16)
			if err := ie.ReportQuantityRelay_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
