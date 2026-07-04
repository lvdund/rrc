// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureReportSCG (line 1848).

var failureReportSCGConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "failureType"},
		{Name: "measResultFreqList", Optional: true},
		{Name: "measResultSCG-Failure", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	FailureReportSCG_FailureType_T310_Expiry             = 0
	FailureReportSCG_FailureType_RandomAccessProblem     = 1
	FailureReportSCG_FailureType_Rlc_MaxNumRetx          = 2
	FailureReportSCG_FailureType_SynchReconfigFailureSCG = 3
	FailureReportSCG_FailureType_Scg_ReconfigFailure     = 4
	FailureReportSCG_FailureType_Srb3_IntegrityFailure   = 5
	FailureReportSCG_FailureType_Other_r16               = 6
	FailureReportSCG_FailureType_Spare1                  = 7
)

var failureReportSCGFailureTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FailureReportSCG_FailureType_T310_Expiry, FailureReportSCG_FailureType_RandomAccessProblem, FailureReportSCG_FailureType_Rlc_MaxNumRetx, FailureReportSCG_FailureType_SynchReconfigFailureSCG, FailureReportSCG_FailureType_Scg_ReconfigFailure, FailureReportSCG_FailureType_Srb3_IntegrityFailure, FailureReportSCG_FailureType_Other_r16, FailureReportSCG_FailureType_Spare1},
}

var failureReportSCGMeasResultSCGFailureConstraints = per.SizeConstraints{}

const (
	FailureReportSCG_Ext_FailureType_v1610_Scg_LbtFailure_r16             = 0
	FailureReportSCG_Ext_FailureType_v1610_BeamFailureRecoveryFailure_r16 = 1
	FailureReportSCG_Ext_FailureType_v1610_T312_Expiry_r16                = 2
	FailureReportSCG_Ext_FailureType_v1610_Bh_RLF_r16                     = 3
	FailureReportSCG_Ext_FailureType_v1610_BeamFailure_r17                = 4
	FailureReportSCG_Ext_FailureType_v1610_Spare3                         = 5
	FailureReportSCG_Ext_FailureType_v1610_Spare2                         = 6
	FailureReportSCG_Ext_FailureType_v1610_Spare1                         = 7
)

var failureReportSCGExtFailureTypeV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FailureReportSCG_Ext_FailureType_v1610_Scg_LbtFailure_r16, FailureReportSCG_Ext_FailureType_v1610_BeamFailureRecoveryFailure_r16, FailureReportSCG_Ext_FailureType_v1610_T312_Expiry_r16, FailureReportSCG_Ext_FailureType_v1610_Bh_RLF_r16, FailureReportSCG_Ext_FailureType_v1610_BeamFailure_r17, FailureReportSCG_Ext_FailureType_v1610_Spare3, FailureReportSCG_Ext_FailureType_v1610_Spare2, FailureReportSCG_Ext_FailureType_v1610_Spare1},
}

var failureReportSCGTimeSCGFailureR17Constraints = per.Constrained(0, 1023)

type FailureReportSCG struct {
	FailureType           int64
	MeasResultFreqList    *MeasResultFreqList
	MeasResultSCG_Failure []byte
	LocationInfo_r16      *LocationInfo_r16
	FailureType_v1610     *int64
	PreviousPSCellId_r17  *struct {
		PhysCellId_r17  PhysCellId
		CarrierFreq_r17 ARFCN_ValueNR
	}
	FailedPSCellId_r17 *struct {
		PhysCellId_r17  PhysCellId
		CarrierFreq_r17 ARFCN_ValueNR
	}
	TimeSCGFailure_r17               *int64
	PerRAInfoList_r17                *PerRAInfoList_r16
	PerRAInfoList_V17b0              *PerRAInfoList_v1660
	PerRAInfoList_v1840              *PerRAInfoList_v1800
	Cho_WithCandidateSCGInfoList_r19 *CHO_WithCandidateSCGInfoList_r19
}

