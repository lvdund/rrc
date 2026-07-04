// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Aerial-Config-r18 (line 4963).

var aerialConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "flightPathUpdateThrConfig-r18", Optional: true},
	},
}

var aerialConfigR18FlightPathUpdateThrConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "flightPathUpdateDistanceThr-r18", Optional: true},
		{Name: "flightPathUpdateTimeThr-r18", Optional: true},
	},
}

var aerialConfigR18FlightPathUpdateThrConfigR18FlightPathUpdateDistanceThrR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateDistanceThr_r18_Release = 0
	Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateDistanceThr_r18_Setup   = 1
)

var aerialConfigR18FlightPathUpdateThrConfigR18FlightPathUpdateTimeThrR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateTimeThr_r18_Release = 0
	Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateTimeThr_r18_Setup   = 1
)

type Aerial_Config_r18 struct {
	FlightPathUpdateThrConfig_r18 *struct {
		FlightPathUpdateDistanceThr_r18 *struct {
			Choice  int
			Release *struct{}
			Setup   *FlightPathUpdateDistanceThr_r18
		}
		FlightPathUpdateTimeThr_r18 *struct {
			Choice  int
			Release *struct{}
			Setup   *FlightPathUpdateTimeThr_r18
		}
	}
}

func (ie *Aerial_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(aerialConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FlightPathUpdateThrConfig_r18 != nil}); err != nil {
		return err
	}

	// 3. flightPathUpdateThrConfig-r18: sequence
	{
		if ie.FlightPathUpdateThrConfig_r18 != nil {
			c := ie.FlightPathUpdateThrConfig_r18
			aerialConfigR18FlightPathUpdateThrConfigR18Seq := e.NewSequenceEncoder(aerialConfigR18FlightPathUpdateThrConfigR18Constraints)
			if err := aerialConfigR18FlightPathUpdateThrConfigR18Seq.EncodePreamble([]bool{c.FlightPathUpdateDistanceThr_r18 != nil, c.FlightPathUpdateTimeThr_r18 != nil}); err != nil {
				return err
			}
			if c.FlightPathUpdateDistanceThr_r18 != nil {
				choiceEnc := e.NewChoiceEncoder(aerialConfigR18FlightPathUpdateThrConfigR18FlightPathUpdateDistanceThrR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.FlightPathUpdateDistanceThr_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.FlightPathUpdateDistanceThr_r18).Choice {
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateDistanceThr_r18_Release:
					if err := e.EncodeNull(); err != nil {
						return err
					}
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateDistanceThr_r18_Setup:
					if err := (*c.FlightPathUpdateDistanceThr_r18).Setup.Encode(e); err != nil {
						return err
					}
				}
			}
			if c.FlightPathUpdateTimeThr_r18 != nil {
				choiceEnc := e.NewChoiceEncoder(aerialConfigR18FlightPathUpdateThrConfigR18FlightPathUpdateTimeThrR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.FlightPathUpdateTimeThr_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.FlightPathUpdateTimeThr_r18).Choice {
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateTimeThr_r18_Release:
					if err := e.EncodeNull(); err != nil {
						return err
					}
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateTimeThr_r18_Setup:
					if err := (*c.FlightPathUpdateTimeThr_r18).Setup.Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (ie *Aerial_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(aerialConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. flightPathUpdateThrConfig-r18: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.FlightPathUpdateThrConfig_r18 = &struct {
				FlightPathUpdateDistanceThr_r18 *struct {
					Choice  int
					Release *struct{}
					Setup   *FlightPathUpdateDistanceThr_r18
				}
				FlightPathUpdateTimeThr_r18 *struct {
					Choice  int
					Release *struct{}
					Setup   *FlightPathUpdateTimeThr_r18
				}
			}{}
			c := ie.FlightPathUpdateThrConfig_r18
			aerialConfigR18FlightPathUpdateThrConfigR18Seq := d.NewSequenceDecoder(aerialConfigR18FlightPathUpdateThrConfigR18Constraints)
			if err := aerialConfigR18FlightPathUpdateThrConfigR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if aerialConfigR18FlightPathUpdateThrConfigR18Seq.IsComponentPresent(0) {
				c.FlightPathUpdateDistanceThr_r18 = &struct {
					Choice  int
					Release *struct{}
					Setup   *FlightPathUpdateDistanceThr_r18
				}{}
				choiceDec := d.NewChoiceDecoder(aerialConfigR18FlightPathUpdateThrConfigR18FlightPathUpdateDistanceThrR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.FlightPathUpdateDistanceThr_r18).Choice = int(index)
				switch index {
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateDistanceThr_r18_Release:
					(*c.FlightPathUpdateDistanceThr_r18).Release = &struct{}{}
					if err := d.DecodeNull(); err != nil {
						return err
					}
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateDistanceThr_r18_Setup:
					(*c.FlightPathUpdateDistanceThr_r18).Setup = new(FlightPathUpdateDistanceThr_r18)
					if err := (*c.FlightPathUpdateDistanceThr_r18).Setup.Decode(d); err != nil {
						return err
					}
				}
			}
			if aerialConfigR18FlightPathUpdateThrConfigR18Seq.IsComponentPresent(1) {
				c.FlightPathUpdateTimeThr_r18 = &struct {
					Choice  int
					Release *struct{}
					Setup   *FlightPathUpdateTimeThr_r18
				}{}
				choiceDec := d.NewChoiceDecoder(aerialConfigR18FlightPathUpdateThrConfigR18FlightPathUpdateTimeThrR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.FlightPathUpdateTimeThr_r18).Choice = int(index)
				switch index {
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateTimeThr_r18_Release:
					(*c.FlightPathUpdateTimeThr_r18).Release = &struct{}{}
					if err := d.DecodeNull(); err != nil {
						return err
					}
				case Aerial_Config_r18_FlightPathUpdateThrConfig_r18_FlightPathUpdateTimeThr_r18_Setup:
					(*c.FlightPathUpdateTimeThr_r18).Setup = new(FlightPathUpdateTimeThr_r18)
					if err := (*c.FlightPathUpdateTimeThr_r18).Setup.Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
