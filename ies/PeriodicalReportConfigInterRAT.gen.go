// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PeriodicalReportConfigInterRAT (line 13489).

var periodicalReportConfigInterRATConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportInterval"},
		{Name: "reportAmount"},
		{Name: "reportQuantity"},
		{Name: "maxReportCells"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	PeriodicalReportConfigInterRAT_ReportAmount_r1       = 0
	PeriodicalReportConfigInterRAT_ReportAmount_r2       = 1
	PeriodicalReportConfigInterRAT_ReportAmount_r4       = 2
	PeriodicalReportConfigInterRAT_ReportAmount_r8       = 3
	PeriodicalReportConfigInterRAT_ReportAmount_r16      = 4
	PeriodicalReportConfigInterRAT_ReportAmount_r32      = 5
	PeriodicalReportConfigInterRAT_ReportAmount_r64      = 6
	PeriodicalReportConfigInterRAT_ReportAmount_Infinity = 7
)

var periodicalReportConfigInterRATReportAmountConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicalReportConfigInterRAT_ReportAmount_r1, PeriodicalReportConfigInterRAT_ReportAmount_r2, PeriodicalReportConfigInterRAT_ReportAmount_r4, PeriodicalReportConfigInterRAT_ReportAmount_r8, PeriodicalReportConfigInterRAT_ReportAmount_r16, PeriodicalReportConfigInterRAT_ReportAmount_r32, PeriodicalReportConfigInterRAT_ReportAmount_r64, PeriodicalReportConfigInterRAT_ReportAmount_Infinity},
}

var periodicalReportConfigInterRATMaxReportCellsConstraints = per.Constrained(1, common.MaxCellReport)

const (
	PeriodicalReportConfigInterRAT_Ext_IncludeCommonLocationInfo_r16_True = 0
)

var periodicalReportConfigInterRATExtIncludeCommonLocationInfoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PeriodicalReportConfigInterRAT_Ext_IncludeCommonLocationInfo_r16_True},
}

var periodicalReportConfigInterRATExtIncludeBTMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfigInterRAT_Ext_IncludeBT_Meas_r16_Release = 0
	PeriodicalReportConfigInterRAT_Ext_IncludeBT_Meas_r16_Setup   = 1
)

var periodicalReportConfigInterRATExtIncludeWLANMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Release = 0
	PeriodicalReportConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Setup   = 1
)

var periodicalReportConfigInterRATExtIncludeSensorMeasR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PeriodicalReportConfigInterRAT_Ext_IncludeSensor_Meas_r16_Release = 0
	PeriodicalReportConfigInterRAT_Ext_IncludeSensor_Meas_r16_Setup   = 1
)

type PeriodicalReportConfigInterRAT struct {
	ReportInterval                ReportInterval
	ReportAmount                  int64
	ReportQuantity                MeasReportQuantity
	MaxReportCells                int64
	ReportQuantityUTRA_FDD_r16    *MeasReportQuantityUTRA_FDD_r16
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
	ReportQuantityRelay_r17 *SL_MeasReportQuantity_r16
}

