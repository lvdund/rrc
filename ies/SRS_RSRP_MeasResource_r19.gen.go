// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRS-RSRP-MeasResource-r19 (line 15795).

var sRSRSRPMeasResourceR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-RSRP-MeasResourceId-r19", Optional: true},
		{Name: "srs-Resource-r19", Optional: true},
		{Name: "qcl-InfoPeriodicSRS-RSRP-MeasResource-r19", Optional: true},
	},
}

type SRS_RSRP_MeasResource_r19 struct {
	Srs_RSRP_MeasResourceId_r19               *SRS_RSRP_MeasResourceId_r19
	Srs_Resource_r19                          *SRS_Resource
	Qcl_InfoPeriodicSRS_RSRP_MeasResource_r19 *TCI_StateId
}

func (ie *SRS_RSRP_MeasResource_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRSRSRPMeasResourceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_RSRP_MeasResourceId_r19 != nil, ie.Srs_Resource_r19 != nil, ie.Qcl_InfoPeriodicSRS_RSRP_MeasResource_r19 != nil}); err != nil {
		return err
	}

	// 3. srs-RSRP-MeasResourceId-r19: ref
	{
		if ie.Srs_RSRP_MeasResourceId_r19 != nil {
			if err := ie.Srs_RSRP_MeasResourceId_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. srs-Resource-r19: ref
	{
		if ie.Srs_Resource_r19 != nil {
			if err := ie.Srs_Resource_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. qcl-InfoPeriodicSRS-RSRP-MeasResource-r19: ref
	{
		if ie.Qcl_InfoPeriodicSRS_RSRP_MeasResource_r19 != nil {
			if err := ie.Qcl_InfoPeriodicSRS_RSRP_MeasResource_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SRS_RSRP_MeasResource_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRSRSRPMeasResourceR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srs-RSRP-MeasResourceId-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_RSRP_MeasResourceId_r19 = new(SRS_RSRP_MeasResourceId_r19)
			if err := ie.Srs_RSRP_MeasResourceId_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. srs-Resource-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_Resource_r19 = new(SRS_Resource)
			if err := ie.Srs_Resource_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. qcl-InfoPeriodicSRS-RSRP-MeasResource-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Qcl_InfoPeriodicSRS_RSRP_MeasResource_r19 = new(TCI_StateId)
			if err := ie.Qcl_InfoPeriodicSRS_RSRP_MeasResource_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
