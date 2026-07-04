// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TA-InfoIoT-r19 (line 10810).

var tAInfoIoTR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ta-CommonIoT-r19", Optional: true},
		{Name: "ta-CommonDriftIoT-r19", Optional: true},
		{Name: "ta-CommonDriftVariantIoT-r19", Optional: true},
	},
}

var tAInfoIoTR19TaCommonIoTR19Constraints = per.Constrained(0, 66485757)

var tAInfoIoTR19TaCommonDriftIoTR19Constraints = per.Constrained(-261935, 261935)

var tAInfoIoTR19TaCommonDriftVariantIoTR19Constraints = per.Constrained(0, 29479)

type TA_InfoIoT_r19 struct {
	Ta_CommonIoT_r19             *int64
	Ta_CommonDriftIoT_r19        *int64
	Ta_CommonDriftVariantIoT_r19 *int64
}

func (ie *TA_InfoIoT_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tAInfoIoTR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ta_CommonIoT_r19 != nil, ie.Ta_CommonDriftIoT_r19 != nil, ie.Ta_CommonDriftVariantIoT_r19 != nil}); err != nil {
		return err
	}

	// 2. ta-CommonIoT-r19: integer
	{
		if ie.Ta_CommonIoT_r19 != nil {
			if err := e.EncodeInteger(*ie.Ta_CommonIoT_r19, tAInfoIoTR19TaCommonIoTR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ta-CommonDriftIoT-r19: integer
	{
		if ie.Ta_CommonDriftIoT_r19 != nil {
			if err := e.EncodeInteger(*ie.Ta_CommonDriftIoT_r19, tAInfoIoTR19TaCommonDriftIoTR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ta-CommonDriftVariantIoT-r19: integer
	{
		if ie.Ta_CommonDriftVariantIoT_r19 != nil {
			if err := e.EncodeInteger(*ie.Ta_CommonDriftVariantIoT_r19, tAInfoIoTR19TaCommonDriftVariantIoTR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *TA_InfoIoT_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tAInfoIoTR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ta-CommonIoT-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(tAInfoIoTR19TaCommonIoTR19Constraints)
			if err != nil {
				return err
			}
			ie.Ta_CommonIoT_r19 = &v
		}
	}

	// 3. ta-CommonDriftIoT-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(tAInfoIoTR19TaCommonDriftIoTR19Constraints)
			if err != nil {
				return err
			}
			ie.Ta_CommonDriftIoT_r19 = &v
		}
	}

	// 4. ta-CommonDriftVariantIoT-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(tAInfoIoTR19TaCommonDriftVariantIoTR19Constraints)
			if err != nil {
				return err
			}
			ie.Ta_CommonDriftVariantIoT_r19 = &v
		}
	}

	return nil
}
