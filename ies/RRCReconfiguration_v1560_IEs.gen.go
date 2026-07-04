// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-v1560-IEs (line 996).

var rRCReconfigurationV1560IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mrdc-SecondaryCellGroupConfig", Optional: true},
		{Name: "radioBearerConfig2", Optional: true},
		{Name: "sk-Counter", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfiguration_v1560_IEsMrdcSecondaryCellGroupConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1560_IEs_Mrdc_SecondaryCellGroupConfig_Release = 0
	RRCReconfiguration_v1560_IEs_Mrdc_SecondaryCellGroupConfig_Setup   = 1
)

var rRCReconfigurationV1560IEsRadioBearerConfig2Constraints = per.SizeConstraints{}

type RRCReconfiguration_v1560_IEs struct {
	Mrdc_SecondaryCellGroupConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *MRDC_SecondaryCellGroupConfig
	}
	RadioBearerConfig2   []byte
	Sk_Counter           *SK_Counter
	NonCriticalExtension *RRCReconfiguration_v1610_IEs
}

func (ie *RRCReconfiguration_v1560_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV1560IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mrdc_SecondaryCellGroupConfig != nil, ie.RadioBearerConfig2 != nil, ie.Sk_Counter != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. mrdc-SecondaryCellGroupConfig: choice
	{
		if ie.Mrdc_SecondaryCellGroupConfig != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1560_IEsMrdcSecondaryCellGroupConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Mrdc_SecondaryCellGroupConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Mrdc_SecondaryCellGroupConfig).Choice {
			case RRCReconfiguration_v1560_IEs_Mrdc_SecondaryCellGroupConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1560_IEs_Mrdc_SecondaryCellGroupConfig_Setup:
				if err := (*ie.Mrdc_SecondaryCellGroupConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Mrdc_SecondaryCellGroupConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. radioBearerConfig2: octet-string
	{
		if ie.RadioBearerConfig2 != nil {
			if err := e.EncodeOctetString(ie.RadioBearerConfig2, rRCReconfigurationV1560IEsRadioBearerConfig2Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sk-Counter: ref
	{
		if ie.Sk_Counter != nil {
			if err := ie.Sk_Counter.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfiguration_v1560_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV1560IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mrdc-SecondaryCellGroupConfig: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Mrdc_SecondaryCellGroupConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *MRDC_SecondaryCellGroupConfig
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1560_IEsMrdcSecondaryCellGroupConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Mrdc_SecondaryCellGroupConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1560_IEs_Mrdc_SecondaryCellGroupConfig_Release:
				(*ie.Mrdc_SecondaryCellGroupConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1560_IEs_Mrdc_SecondaryCellGroupConfig_Setup:
				(*ie.Mrdc_SecondaryCellGroupConfig).Setup = new(MRDC_SecondaryCellGroupConfig)
				if err := (*ie.Mrdc_SecondaryCellGroupConfig).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. radioBearerConfig2: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(rRCReconfigurationV1560IEsRadioBearerConfig2Constraints)
			if err != nil {
				return err
			}
			ie.RadioBearerConfig2 = v
		}
	}

	// 4. sk-Counter: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sk_Counter = new(SK_Counter)
			if err := ie.Sk_Counter.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(3) {
			ie.NonCriticalExtension = new(RRCReconfiguration_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
