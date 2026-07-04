// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CodebookParametersHARQ-ACK-PUSCH-r18 (line 19027).

var codebookParametersHARQACKPUSCHR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "multiplexingType1-r18", Optional: true},
		{Name: "multiplexingType2-r18", Optional: true},
		{Name: "multiplexingType3-r18", Optional: true},
		{Name: "pucch-DiffResource-PDSCH-r18", Optional: true},
		{Name: "diffCB-Size-PDSCH-r18", Optional: true},
	},
}

const (
	CodebookParametersHARQ_ACK_PUSCH_r18_MultiplexingType1_r18_Supported = 0
)

var codebookParametersHARQACKPUSCHR18MultiplexingType1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersHARQ_ACK_PUSCH_r18_MultiplexingType1_r18_Supported},
}

const (
	CodebookParametersHARQ_ACK_PUSCH_r18_MultiplexingType2_r18_Supported = 0
)

var codebookParametersHARQACKPUSCHR18MultiplexingType2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersHARQ_ACK_PUSCH_r18_MultiplexingType2_r18_Supported},
}

const (
	CodebookParametersHARQ_ACK_PUSCH_r18_MultiplexingType3_r18_Supported = 0
)

var codebookParametersHARQACKPUSCHR18MultiplexingType3R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersHARQ_ACK_PUSCH_r18_MultiplexingType3_r18_Supported},
}

const (
	CodebookParametersHARQ_ACK_PUSCH_r18_Pucch_DiffResource_PDSCH_r18_Supported = 0
)

var codebookParametersHARQACKPUSCHR18PucchDiffResourcePDSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersHARQ_ACK_PUSCH_r18_Pucch_DiffResource_PDSCH_r18_Supported},
}

const (
	CodebookParametersHARQ_ACK_PUSCH_r18_DiffCB_Size_PDSCH_r18_Supported = 0
)

var codebookParametersHARQACKPUSCHR18DiffCBSizePDSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CodebookParametersHARQ_ACK_PUSCH_r18_DiffCB_Size_PDSCH_r18_Supported},
}

type CodebookParametersHARQ_ACK_PUSCH_r18 struct {
	MultiplexingType1_r18        *int64
	MultiplexingType2_r18        *int64
	MultiplexingType3_r18        *int64
	Pucch_DiffResource_PDSCH_r18 *int64
	DiffCB_Size_PDSCH_r18        *int64
}

func (ie *CodebookParametersHARQ_ACK_PUSCH_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(codebookParametersHARQACKPUSCHR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MultiplexingType1_r18 != nil, ie.MultiplexingType2_r18 != nil, ie.MultiplexingType3_r18 != nil, ie.Pucch_DiffResource_PDSCH_r18 != nil, ie.DiffCB_Size_PDSCH_r18 != nil}); err != nil {
		return err
	}

	// 2. multiplexingType1-r18: enumerated
	{
		if ie.MultiplexingType1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultiplexingType1_r18, codebookParametersHARQACKPUSCHR18MultiplexingType1R18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. multiplexingType2-r18: enumerated
	{
		if ie.MultiplexingType2_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultiplexingType2_r18, codebookParametersHARQACKPUSCHR18MultiplexingType2R18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. multiplexingType3-r18: enumerated
	{
		if ie.MultiplexingType3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultiplexingType3_r18, codebookParametersHARQACKPUSCHR18MultiplexingType3R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. pucch-DiffResource-PDSCH-r18: enumerated
	{
		if ie.Pucch_DiffResource_PDSCH_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_DiffResource_PDSCH_r18, codebookParametersHARQACKPUSCHR18PucchDiffResourcePDSCHR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. diffCB-Size-PDSCH-r18: enumerated
	{
		if ie.DiffCB_Size_PDSCH_r18 != nil {
			if err := e.EncodeEnumerated(*ie.DiffCB_Size_PDSCH_r18, codebookParametersHARQACKPUSCHR18DiffCBSizePDSCHR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CodebookParametersHARQ_ACK_PUSCH_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(codebookParametersHARQACKPUSCHR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. multiplexingType1-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(codebookParametersHARQACKPUSCHR18MultiplexingType1R18Constraints)
			if err != nil {
				return err
			}
			ie.MultiplexingType1_r18 = &idx
		}
	}

	// 3. multiplexingType2-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(codebookParametersHARQACKPUSCHR18MultiplexingType2R18Constraints)
			if err != nil {
				return err
			}
			ie.MultiplexingType2_r18 = &idx
		}
	}

	// 4. multiplexingType3-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(codebookParametersHARQACKPUSCHR18MultiplexingType3R18Constraints)
			if err != nil {
				return err
			}
			ie.MultiplexingType3_r18 = &idx
		}
	}

	// 5. pucch-DiffResource-PDSCH-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(codebookParametersHARQACKPUSCHR18PucchDiffResourcePDSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_DiffResource_PDSCH_r18 = &idx
		}
	}

	// 6. diffCB-Size-PDSCH-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(codebookParametersHARQACKPUSCHR18DiffCBSizePDSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.DiffCB_Size_PDSCH_r18 = &idx
		}
	}

	return nil
}
