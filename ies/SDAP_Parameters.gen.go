// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SDAP-Parameters (line 24841).

var sDAPParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "as-ReflectiveQoS", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	SDAP_Parameters_As_ReflectiveQoS_True = 0
)

var sDAPParametersAsReflectiveQoSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDAP_Parameters_As_ReflectiveQoS_True},
}

const (
	SDAP_Parameters_Ext_Sdap_QOS_IAB_r16_Supported = 0
)

var sDAPParametersExtSdapQOSIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDAP_Parameters_Ext_Sdap_QOS_IAB_r16_Supported},
}

const (
	SDAP_Parameters_Ext_SdapHeaderIAB_r16_Supported = 0
)

var sDAPParametersExtSdapHeaderIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDAP_Parameters_Ext_SdapHeaderIAB_r16_Supported},
}

const (
	SDAP_Parameters_Ext_Sdap_QOS_NCR_r18_Supported = 0
)

var sDAPParametersExtSdapQOSNCRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDAP_Parameters_Ext_Sdap_QOS_NCR_r18_Supported},
}

const (
	SDAP_Parameters_Ext_Sdap_HeaderNCR_r18_Supported = 0
)

var sDAPParametersExtSdapHeaderNCRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDAP_Parameters_Ext_Sdap_HeaderNCR_r18_Supported},
}

type SDAP_Parameters struct {
	As_ReflectiveQoS   *int64
	Sdap_QOS_IAB_r16   *int64
	SdapHeaderIAB_r16  *int64
	Sdap_QOS_NCR_r18   *int64
	Sdap_HeaderNCR_r18 *int64
}

func (ie *SDAP_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDAPParametersConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sdap_QOS_IAB_r16 != nil || ie.SdapHeaderIAB_r16 != nil
	hasExtGroup1 := ie.Sdap_QOS_NCR_r18 != nil || ie.Sdap_HeaderNCR_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.As_ReflectiveQoS != nil}); err != nil {
		return err
	}

	// 3. as-ReflectiveQoS: enumerated
	{
		if ie.As_ReflectiveQoS != nil {
			if err := e.EncodeEnumerated(*ie.As_ReflectiveQoS, sDAPParametersAsReflectiveQoSConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sdap-QOS-IAB-r16", Optional: true},
					{Name: "sdapHeaderIAB-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sdap_QOS_IAB_r16 != nil, ie.SdapHeaderIAB_r16 != nil}); err != nil {
				return err
			}

			if ie.Sdap_QOS_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Sdap_QOS_IAB_r16, sDAPParametersExtSdapQOSIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.SdapHeaderIAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SdapHeaderIAB_r16, sDAPParametersExtSdapHeaderIABR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sdap-QOS-NCR-r18", Optional: true},
					{Name: "sdap-HeaderNCR-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sdap_QOS_NCR_r18 != nil, ie.Sdap_HeaderNCR_r18 != nil}); err != nil {
				return err
			}

			if ie.Sdap_QOS_NCR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sdap_QOS_NCR_r18, sDAPParametersExtSdapQOSNCRR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sdap_HeaderNCR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sdap_HeaderNCR_r18, sDAPParametersExtSdapHeaderNCRR18Constraints); err != nil {
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

func (ie *SDAP_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDAPParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. as-ReflectiveQoS: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sDAPParametersAsReflectiveQoSConstraints)
			if err != nil {
				return err
			}
			ie.As_ReflectiveQoS = &idx
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
				{Name: "sdap-QOS-IAB-r16", Optional: true},
				{Name: "sdapHeaderIAB-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sDAPParametersExtSdapQOSIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Sdap_QOS_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sDAPParametersExtSdapHeaderIABR16Constraints)
			if err != nil {
				return err
			}
			ie.SdapHeaderIAB_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sdap-QOS-NCR-r18", Optional: true},
				{Name: "sdap-HeaderNCR-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sDAPParametersExtSdapQOSNCRR18Constraints)
			if err != nil {
				return err
			}
			ie.Sdap_QOS_NCR_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sDAPParametersExtSdapHeaderNCRR18Constraints)
			if err != nil {
				return err
			}
			ie.Sdap_HeaderNCR_r18 = &v
		}
	}

	return nil
}
