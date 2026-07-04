// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FailureReportMCG-r16 (line 705).

var failureReportMCGR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "failureType-r16", Optional: true},
		{Name: "measResultFreqList-r16", Optional: true},
		{Name: "measResultFreqListEUTRA-r16", Optional: true},
		{Name: "measResultSCG-r16", Optional: true},
		{Name: "measResultSCG-EUTRA-r16", Optional: true},
		{Name: "measResultFreqListUTRA-FDD-r16", Optional: true},
	},
}

const (
	FailureReportMCG_r16_FailureType_r16_T310_Expiry                    = 0
	FailureReportMCG_r16_FailureType_r16_RandomAccessProblem            = 1
	FailureReportMCG_r16_FailureType_r16_Rlc_MaxNumRetx                 = 2
	FailureReportMCG_r16_FailureType_r16_T312_Expiry_r16                = 3
	FailureReportMCG_r16_FailureType_r16_Lbt_Failure_r16                = 4
	FailureReportMCG_r16_FailureType_r16_BeamFailureRecoveryFailure_r16 = 5
	FailureReportMCG_r16_FailureType_r16_Bh_RLF_r16                     = 6
	FailureReportMCG_r16_FailureType_r16_Spare1                         = 7
)

var failureReportMCGR16FailureTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FailureReportMCG_r16_FailureType_r16_T310_Expiry, FailureReportMCG_r16_FailureType_r16_RandomAccessProblem, FailureReportMCG_r16_FailureType_r16_Rlc_MaxNumRetx, FailureReportMCG_r16_FailureType_r16_T312_Expiry_r16, FailureReportMCG_r16_FailureType_r16_Lbt_Failure_r16, FailureReportMCG_r16_FailureType_r16_BeamFailureRecoveryFailure_r16, FailureReportMCG_r16_FailureType_r16_Bh_RLF_r16, FailureReportMCG_r16_FailureType_r16_Spare1},
}

var failureReportMCGR16MeasResultSCGR16Constraints = per.SizeConstraints{}

var failureReportMCGR16MeasResultSCGEUTRAR16Constraints = per.SizeConstraints{}

type FailureReportMCG_r16 struct {
	FailureType_r16                *int64
	MeasResultFreqList_r16         *MeasResultList2NR
	MeasResultFreqListEUTRA_r16    *MeasResultList2EUTRA
	MeasResultSCG_r16              []byte
	MeasResultSCG_EUTRA_r16        []byte
	MeasResultFreqListUTRA_FDD_r16 *MeasResultList2UTRA
}

func (ie *FailureReportMCG_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(failureReportMCGR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FailureType_r16 != nil, ie.MeasResultFreqList_r16 != nil, ie.MeasResultFreqListEUTRA_r16 != nil, ie.MeasResultSCG_r16 != nil, ie.MeasResultSCG_EUTRA_r16 != nil, ie.MeasResultFreqListUTRA_FDD_r16 != nil}); err != nil {
		return err
	}

	// 3. failureType-r16: enumerated
	{
		if ie.FailureType_r16 != nil {
			if err := e.EncodeEnumerated(*ie.FailureType_r16, failureReportMCGR16FailureTypeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. measResultFreqList-r16: ref
	{
		if ie.MeasResultFreqList_r16 != nil {
			if err := ie.MeasResultFreqList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. measResultFreqListEUTRA-r16: ref
	{
		if ie.MeasResultFreqListEUTRA_r16 != nil {
			if err := ie.MeasResultFreqListEUTRA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. measResultSCG-r16: octet-string
	{
		if ie.MeasResultSCG_r16 != nil {
			if err := e.EncodeOctetString(ie.MeasResultSCG_r16, failureReportMCGR16MeasResultSCGR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. measResultSCG-EUTRA-r16: octet-string
	{
		if ie.MeasResultSCG_EUTRA_r16 != nil {
			if err := e.EncodeOctetString(ie.MeasResultSCG_EUTRA_r16, failureReportMCGR16MeasResultSCGEUTRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. measResultFreqListUTRA-FDD-r16: ref
	{
		if ie.MeasResultFreqListUTRA_FDD_r16 != nil {
			if err := ie.MeasResultFreqListUTRA_FDD_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FailureReportMCG_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(failureReportMCGR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. failureType-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(failureReportMCGR16FailureTypeR16Constraints)
			if err != nil {
				return err
			}
			ie.FailureType_r16 = &idx
		}
	}

	// 4. measResultFreqList-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.MeasResultFreqList_r16 = new(MeasResultList2NR)
			if err := ie.MeasResultFreqList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. measResultFreqListEUTRA-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultFreqListEUTRA_r16 = new(MeasResultList2EUTRA)
			if err := ie.MeasResultFreqListEUTRA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. measResultSCG-r16: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(failureReportMCGR16MeasResultSCGR16Constraints)
			if err != nil {
				return err
			}
			ie.MeasResultSCG_r16 = v
		}
	}

	// 7. measResultSCG-EUTRA-r16: octet-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeOctetString(failureReportMCGR16MeasResultSCGEUTRAR16Constraints)
			if err != nil {
				return err
			}
			ie.MeasResultSCG_EUTRA_r16 = v
		}
	}

	// 8. measResultFreqListUTRA-FDD-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.MeasResultFreqListUTRA_FDD_r16 = new(MeasResultList2UTRA)
			if err := ie.MeasResultFreqListUTRA_FDD_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
