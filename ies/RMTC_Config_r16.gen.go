// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RMTC-Config-r16 (line 9561).

var rMTCConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rmtc-Periodicity-r16"},
		{Name: "rmtc-SubframeOffset-r16", Optional: true},
		{Name: "measDurationSymbols-r16"},
		{Name: "rmtc-Frequency-r16"},
		{Name: "ref-SCS-CP-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	RMTC_Config_r16_Rmtc_Periodicity_r16_Ms40  = 0
	RMTC_Config_r16_Rmtc_Periodicity_r16_Ms80  = 1
	RMTC_Config_r16_Rmtc_Periodicity_r16_Ms160 = 2
	RMTC_Config_r16_Rmtc_Periodicity_r16_Ms320 = 3
	RMTC_Config_r16_Rmtc_Periodicity_r16_Ms640 = 4
)

var rMTCConfigR16RmtcPeriodicityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RMTC_Config_r16_Rmtc_Periodicity_r16_Ms40, RMTC_Config_r16_Rmtc_Periodicity_r16_Ms80, RMTC_Config_r16_Rmtc_Periodicity_r16_Ms160, RMTC_Config_r16_Rmtc_Periodicity_r16_Ms320, RMTC_Config_r16_Rmtc_Periodicity_r16_Ms640},
}

var rMTCConfigR16RmtcSubframeOffsetR16Constraints = per.Constrained(0, 639)

const (
	RMTC_Config_r16_MeasDurationSymbols_r16_Sym1      = 0
	RMTC_Config_r16_MeasDurationSymbols_r16_Sym14or12 = 1
	RMTC_Config_r16_MeasDurationSymbols_r16_Sym28or24 = 2
	RMTC_Config_r16_MeasDurationSymbols_r16_Sym42or36 = 3
	RMTC_Config_r16_MeasDurationSymbols_r16_Sym70or60 = 4
)

var rMTCConfigR16MeasDurationSymbolsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RMTC_Config_r16_MeasDurationSymbols_r16_Sym1, RMTC_Config_r16_MeasDurationSymbols_r16_Sym14or12, RMTC_Config_r16_MeasDurationSymbols_r16_Sym28or24, RMTC_Config_r16_MeasDurationSymbols_r16_Sym42or36, RMTC_Config_r16_MeasDurationSymbols_r16_Sym70or60},
}

const (
	RMTC_Config_r16_Ref_SCS_CP_r16_KHz15     = 0
	RMTC_Config_r16_Ref_SCS_CP_r16_KHz30     = 1
	RMTC_Config_r16_Ref_SCS_CP_r16_KHz60_NCP = 2
	RMTC_Config_r16_Ref_SCS_CP_r16_KHz60_ECP = 3
)

var rMTCConfigR16RefSCSCPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RMTC_Config_r16_Ref_SCS_CP_r16_KHz15, RMTC_Config_r16_Ref_SCS_CP_r16_KHz30, RMTC_Config_r16_Ref_SCS_CP_r16_KHz60_NCP, RMTC_Config_r16_Ref_SCS_CP_r16_KHz60_ECP},
}

const (
	RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz100  = 0
	RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz400  = 1
	RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz800  = 2
	RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz1600 = 3
	RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz2000 = 4
)

var rMTCConfigR16ExtRmtcBandwidthR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz100, RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz400, RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz800, RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz1600, RMTC_Config_r16_Ext_Rmtc_Bandwidth_r17_Mhz2000},
}

const (
	RMTC_Config_r16_Ext_MeasDurationSymbols_v1700_Sym140  = 0
	RMTC_Config_r16_Ext_MeasDurationSymbols_v1700_Sym560  = 1
	RMTC_Config_r16_Ext_MeasDurationSymbols_v1700_Sym1120 = 2
)

var rMTCConfigR16ExtMeasDurationSymbolsV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RMTC_Config_r16_Ext_MeasDurationSymbols_v1700_Sym140, RMTC_Config_r16_Ext_MeasDurationSymbols_v1700_Sym560, RMTC_Config_r16_Ext_MeasDurationSymbols_v1700_Sym1120},
}

