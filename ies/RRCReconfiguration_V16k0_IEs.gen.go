// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-v16k0-IEs (line 1066).

var rRCReconfigurationV16k0IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfigDedicatedNR-v16k0", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfiguration_V16k0_IEsSlConfigDedicatedNRV16k0Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_V16k0_IEs_Sl_ConfigDedicatedNR_V16k0_Release = 0
	RRCReconfiguration_V16k0_IEs_Sl_ConfigDedicatedNR_V16k0_Setup   = 1
)

type RRCReconfiguration_V16k0_IEs struct {
	Sl_ConfigDedicatedNR_V16k0 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_ConfigDedicatedNR_V16k0
	}
	NonCriticalExtension *struct{}
}

func (ie *RRCReconfiguration_V16k0_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV16k0IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfigDedicatedNR_V16k0 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedNR-v16k0: choice
	{
		if ie.Sl_ConfigDedicatedNR_V16k0 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_V16k0_IEsSlConfigDedicatedNRV16k0Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_ConfigDedicatedNR_V16k0).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_ConfigDedicatedNR_V16k0).Choice {
			case RRCReconfiguration_V16k0_IEs_Sl_ConfigDedicatedNR_V16k0_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_V16k0_IEs_Sl_ConfigDedicatedNR_V16k0_Setup:
				if err := (*ie.Sl_ConfigDedicatedNR_V16k0).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_ConfigDedicatedNR_V16k0).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *RRCReconfiguration_V16k0_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV16k0IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedNR-v16k0: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_ConfigDedicatedNR_V16k0 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_ConfigDedicatedNR_V16k0
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_V16k0_IEsSlConfigDedicatedNRV16k0Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_ConfigDedicatedNR_V16k0).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_V16k0_IEs_Sl_ConfigDedicatedNR_V16k0_Release:
				(*ie.Sl_ConfigDedicatedNR_V16k0).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_V16k0_IEs_Sl_ConfigDedicatedNR_V16k0_Setup:
				(*ie.Sl_ConfigDedicatedNR_V16k0).Setup = new(SL_ConfigDedicatedNR_V16k0)
				if err := (*ie.Sl_ConfigDedicatedNR_V16k0).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
