// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-ReportCJTC-r19 (line 7330).

var cSIReportCJTCR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "associatedSRS-ResourceSet-r19", Optional: true},
		{Name: "linkedCJTC-Report-r19", Optional: true},
		{Name: "valueOfAD-r19", Optional: true},
		{Name: "valueOfMD-r19", Optional: true},
		{Name: "valueOfAFO-r19", Optional: true},
		{Name: "valueOfMFO-r19", Optional: true},
		{Name: "valueOfMPhi-r19", Optional: true},
		{Name: "subbandSizeCJTC-r19", Optional: true},
		{Name: "nrofSubbandsPO-r19", Optional: true},
	},
}

const (
	CSI_ReportCJTC_r19_ValueOfAD_r19_Dot5 = 0
	CSI_ReportCJTC_r19_ValueOfAD_r19_One  = 1
)

var cSIReportCJTCR19ValueOfADR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportCJTC_r19_ValueOfAD_r19_Dot5, CSI_ReportCJTC_r19_ValueOfAD_r19_One},
}

const (
	CSI_ReportCJTC_r19_ValueOfMD_r19_N32  = 0
	CSI_ReportCJTC_r19_ValueOfMD_r19_N64  = 1
	CSI_ReportCJTC_r19_ValueOfMD_r19_N128 = 2
	CSI_ReportCJTC_r19_ValueOfMD_r19_N256 = 3
)

var cSIReportCJTCR19ValueOfMDR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportCJTC_r19_ValueOfMD_r19_N32, CSI_ReportCJTC_r19_ValueOfMD_r19_N64, CSI_ReportCJTC_r19_ValueOfMD_r19_N128, CSI_ReportCJTC_r19_ValueOfMD_r19_N256},
}

const (
	CSI_ReportCJTC_r19_ValueOfAFO_r19_ZeroDot1 = 0
	CSI_ReportCJTC_r19_ValueOfAFO_r19_ZeroDot2 = 1
)

var cSIReportCJTCR19ValueOfAFOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportCJTC_r19_ValueOfAFO_r19_ZeroDot1, CSI_ReportCJTC_r19_ValueOfAFO_r19_ZeroDot2},
}

const (
	CSI_ReportCJTC_r19_ValueOfMFO_r19_N16  = 0
	CSI_ReportCJTC_r19_ValueOfMFO_r19_N32  = 1
	CSI_ReportCJTC_r19_ValueOfMFO_r19_N256 = 2
)

var cSIReportCJTCR19ValueOfMFOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportCJTC_r19_ValueOfMFO_r19_N16, CSI_ReportCJTC_r19_ValueOfMFO_r19_N32, CSI_ReportCJTC_r19_ValueOfMFO_r19_N256},
}

const (
	CSI_ReportCJTC_r19_ValueOfMPhi_r19_N16 = 0
	CSI_ReportCJTC_r19_ValueOfMPhi_r19_N32 = 1
)

var cSIReportCJTCR19ValueOfMPhiR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportCJTC_r19_ValueOfMPhi_r19_N16, CSI_ReportCJTC_r19_ValueOfMPhi_r19_N32},
}

const (
	CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N1       = 0
	CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N2       = 1
	CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N4       = 2
	CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N8       = 3
	CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N16      = 4
	CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_Wideband = 5
)

var cSIReportCJTCR19SubbandSizeCJTCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N1, CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N2, CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N4, CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N8, CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_N16, CSI_ReportCJTC_r19_SubbandSizeCJTC_r19_Wideband},
}

var cSIReportCJTCR19NrofSubbandsPOR19Constraints = per.SizeRange(1, 16)

var cSIReportCJTCR19AssociatedSRSResourceSetR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-ResourceSetId-r19"},
		{Name: "srs-ResourceId-r19"},
		{Name: "referenceAntennaPort-r19", Optional: true},
	},
}

type CSI_ReportCJTC_r19 struct {
	AssociatedSRS_ResourceSet_r19 *struct {
		Srs_ResourceSetId_r19    SRS_ResourceSetId
		Srs_ResourceId_r19       SRS_ResourceId
		ReferenceAntennaPort_r19 *int64
	}
	LinkedCJTC_Report_r19 *CSI_ReportConfigId
	ValueOfAD_r19         *int64
	ValueOfMD_r19         *int64
	ValueOfAFO_r19        *int64
	ValueOfMFO_r19        *int64
	ValueOfMPhi_r19       *int64
	SubbandSizeCJTC_r19   *int64
	NrofSubbandsPO_r19    []int64
}

