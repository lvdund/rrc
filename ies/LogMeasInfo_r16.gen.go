// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LogMeasInfo-r16 (line 3003).

var logMeasInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "locationInfo-r16", Optional: true},
		{Name: "relativeTimeStamp-r16"},
		{Name: "servCellIdentity-r16", Optional: true},
		{Name: "measResultServingCell-r16", Optional: true},
		{Name: "measResultNeighCells-r16"},
		{Name: "anyCellSelectionDetected-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var logMeasInfoR16RelativeTimeStampR16Constraints = per.Constrained(0, 7200)

const (
	LogMeasInfo_r16_AnyCellSelectionDetected_r16_True = 0
)

var logMeasInfoR16AnyCellSelectionDetectedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogMeasInfo_r16_AnyCellSelectionDetected_r16_True},
}

const (
	LogMeasInfo_r16_Ext_InDeviceCoexDetected_r17_True = 0
)

var logMeasInfoR16ExtInDeviceCoexDetectedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LogMeasInfo_r16_Ext_InDeviceCoexDetected_r17_True},
}

var logMeasInfoR16MeasResultNeighCellsR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "measResultNeighCellListNR", Optional: true},
		{Name: "measResultNeighCellListEUTRA", Optional: true},
	},
}

type LogMeasInfo_r16 struct {
	LocationInfo_r16          *LocationInfo_r16
	RelativeTimeStamp_r16     int64
	ServCellIdentity_r16      *CGI_Info_Logging_r16
	MeasResultServingCell_r16 *MeasResultServingCell_r16
	MeasResultNeighCells_r16  struct {
		MeasResultNeighCellListNR    *MeasResultListLogging2NR_r16
		MeasResultNeighCellListEUTRA *MeasResultList2EUTRA_r16
	}
	AnyCellSelectionDetected_r16 *int64
	InDeviceCoexDetected_r17     *int64
	Nsag_ID_r19                  *NSAG_ID_r17
	ReselectedCellId_r19         *CGI_Info_Logging_r16
}

func (ie *LogMeasInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(logMeasInfoR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.InDeviceCoexDetected_r17 != nil
	hasExtGroup1 := ie.Nsag_ID_r19 != nil || ie.ReselectedCellId_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LocationInfo_r16 != nil, ie.ServCellIdentity_r16 != nil, ie.MeasResultServingCell_r16 != nil, ie.AnyCellSelectionDetected_r16 != nil}); err != nil {
		return err
	}

	// 3. locationInfo-r16: ref
	{
		if ie.LocationInfo_r16 != nil {
			if err := ie.LocationInfo_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. relativeTimeStamp-r16: integer
	{
		if err := e.EncodeInteger(ie.RelativeTimeStamp_r16, logMeasInfoR16RelativeTimeStampR16Constraints); err != nil {
			return err
		}
	}

	// 5. servCellIdentity-r16: ref
	{
		if ie.ServCellIdentity_r16 != nil {
			if err := ie.ServCellIdentity_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. measResultServingCell-r16: ref
	{
		if ie.MeasResultServingCell_r16 != nil {
			if err := ie.MeasResultServingCell_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. measResultNeighCells-r16: sequence
	{
		{
			c := &ie.MeasResultNeighCells_r16
			logMeasInfoR16MeasResultNeighCellsR16Seq := e.NewSequenceEncoder(logMeasInfoR16MeasResultNeighCellsR16Constraints)
			if err := logMeasInfoR16MeasResultNeighCellsR16Seq.EncodePreamble([]bool{c.MeasResultNeighCellListNR != nil, c.MeasResultNeighCellListEUTRA != nil}); err != nil {
				return err
			}
			if c.MeasResultNeighCellListNR != nil {
				if err := c.MeasResultNeighCellListNR.Encode(e); err != nil {
					return err
				}
			}
			if c.MeasResultNeighCellListEUTRA != nil {
				if err := c.MeasResultNeighCellListEUTRA.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. anyCellSelectionDetected-r16: enumerated
	{
		if ie.AnyCellSelectionDetected_r16 != nil {
			if err := e.EncodeEnumerated(*ie.AnyCellSelectionDetected_r16, logMeasInfoR16AnyCellSelectionDetectedR16Constraints); err != nil {
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
					{Name: "inDeviceCoexDetected-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.InDeviceCoexDetected_r17 != nil}); err != nil {
				return err
			}

			if ie.InDeviceCoexDetected_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.InDeviceCoexDetected_r17, logMeasInfoR16ExtInDeviceCoexDetectedR17Constraints); err != nil {
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
					{Name: "nsag-ID-r19", Optional: true},
					{Name: "reselectedCellId-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Nsag_ID_r19 != nil, ie.ReselectedCellId_r19 != nil}); err != nil {
				return err
			}

			if ie.Nsag_ID_r19 != nil {
				if err := ie.Nsag_ID_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ReselectedCellId_r19 != nil {
				if err := ie.ReselectedCellId_r19.Encode(ex); err != nil {
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

func (ie *LogMeasInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(logMeasInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. locationInfo-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.LocationInfo_r16 = new(LocationInfo_r16)
			if err := ie.LocationInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. relativeTimeStamp-r16: integer
	{
		v1, err := d.DecodeInteger(logMeasInfoR16RelativeTimeStampR16Constraints)
		if err != nil {
			return err
		}
		ie.RelativeTimeStamp_r16 = v1
	}

	// 5. servCellIdentity-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.ServCellIdentity_r16 = new(CGI_Info_Logging_r16)
			if err := ie.ServCellIdentity_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. measResultServingCell-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.MeasResultServingCell_r16 = new(MeasResultServingCell_r16)
			if err := ie.MeasResultServingCell_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. measResultNeighCells-r16: sequence
	{
		{
			c := &ie.MeasResultNeighCells_r16
			logMeasInfoR16MeasResultNeighCellsR16Seq := d.NewSequenceDecoder(logMeasInfoR16MeasResultNeighCellsR16Constraints)
			if err := logMeasInfoR16MeasResultNeighCellsR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if logMeasInfoR16MeasResultNeighCellsR16Seq.IsComponentPresent(0) {
				c.MeasResultNeighCellListNR = new(MeasResultListLogging2NR_r16)
				if err := (*c.MeasResultNeighCellListNR).Decode(d); err != nil {
					return err
				}
			}
			if logMeasInfoR16MeasResultNeighCellsR16Seq.IsComponentPresent(1) {
				c.MeasResultNeighCellListEUTRA = new(MeasResultList2EUTRA_r16)
				if err := (*c.MeasResultNeighCellListEUTRA).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. anyCellSelectionDetected-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(logMeasInfoR16AnyCellSelectionDetectedR16Constraints)
			if err != nil {
				return err
			}
			ie.AnyCellSelectionDetected_r16 = &idx
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
				{Name: "inDeviceCoexDetected-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(logMeasInfoR16ExtInDeviceCoexDetectedR17Constraints)
			if err != nil {
				return err
			}
			ie.InDeviceCoexDetected_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nsag-ID-r19", Optional: true},
				{Name: "reselectedCellId-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Nsag_ID_r19 = new(NSAG_ID_r17)
			if err := ie.Nsag_ID_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ReselectedCellId_r19 = new(CGI_Info_Logging_r16)
			if err := ie.ReselectedCellId_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
