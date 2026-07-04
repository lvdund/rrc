// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GeneralParametersMRDC-XDD-Diff (line 25626).

var generalParametersMRDCXDDDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "splitSRB-WithOneUL-Path", Optional: true},
		{Name: "splitDRB-withUL-Both-MCG-SCG", Optional: true},
		{Name: "srb3", Optional: true},
		{Name: "dummy", Optional: true},
	},
}

const (
	GeneralParametersMRDC_XDD_Diff_SplitSRB_WithOneUL_Path_Supported = 0
)

var generalParametersMRDCXDDDiffSplitSRBWithOneULPathConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GeneralParametersMRDC_XDD_Diff_SplitSRB_WithOneUL_Path_Supported},
}

const (
	GeneralParametersMRDC_XDD_Diff_SplitDRB_WithUL_Both_MCG_SCG_Supported = 0
)

var generalParametersMRDCXDDDiffSplitDRBWithULBothMCGSCGConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GeneralParametersMRDC_XDD_Diff_SplitDRB_WithUL_Both_MCG_SCG_Supported},
}

const (
	GeneralParametersMRDC_XDD_Diff_Srb3_Supported = 0
)

var generalParametersMRDCXDDDiffSrb3Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GeneralParametersMRDC_XDD_Diff_Srb3_Supported},
}

const (
	GeneralParametersMRDC_XDD_Diff_Dummy_Supported = 0
)

var generalParametersMRDCXDDDiffDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GeneralParametersMRDC_XDD_Diff_Dummy_Supported},
}

type GeneralParametersMRDC_XDD_Diff struct {
	SplitSRB_WithOneUL_Path      *int64
	SplitDRB_WithUL_Both_MCG_SCG *int64
	Srb3                         *int64
	Dummy                        *int64
}

func (ie *GeneralParametersMRDC_XDD_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(generalParametersMRDCXDDDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SplitSRB_WithOneUL_Path != nil, ie.SplitDRB_WithUL_Both_MCG_SCG != nil, ie.Srb3 != nil, ie.Dummy != nil}); err != nil {
		return err
	}

	// 3. splitSRB-WithOneUL-Path: enumerated
	{
		if ie.SplitSRB_WithOneUL_Path != nil {
			if err := e.EncodeEnumerated(*ie.SplitSRB_WithOneUL_Path, generalParametersMRDCXDDDiffSplitSRBWithOneULPathConstraints); err != nil {
				return err
			}
		}
	}

	// 4. splitDRB-withUL-Both-MCG-SCG: enumerated
	{
		if ie.SplitDRB_WithUL_Both_MCG_SCG != nil {
			if err := e.EncodeEnumerated(*ie.SplitDRB_WithUL_Both_MCG_SCG, generalParametersMRDCXDDDiffSplitDRBWithULBothMCGSCGConstraints); err != nil {
				return err
			}
		}
	}

	// 5. srb3: enumerated
	{
		if ie.Srb3 != nil {
			if err := e.EncodeEnumerated(*ie.Srb3, generalParametersMRDCXDDDiffSrb3Constraints); err != nil {
				return err
			}
		}
	}

	// 6. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, generalParametersMRDCXDDDiffDummyConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *GeneralParametersMRDC_XDD_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(generalParametersMRDCXDDDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. splitSRB-WithOneUL-Path: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(generalParametersMRDCXDDDiffSplitSRBWithOneULPathConstraints)
			if err != nil {
				return err
			}
			ie.SplitSRB_WithOneUL_Path = &idx
		}
	}

	// 4. splitDRB-withUL-Both-MCG-SCG: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(generalParametersMRDCXDDDiffSplitDRBWithULBothMCGSCGConstraints)
			if err != nil {
				return err
			}
			ie.SplitDRB_WithUL_Both_MCG_SCG = &idx
		}
	}

	// 5. srb3: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(generalParametersMRDCXDDDiffSrb3Constraints)
			if err != nil {
				return err
			}
			ie.Srb3 = &idx
		}
	}

	// 6. dummy: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(generalParametersMRDCXDDDiffDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	return nil
}
