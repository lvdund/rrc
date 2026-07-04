// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SharedSpectrumChAccessParamsPerBand-v1630 (line 24944).

var sharedSpectrumChAccessParamsPerBandV1630Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-ReceptionIntraCellGuardband-r16", Optional: true},
		{Name: "dl-ReceptionLBT-subsetRB-r16", Optional: true},
	},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1630_Dl_ReceptionIntraCellGuardband_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1630DlReceptionIntraCellGuardbandR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1630_Dl_ReceptionIntraCellGuardband_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1630_Dl_ReceptionLBT_SubsetRB_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1630DlReceptionLBTSubsetRBR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1630_Dl_ReceptionLBT_SubsetRB_r16_Supported},
}

type SharedSpectrumChAccessParamsPerBand_v1630 struct {
	Dl_ReceptionIntraCellGuardband_r16 *int64
	Dl_ReceptionLBT_SubsetRB_r16       *int64
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1630) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sharedSpectrumChAccessParamsPerBandV1630Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_ReceptionIntraCellGuardband_r16 != nil, ie.Dl_ReceptionLBT_SubsetRB_r16 != nil}); err != nil {
		return err
	}

	// 2. dl-ReceptionIntraCellGuardband-r16: enumerated
	{
		if ie.Dl_ReceptionIntraCellGuardband_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_ReceptionIntraCellGuardband_r16, sharedSpectrumChAccessParamsPerBandV1630DlReceptionIntraCellGuardbandR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. dl-ReceptionLBT-subsetRB-r16: enumerated
	{
		if ie.Dl_ReceptionLBT_SubsetRB_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dl_ReceptionLBT_SubsetRB_r16, sharedSpectrumChAccessParamsPerBandV1630DlReceptionLBTSubsetRBR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1630) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sharedSpectrumChAccessParamsPerBandV1630Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dl-ReceptionIntraCellGuardband-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1630DlReceptionIntraCellGuardbandR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_ReceptionIntraCellGuardband_r16 = &idx
		}
	}

	// 3. dl-ReceptionLBT-subsetRB-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1630DlReceptionLBTSubsetRBR16Constraints)
			if err != nil {
				return err
			}
			ie.Dl_ReceptionLBT_SubsetRB_r16 = &idx
		}
	}

	return nil
}