const (
	RMTC_Config_r16_Ext_Ref_SCS_CP_v1700_KHz120 = 0
	RMTC_Config_r16_Ext_Ref_SCS_CP_v1700_KHz480 = 1
	RMTC_Config_r16_Ext_Ref_SCS_CP_v1700_KHz960 = 2
)

var rMTCConfigR16ExtRefSCSCPV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RMTC_Config_r16_Ext_Ref_SCS_CP_v1700_KHz120, RMTC_Config_r16_Ext_Ref_SCS_CP_v1700_KHz480, RMTC_Config_r16_Ext_Ref_SCS_CP_v1700_KHz960},
}

var rMTCConfigR16ExtTciStateInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tci-StateId-r17"},
		{Name: "ref-ServCellId-r17", Optional: true},
	},
}

type RMTC_Config_r16 struct {
	Rmtc_Periodicity_r16      int64
	Rmtc_SubframeOffset_r16   *int64
	MeasDurationSymbols_r16   int64
	Rmtc_Frequency_r16        ARFCN_ValueNR
	Ref_SCS_CP_r16            int64
	Rmtc_Bandwidth_r17        *int64
	MeasDurationSymbols_v1700 *int64
	Ref_SCS_CP_v1700          *int64
	Tci_StateInfo_r17         *struct {
		Tci_StateId_r17    TCI_StateId
		Ref_ServCellId_r17 *ServCellIndex
	}
	Ref_BWPId_r17 *BWP_Id
}

