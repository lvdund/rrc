// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureReportSCG-EUTRA (line 1907).

var failureReportSCGEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "failureType"},
		{Name: "measResultFreqListMRDC", Optional: true},
		{Name: "measResultSCG-FailureMRDC", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	FailureReportSCG_EUTRA_FailureType_T313_Expiry         = 0
	FailureReportSCG_EUTRA_FailureType_RandomAccessProblem = 1
	FailureReportSCG_EUTRA_FailureType_Rlc_MaxNumRetx      = 2
	FailureReportSCG_EUTRA_FailureType_Scg_ChangeFailure   = 3
	FailureReportSCG_EUTRA_FailureType_Spare4              = 4
	FailureReportSCG_EUTRA_FailureType_Spare3              = 5
	FailureReportSCG_EUTRA_FailureType_Spare2              = 6
	FailureReportSCG_EUTRA_FailureType_Spare1              = 7
)

var failureReportSCGEUTRAFailureTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FailureReportSCG_EUTRA_FailureType_T313_Expiry, FailureReportSCG_EUTRA_FailureType_RandomAccessProblem, FailureReportSCG_EUTRA_FailureType_Rlc_MaxNumRetx, FailureReportSCG_EUTRA_FailureType_Scg_ChangeFailure, FailureReportSCG_EUTRA_FailureType_Spare4, FailureReportSCG_EUTRA_FailureType_Spare3, FailureReportSCG_EUTRA_FailureType_Spare2, FailureReportSCG_EUTRA_FailureType_Spare1},
}

var failureReportSCGEUTRAMeasResultSCGFailureMRDCConstraints = per.SizeConstraints{}

type FailureReportSCG_EUTRA struct {
	FailureType               int64
	MeasResultFreqListMRDC    *MeasResultFreqListFailMRDC
	MeasResultSCG_FailureMRDC []byte
	LocationInfo_r16          *LocationInfo_r16
}

func (ie *FailureReportSCG_EUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureReportSCGEUTRAConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.LocationInfo_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultFreqListMRDC != nil, ie.MeasResultSCG_FailureMRDC != nil}); err != nil {
		return err
	}

	// 3. failureType: enumerated
	{
		if err := e.EncodeEnumerated(ie.FailureType, failureReportSCGEUTRAFailureTypeConstraints); err != nil {
			return err
		}
	}

	// 4. measResultFreqListMRDC: ref
	{
		if ie.MeasResultFreqListMRDC != nil {
			if err := ie.MeasResultFreqListMRDC.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. measResultSCG-FailureMRDC: octet-string
	{
		if ie.MeasResultSCG_FailureMRDC != nil {
			if err := e.EncodeOctetString(ie.MeasResultSCG_FailureMRDC, failureReportSCGEUTRAMeasResultSCGFailureMRDCConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "locationInfo-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LocationInfo_r16 != nil}); err != nil {
				return err
			}

			if ie.LocationInfo_r16 != nil {
				if err := ie.LocationInfo_r16.Encode(ex); err != nil {
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

func (ie *FailureReportSCG_EUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureReportSCGEUTRAConstraints)

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
		v0, err := d.DecodeEnumerated(failureReportSCGEUTRAFailureTypeConstraints)
		if err != nil {
			return err
		}
		ie.FailureType = v0
	}

	// 4. measResultFreqListMRDC: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasResultFreqListMRDC = new(MeasResultFreqListFailMRDC)
			if err := ie.MeasResultFreqListMRDC.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. measResultSCG-FailureMRDC: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(failureReportSCGEUTRAMeasResultSCGFailureMRDCConstraints)
			if err != nil {
				return err
			}
			ie.MeasResultSCG_FailureMRDC = v
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
	}

	return nil
}