func (ie *PeriodicalReportConfigInterRAT) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(periodicalReportConfigInterRATConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReportQuantityUTRA_FDD_r16 != nil
	hasExtGroup1 := ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil
	hasExtGroup2 := ie.ReportQuantityRelay_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. reportInterval: ref
	{
		if err := ie.ReportInterval.Encode(e); err != nil {
			return err
		}
	}

	// 3. reportAmount: enumerated
	{
		if err := e.EncodeEnumerated(ie.ReportAmount, periodicalReportConfigInterRATReportAmountConstraints); err != nil {
			return err
		}
	}

	// 4. reportQuantity: ref
	{
		if err := ie.ReportQuantity.Encode(e); err != nil {
			return err
		}
	}

	// 5. maxReportCells: integer
	{
		if err := e.EncodeInteger(ie.MaxReportCells, periodicalReportConfigInterRATMaxReportCellsConstraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportQuantityUTRA-FDD-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportQuantityUTRA_FDD_r16 != nil}); err != nil {
				return err
			}

			if ie.ReportQuantityUTRA_FDD_r16 != nil {
				if err := ie.ReportQuantityUTRA_FDD_r16.Encode(ex); err != nil {
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
					{Name: "includeCommonLocationInfo-r16", Optional: true},
					{Name: "includeBT-Meas-r16", Optional: true},
					{Name: "includeWLAN-Meas-r16", Optional: true},
					{Name: "includeSensor-Meas-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IncludeCommonLocationInfo_r16 != nil, ie.IncludeBT_Meas_r16 != nil, ie.IncludeWLAN_Meas_r16 != nil, ie.IncludeSensor_Meas_r16 != nil}); err != nil {
				return err
			}

			if ie.IncludeCommonLocationInfo_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IncludeCommonLocationInfo_r16, periodicalReportConfigInterRATExtIncludeCommonLocationInfoR16Constraints); err != nil {
					return err
				}
			}

			if ie.IncludeBT_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigInterRATExtIncludeBTMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeBT_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeBT_Meas_r16).Choice {
				case PeriodicalReportConfigInterRAT_Ext_IncludeBT_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfigInterRAT_Ext_IncludeBT_Meas_r16_Setup:
					if err := (*ie.IncludeBT_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeWLAN_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigInterRATExtIncludeWLANMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeWLAN_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeWLAN_Meas_r16).Choice {
				case PeriodicalReportConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Setup:
					if err := (*ie.IncludeWLAN_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncludeSensor_Meas_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(periodicalReportConfigInterRATExtIncludeSensorMeasR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IncludeSensor_Meas_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IncludeSensor_Meas_r16).Choice {
				case PeriodicalReportConfigInterRAT_Ext_IncludeSensor_Meas_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PeriodicalReportConfigInterRAT_Ext_IncludeSensor_Meas_r16_Setup:
					if err := (*ie.IncludeSensor_Meas_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportQuantityRelay-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportQuantityRelay_r17 != nil}); err != nil {
				return err
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

func (ie *PeriodicalReportConfigInterRAT) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(periodicalReportConfigInterRATConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. reportInterval: ref
	{
		if err := ie.ReportInterval.Decode(d); err != nil {
			return err
		}
	}

	// 3. reportAmount: enumerated
	{
		v1, err := d.DecodeEnumerated(periodicalReportConfigInterRATReportAmountConstraints)
		if err != nil {
			return err
		}
		ie.ReportAmount = v1
	}

	// 4. reportQuantity: ref
	{
		if err := ie.ReportQuantity.Decode(d); err != nil {
			return err
		}
	}

	// 5. maxReportCells: integer
	{
		v3, err := d.DecodeInteger(periodicalReportConfigInterRATMaxReportCellsConstraints)
		if err != nil {
			return err
		}
		ie.MaxReportCells = v3
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
				{Name: "reportQuantityUTRA-FDD-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ReportQuantityUTRA_FDD_r16 = new(MeasReportQuantityUTRA_FDD_r16)
			if err := ie.ReportQuantityUTRA_FDD_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "includeCommonLocationInfo-r16", Optional: true},
				{Name: "includeBT-Meas-r16", Optional: true},
				{Name: "includeWLAN-Meas-r16", Optional: true},
				{Name: "includeSensor-Meas-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(periodicalReportConfigInterRATExtIncludeCommonLocationInfoR16Constraints)
			if err != nil {
				return err
			}
			ie.IncludeCommonLocationInfo_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.IncludeBT_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BT_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigInterRATExtIncludeBTMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeBT_Meas_r16).Choice = int(index)
			switch index {
			case PeriodicalReportConfigInterRAT_Ext_IncludeBT_Meas_r16_Release:
				(*ie.IncludeBT_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfigInterRAT_Ext_IncludeBT_Meas_r16_Setup:
				(*ie.IncludeBT_Meas_r16).Setup = new(BT_NameList_r16)
				if err := (*ie.IncludeBT_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.IncludeWLAN_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *WLAN_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigInterRATExtIncludeWLANMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeWLAN_Meas_r16).Choice = int(index)
			switch index {
			case PeriodicalReportConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Release:
				(*ie.IncludeWLAN_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfigInterRAT_Ext_IncludeWLAN_Meas_r16_Setup:
				(*ie.IncludeWLAN_Meas_r16).Setup = new(WLAN_NameList_r16)
				if err := (*ie.IncludeWLAN_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.IncludeSensor_Meas_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Sensor_NameList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(periodicalReportConfigInterRATExtIncludeSensorMeasR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IncludeSensor_Meas_r16).Choice = int(index)
			switch index {
			case PeriodicalReportConfigInterRAT_Ext_IncludeSensor_Meas_r16_Release:
				(*ie.IncludeSensor_Meas_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PeriodicalReportConfigInterRAT_Ext_IncludeSensor_Meas_r16_Setup:
				(*ie.IncludeSensor_Meas_r16).Setup = new(Sensor_NameList_r16)
				if err := (*ie.IncludeSensor_Meas_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "reportQuantityRelay-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ReportQuantityRelay_r17 = new(SL_MeasReportQuantity_r16)
			if err := ie.ReportQuantityRelay_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
