// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IMS-Parameters (line 20858).

var iMSParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ims-ParametersCommon", Optional: true},
		{Name: "ims-ParametersFRX-Diff", Optional: true},
	},
}

type IMS_Parameters struct {
	Ims_ParametersCommon   *IMS_ParametersCommon
	Ims_ParametersFRX_Diff *IMS_ParametersFRX_Diff
}

func (ie *IMS_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iMSParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ims_ParametersCommon != nil, ie.Ims_ParametersFRX_Diff != nil}); err != nil {
		return err
	}

	// 3. ims-ParametersCommon: ref
	{
		if ie.Ims_ParametersCommon != nil {
			if err := ie.Ims_ParametersCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ims-ParametersFRX-Diff: ref
	{
		if ie.Ims_ParametersFRX_Diff != nil {
			if err := ie.Ims_ParametersFRX_Diff.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IMS_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iMSParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ims-ParametersCommon: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ims_ParametersCommon = new(IMS_ParametersCommon)
			if err := ie.Ims_ParametersCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ims-ParametersFRX-Diff: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ims_ParametersFRX_Diff = new(IMS_ParametersFRX_Diff)
			if err := ie.Ims_ParametersFRX_Diff.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