func (ie *FailureReportSCG) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureReportSCGConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.LocationInfo_r16 != nil || ie.FailureType_v1610 != nil
	hasExtGroup1 := ie.PreviousPSCellId_r17 != nil || ie.FailedPSCellId_r17 != nil || ie.TimeSCGFailure_r17 != nil || ie.PerRAInfoList_r17 != nil
	hasExtGroup2 := ie.PerRAInfoList_V17b0 != nil
	hasExtGroup3 := ie.PerRAInfoList_v1840 != nil
	hasExtGroup4 := ie.Cho_WithCandidateSCGInfoList_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultFreqList != nil, ie.MeasResultSCG_Failure != nil}); err != nil {
		return err
	}

	// 3. failureType: enumerated
	{
		if err := e.EncodeEnumerated(ie.FailureType, failureReportSCGFailureTypeConstraints); err != nil {
			return err
		}
	}

	// 4. measResultFreqList: ref
	{
		if ie.MeasResultFreqList != nil {
			if err := ie.MeasResultFreqList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. measResultSCG-Failure: octet-string
	{
		if ie.MeasResultSCG_Failure != nil {
			if err := e.EncodeOctetString(ie.MeasResultSCG_Failure, failureReportSCGMeasResultSCGFailureConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "locationInfo-r16", Optional: true},
					{Name: "failureType-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LocationInfo_r16 != nil, ie.FailureType_v1610 != nil}); err != nil {
				return err
			}

			if ie.LocationInfo_r16 != nil {
				if err := ie.LocationInfo_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.FailureType_v1610 != nil {
				if err := ex.EncodeEnumerated(*ie.FailureType_v1610, failureReportSCGExtFailureTypeV1610Constraints); err != nil {
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
					{Name: "previousPSCellId-r17", Optional: true},
					{Name: "failedPSCellId-r17", Optional: true},
					{Name: "timeSCGFailure-r17", Optional: true},
					{Name: "perRAInfoList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PreviousPSCellId_r17 != nil, ie.FailedPSCellId_r17 != nil, ie.TimeSCGFailure_r17 != nil, ie.PerRAInfoList_r17 != nil}); err != nil {
				return err
			}

			if ie.PreviousPSCellId_r17 != nil {
				c := ie.PreviousPSCellId_r17
				if err := c.PhysCellId_r17.Encode(ex); err != nil {
					return err
				}
				if err := c.CarrierFreq_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.FailedPSCellId_r17 != nil {
				c := ie.FailedPSCellId_r17
				if err := c.PhysCellId_r17.Encode(ex); err != nil {
					return err
				}
				if err := c.CarrierFreq_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.TimeSCGFailure_r17 != nil {
				if err := ex.EncodeInteger(*ie.TimeSCGFailure_r17, failureReportSCGTimeSCGFailureR17Constraints); err != nil {
					return err
				}
			}

			if ie.PerRAInfoList_r17 != nil {
				if err := ie.PerRAInfoList_r17.Encode(ex); err != nil {
					return err
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
					{Name: "perRAInfoList-v17b0", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PerRAInfoList_V17b0 != nil}); err != nil {
				return err
			}

			if ie.PerRAInfoList_V17b0 != nil {
				if err := ie.PerRAInfoList_V17b0.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "perRAInfoList-v1840", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PerRAInfoList_v1840 != nil}); err != nil {
				return err
			}

			if ie.PerRAInfoList_v1840 != nil {
				if err := ie.PerRAInfoList_v1840.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "cho-WithCandidateSCGInfoList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cho_WithCandidateSCGInfoList_r19 != nil}); err != nil {
				return err
			}

			if ie.Cho_WithCandidateSCGInfoList_r19 != nil {
				if err := ie.Cho_WithCandidateSCGInfoList_r19.Encode(ex); err != nil {
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

func (ie *FailureReportSCG) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureReportSCGConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. failureType: enumerated
	{
		v0, err := d.DecodeEnumerated(failureReportSCGFailureTypeConstraints)
		if err != nil {
			return err
		}
		ie.FailureType = v0
	}

	// 4. measResultFreqList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasResultFreqList = new(MeasResultFreqList)
			if err := ie.MeasResultFreqList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. measResultSCG-Failure: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(failureReportSCGMeasResultSCGFailureConstraints)
			if err != nil {
				return err
			}
			ie.MeasResultSCG_Failure = v
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
				{Name: "locationInfo-r16", Optional: true},
				{Name: "failureType-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.LocationInfo_r16 = new(LocationInfo_r16)
			if err := ie.LocationInfo_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(failureReportSCGExtFailureTypeV1610Constraints)
			if err != nil {
				return err
			}
			ie.FailureType_v1610 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "previousPSCellId-r17", Optional: true},
				{Name: "failedPSCellId-r17", Optional: true},
				{Name: "timeSCGFailure-r17", Optional: true},
				{Name: "perRAInfoList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PreviousPSCellId_r17 = &struct {
				PhysCellId_r17  PhysCellId
				CarrierFreq_r17 ARFCN_ValueNR
			}{}
			c := ie.PreviousPSCellId_r17
			{
				if err := c.PhysCellId_r17.Decode(dx); err != nil {
					return err
				}
			}
			{
				if err := c.CarrierFreq_r17.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.FailedPSCellId_r17 = &struct {
				PhysCellId_r17  PhysCellId
				CarrierFreq_r17 ARFCN_ValueNR
			}{}
			c := ie.FailedPSCellId_r17
			{
				if err := c.PhysCellId_r17.Decode(dx); err != nil {
					return err
				}
			}
			{
				if err := c.CarrierFreq_r17.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(failureReportSCGTimeSCGFailureR17Constraints)
			if err != nil {
				return err
			}
			ie.TimeSCGFailure_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.PerRAInfoList_r17 = new(PerRAInfoList_r16)
			if err := ie.PerRAInfoList_r17.Decode(dx); err != nil {
				return err
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
				{Name: "perRAInfoList-v17b0", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PerRAInfoList_V17b0 = new(PerRAInfoList_v1660)
			if err := ie.PerRAInfoList_V17b0.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "perRAInfoList-v1840", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PerRAInfoList_v1840 = new(PerRAInfoList_v1800)
			if err := ie.PerRAInfoList_v1840.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cho-WithCandidateSCGInfoList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Cho_WithCandidateSCGInfoList_r19 = new(CHO_WithCandidateSCGInfoList_r19)
			if err := ie.Cho_WithCandidateSCGInfoList_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