func (ie *RMTC_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rMTCConfigR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Rmtc_Bandwidth_r17 != nil || ie.MeasDurationSymbols_v1700 != nil || ie.Ref_SCS_CP_v1700 != nil || ie.Tci_StateInfo_r17 != nil
	hasExtGroup1 := ie.Ref_BWPId_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rmtc_SubframeOffset_r16 != nil}); err != nil {
		return err
	}

	// 3. rmtc-Periodicity-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Rmtc_Periodicity_r16, rMTCConfigR16RmtcPeriodicityR16Constraints); err != nil {
			return err
		}
	}

	// 4. rmtc-SubframeOffset-r16: integer
	{
		if ie.Rmtc_SubframeOffset_r16 != nil {
			if err := e.EncodeInteger(*ie.Rmtc_SubframeOffset_r16, rMTCConfigR16RmtcSubframeOffsetR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. measDurationSymbols-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MeasDurationSymbols_r16, rMTCConfigR16MeasDurationSymbolsR16Constraints); err != nil {
			return err
		}
	}

	// 6. rmtc-Frequency-r16: ref
	{
		if err := ie.Rmtc_Frequency_r16.Encode(e); err != nil {
			return err
		}
	}

	// 7. ref-SCS-CP-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ref_SCS_CP_r16, rMTCConfigR16RefSCSCPR16Constraints); err != nil {
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
					{Name: "rmtc-Bandwidth-r17", Optional: true},
					{Name: "measDurationSymbols-v1700", Optional: true},
					{Name: "ref-SCS-CP-v1700", Optional: true},
					{Name: "tci-StateInfo-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rmtc_Bandwidth_r17 != nil, ie.MeasDurationSymbols_v1700 != nil, ie.Ref_SCS_CP_v1700 != nil, ie.Tci_StateInfo_r17 != nil}); err != nil {
				return err
			}

			if ie.Rmtc_Bandwidth_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rmtc_Bandwidth_r17, rMTCConfigR16ExtRmtcBandwidthR17Constraints); err != nil {
					return err
				}
			}

			if ie.MeasDurationSymbols_v1700 != nil {
				if err := ex.EncodeEnumerated(*ie.MeasDurationSymbols_v1700, rMTCConfigR16ExtMeasDurationSymbolsV1700Constraints); err != nil {
					return err
				}
			}

			if ie.Ref_SCS_CP_v1700 != nil {
				if err := ex.EncodeEnumerated(*ie.Ref_SCS_CP_v1700, rMTCConfigR16ExtRefSCSCPV1700Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_StateInfo_r17 != nil {
				c := ie.Tci_StateInfo_r17
				rMTCConfigR16ExtTciStateInfoR17Seq := ex.NewSequenceEncoder(rMTCConfigR16ExtTciStateInfoR17Constraints)
				if err := rMTCConfigR16ExtTciStateInfoR17Seq.EncodePreamble([]bool{c.Ref_ServCellId_r17 != nil}); err != nil {
					return err
				}
				if err := c.Tci_StateId_r17.Encode(ex); err != nil {
					return err
				}
				if c.Ref_ServCellId_r17 != nil {
					if err := c.Ref_ServCellId_r17.Encode(ex); err != nil {
						return err
					}
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
					{Name: "ref-BWPId-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ref_BWPId_r17 != nil}); err != nil {
				return err
			}

			if ie.Ref_BWPId_r17 != nil {
				if err := ie.Ref_BWPId_r17.Encode(ex); err != nil {
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

func (ie *RMTC_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rMTCConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rmtc-Periodicity-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(rMTCConfigR16RmtcPeriodicityR16Constraints)
		if err != nil {
			return err
		}
		ie.Rmtc_Periodicity_r16 = v0
	}

	// 4. rmtc-SubframeOffset-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(rMTCConfigR16RmtcSubframeOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Rmtc_SubframeOffset_r16 = &v
		}
	}

	// 5. measDurationSymbols-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(rMTCConfigR16MeasDurationSymbolsR16Constraints)
		if err != nil {
			return err
		}
		ie.MeasDurationSymbols_r16 = v2
	}

	// 6. rmtc-Frequency-r16: ref
	{
		if err := ie.Rmtc_Frequency_r16.Decode(d); err != nil {
			return err
		}
	}

	// 7. ref-SCS-CP-r16: enumerated
	{
		v4, err := d.DecodeEnumerated(rMTCConfigR16RefSCSCPR16Constraints)
		if err != nil {
			return err
		}
		ie.Ref_SCS_CP_r16 = v4
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
				{Name: "rmtc-Bandwidth-r17", Optional: true},
				{Name: "measDurationSymbols-v1700", Optional: true},
				{Name: "ref-SCS-CP-v1700", Optional: true},
				{Name: "tci-StateInfo-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rMTCConfigR16ExtRmtcBandwidthR17Constraints)
			if err != nil {
				return err
			}
			ie.Rmtc_Bandwidth_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(rMTCConfigR16ExtMeasDurationSymbolsV1700Constraints)
			if err != nil {
				return err
			}
			ie.MeasDurationSymbols_v1700 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(rMTCConfigR16ExtRefSCSCPV1700Constraints)
			if err != nil {
				return err
			}
			ie.Ref_SCS_CP_v1700 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Tci_StateInfo_r17 = &struct {
				Tci_StateId_r17    TCI_StateId
				Ref_ServCellId_r17 *ServCellIndex
			}{}
			c := ie.Tci_StateInfo_r17
			rMTCConfigR16ExtTciStateInfoR17Seq := dx.NewSequenceDecoder(rMTCConfigR16ExtTciStateInfoR17Constraints)
			if err := rMTCConfigR16ExtTciStateInfoR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.Tci_StateId_r17.Decode(dx); err != nil {
					return err
				}
			}
			if rMTCConfigR16ExtTciStateInfoR17Seq.IsComponentPresent(1) {
				c.Ref_ServCellId_r17 = new(ServCellIndex)
				if err := (*c.Ref_ServCellId_r17).Decode(dx); err != nil {
					return err
				}
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
				{Name: "ref-BWPId-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ref_BWPId_r17 = new(BWP_Id)
			if err := ie.Ref_BWPId_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
