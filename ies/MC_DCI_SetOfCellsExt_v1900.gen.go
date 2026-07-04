// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MC-DCI-SetOfCellsExt-v1900 (line 14867).

var mCDCISetOfCellsExtV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tdra-FieldIndexListDCI-1-3-r19", Optional: true},
		{Name: "tdra-FieldIndexListDCI-0-3-r19", Optional: true},
	},
}

var mCDCISetOfCellsExtV1900TdraFieldIndexListDCI13R19Constraints = per.SizeRange(1, 64)

var mCDCISetOfCellsExtV1900TdraFieldIndexListDCI03R19Constraints = per.SizeRange(1, 128)

type MC_DCI_SetOfCellsExt_v1900 struct {
	Tdra_FieldIndexListDCI_1_3_r19 []TDRA_FieldIndexDCI_1_3_r19
	Tdra_FieldIndexListDCI_0_3_r19 []TDRA_FieldIndexDCI_0_3_r19
}

func (ie *MC_DCI_SetOfCellsExt_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mCDCISetOfCellsExtV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tdra_FieldIndexListDCI_1_3_r19 != nil, ie.Tdra_FieldIndexListDCI_0_3_r19 != nil}); err != nil {
		return err
	}

	// 2. tdra-FieldIndexListDCI-1-3-r19: sequence-of
	{
		if ie.Tdra_FieldIndexListDCI_1_3_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsExtV1900TdraFieldIndexListDCI13R19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tdra_FieldIndexListDCI_1_3_r19))); err != nil {
				return err
			}
			for i := range ie.Tdra_FieldIndexListDCI_1_3_r19 {
				if err := ie.Tdra_FieldIndexListDCI_1_3_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. tdra-FieldIndexListDCI-0-3-r19: sequence-of
	{
		if ie.Tdra_FieldIndexListDCI_0_3_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(mCDCISetOfCellsExtV1900TdraFieldIndexListDCI03R19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tdra_FieldIndexListDCI_0_3_r19))); err != nil {
				return err
			}
			for i := range ie.Tdra_FieldIndexListDCI_0_3_r19 {
				if err := ie.Tdra_FieldIndexListDCI_0_3_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MC_DCI_SetOfCellsExt_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mCDCISetOfCellsExtV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. tdra-FieldIndexListDCI-1-3-r19: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsExtV1900TdraFieldIndexListDCI13R19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tdra_FieldIndexListDCI_1_3_r19 = make([]TDRA_FieldIndexDCI_1_3_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tdra_FieldIndexListDCI_1_3_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. tdra-FieldIndexListDCI-0-3-r19: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(mCDCISetOfCellsExtV1900TdraFieldIndexListDCI03R19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tdra_FieldIndexListDCI_0_3_r19 = make([]TDRA_FieldIndexDCI_0_3_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tdra_FieldIndexListDCI_0_3_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