func (ie *CSI_ReportCJTC_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIReportCJTCR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AssociatedSRS_ResourceSet_r19 != nil, ie.LinkedCJTC_Report_r19 != nil, ie.ValueOfAD_r19 != nil, ie.ValueOfMD_r19 != nil, ie.ValueOfAFO_r19 != nil, ie.ValueOfMFO_r19 != nil, ie.ValueOfMPhi_r19 != nil, ie.SubbandSizeCJTC_r19 != nil, ie.NrofSubbandsPO_r19 != nil}); err != nil {
		return err
	}

	// 2. associatedSRS-ResourceSet-r19: sequence
	{
		if ie.AssociatedSRS_ResourceSet_r19 != nil {
			c := ie.AssociatedSRS_ResourceSet_r19
			cSIReportCJTCR19AssociatedSRSResourceSetR19Seq := e.NewSequenceEncoder(cSIReportCJTCR19AssociatedSRSResourceSetR19Constraints)
			if err := cSIReportCJTCR19AssociatedSRSResourceSetR19Seq.EncodePreamble([]bool{c.ReferenceAntennaPort_r19 != nil}); err != nil {
				return err
			}
			if err := c.Srs_ResourceSetId_r19.Encode(e); err != nil {
				return err
			}
			if err := c.Srs_ResourceId_r19.Encode(e); err != nil {
				return err
			}
			if c.ReferenceAntennaPort_r19 != nil {
				if err := e.EncodeInteger((*c.ReferenceAntennaPort_r19), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
		}
	}

	// 3. linkedCJTC-Report-r19: ref
	{
		if ie.LinkedCJTC_Report_r19 != nil {
			if err := ie.LinkedCJTC_Report_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. valueOfAD-r19: enumerated
	{
		if ie.ValueOfAD_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ValueOfAD_r19, cSIReportCJTCR19ValueOfADR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. valueOfMD-r19: enumerated
	{
		if ie.ValueOfMD_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ValueOfMD_r19, cSIReportCJTCR19ValueOfMDR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. valueOfAFO-r19: enumerated
	{
		if ie.ValueOfAFO_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ValueOfAFO_r19, cSIReportCJTCR19ValueOfAFOR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. valueOfMFO-r19: enumerated
	{
		if ie.ValueOfMFO_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ValueOfMFO_r19, cSIReportCJTCR19ValueOfMFOR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. valueOfMPhi-r19: enumerated
	{
		if ie.ValueOfMPhi_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ValueOfMPhi_r19, cSIReportCJTCR19ValueOfMPhiR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. subbandSizeCJTC-r19: enumerated
	{
		if ie.SubbandSizeCJTC_r19 != nil {
			if err := e.EncodeEnumerated(*ie.SubbandSizeCJTC_r19, cSIReportCJTCR19SubbandSizeCJTCR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. nrofSubbandsPO-r19: sequence-of
	{
		if ie.NrofSubbandsPO_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(cSIReportCJTCR19NrofSubbandsPOR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.NrofSubbandsPO_r19))); err != nil {
				return err
			}
			for i := range ie.NrofSubbandsPO_r19 {
				if err := e.EncodeInteger(ie.NrofSubbandsPO_r19[i], per.Constrained(1, 275)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CSI_ReportCJTC_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIReportCJTCR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. associatedSRS-ResourceSet-r19: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.AssociatedSRS_ResourceSet_r19 = &struct {
				Srs_ResourceSetId_r19    SRS_ResourceSetId
				Srs_ResourceId_r19       SRS_ResourceId
				ReferenceAntennaPort_r19 *int64
			}{}
			c := ie.AssociatedSRS_ResourceSet_r19
			cSIReportCJTCR19AssociatedSRSResourceSetR19Seq := d.NewSequenceDecoder(cSIReportCJTCR19AssociatedSRSResourceSetR19Constraints)
			if err := cSIReportCJTCR19AssociatedSRSResourceSetR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				if err := c.Srs_ResourceSetId_r19.Decode(d); err != nil {
					return err
				}
			}
			{
				if err := c.Srs_ResourceId_r19.Decode(d); err != nil {
					return err
				}
			}
			if cSIReportCJTCR19AssociatedSRSResourceSetR19Seq.IsComponentPresent(2) {
				c.ReferenceAntennaPort_r19 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.ReferenceAntennaPort_r19) = v
			}
		}
	}

	// 3. linkedCJTC-Report-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.LinkedCJTC_Report_r19 = new(CSI_ReportConfigId)
			if err := ie.LinkedCJTC_Report_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. valueOfAD-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cSIReportCJTCR19ValueOfADR19Constraints)
			if err != nil {
				return err
			}
			ie.ValueOfAD_r19 = &idx
		}
	}

	// 5. valueOfMD-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cSIReportCJTCR19ValueOfMDR19Constraints)
			if err != nil {
				return err
			}
			ie.ValueOfMD_r19 = &idx
		}
	}

	// 6. valueOfAFO-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cSIReportCJTCR19ValueOfAFOR19Constraints)
			if err != nil {
				return err
			}
			ie.ValueOfAFO_r19 = &idx
		}
	}

	// 7. valueOfMFO-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cSIReportCJTCR19ValueOfMFOR19Constraints)
			if err != nil {
				return err
			}
			ie.ValueOfMFO_r19 = &idx
		}
	}

	// 8. valueOfMPhi-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(cSIReportCJTCR19ValueOfMPhiR19Constraints)
			if err != nil {
				return err
			}
			ie.ValueOfMPhi_r19 = &idx
		}
	}

	// 9. subbandSizeCJTC-r19: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(cSIReportCJTCR19SubbandSizeCJTCR19Constraints)
			if err != nil {
				return err
			}
			ie.SubbandSizeCJTC_r19 = &idx
		}
	}

	// 10. nrofSubbandsPO-r19: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(cSIReportCJTCR19NrofSubbandsPOR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NrofSubbandsPO_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(1, 275))
				if err != nil {
					return err
				}
				ie.NrofSubbandsPO_r19[i] = v
			}
		}
	}

	return nil
}
