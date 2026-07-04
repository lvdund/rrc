// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReferenceSignalConfig (line 9496).

var referenceSignalConfigConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-ConfigMobility", Optional: true},
		{Name: "csi-rs-ResourceConfigMobility", Optional: true},
	},
}

var referenceSignalConfigCsiRsResourceConfigMobilityConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ReferenceSignalConfig_Csi_Rs_ResourceConfigMobility_Release = 0
	ReferenceSignalConfig_Csi_Rs_ResourceConfigMobility_Setup   = 1
)

type ReferenceSignalConfig struct {
	Ssb_ConfigMobility            *SSB_ConfigMobility
	Csi_Rs_ResourceConfigMobility *struct {
		Choice  int
		Release *struct{}
		Setup   *CSI_RS_ResourceConfigMobility
	}
}

func (ie *ReferenceSignalConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(referenceSignalConfigConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_ConfigMobility != nil, ie.Csi_Rs_ResourceConfigMobility != nil}); err != nil {
		return err
	}

	// 2. ssb-ConfigMobility: ref
	{
		if ie.Ssb_ConfigMobility != nil {
			if err := ie.Ssb_ConfigMobility.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. csi-rs-ResourceConfigMobility: choice
	{
		if ie.Csi_Rs_ResourceConfigMobility != nil {
			choiceEnc := e.NewChoiceEncoder(referenceSignalConfigCsiRsResourceConfigMobilityConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Csi_Rs_ResourceConfigMobility).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Csi_Rs_ResourceConfigMobility).Choice {
			case ReferenceSignalConfig_Csi_Rs_ResourceConfigMobility_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ReferenceSignalConfig_Csi_Rs_ResourceConfigMobility_Setup:
				if err := (*ie.Csi_Rs_ResourceConfigMobility).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Csi_Rs_ResourceConfigMobility).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *ReferenceSignalConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(referenceSignalConfigConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ssb-ConfigMobility: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ssb_ConfigMobility = new(SSB_ConfigMobility)
			if err := ie.Ssb_ConfigMobility.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. csi-rs-ResourceConfigMobility: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Csi_Rs_ResourceConfigMobility = &struct {
				Choice  int
				Release *struct{}
				Setup   *CSI_RS_ResourceConfigMobility
			}{}
			choiceDec := d.NewChoiceDecoder(referenceSignalConfigCsiRsResourceConfigMobilityConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Csi_Rs_ResourceConfigMobility).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ReferenceSignalConfig_Csi_Rs_ResourceConfigMobility_Release:
				(*ie.Csi_Rs_ResourceConfigMobility).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ReferenceSignalConfig_Csi_Rs_ResourceConfigMobility_Setup:
				(*ie.Csi_Rs_ResourceConfigMobility).Setup = new(CSI_RS_ResourceConfigMobility)
				if err := (*ie.Csi_Rs_ResourceConfigMobility).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
