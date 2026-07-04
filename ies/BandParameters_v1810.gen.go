// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParameters-v1810 (line 17095).

var bandParametersV1810Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-AntennaSwitching8T8R-r18", Optional: true},
	},
}

var bandParametersV1810SrsAntennaSwitching8T8RR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "antennaSwitch8T8R-r18", Optional: true},
		{Name: "downgradeConfig-r18", Optional: true},
		{Name: "entryNumberAffect-r18", Optional: true},
		{Name: "entryNumberSwitch-r18", Optional: true},
	},
}

const (
	BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_AntennaSwitch8T8R_r18_NoTdm       = 0
	BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_AntennaSwitch8T8R_r18_TdmAndNoTdm = 1
)

var bandParametersV1810SrsAntennaSwitching8T8RR18AntennaSwitch8T8RR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_AntennaSwitch8T8R_r18_NoTdm, BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_AntennaSwitch8T8R_r18_TdmAndNoTdm},
}

var bandParametersV1810SrsAntennaSwitching8T8RR18DowngradeConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "empty-r18"},
		{Name: "downgrade-r18"},
	},
}

const (
	BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_DowngradeConfig_r18_Empty_r18     = 0
	BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_DowngradeConfig_r18_Downgrade_r18 = 1
)

type BandParameters_v1810 struct {
	Srs_AntennaSwitching8T8R_r18 *struct {
		AntennaSwitch8T8R_r18 *int64
		DowngradeConfig_r18   *struct {
			Choice        int
			Empty_r18     *struct{}
			Downgrade_r18 *per.BitString
		}
		EntryNumberAffect_r18 *int64
		EntryNumberSwitch_r18 *int64
	}
}

func (ie *BandParameters_v1810) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1810Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_AntennaSwitching8T8R_r18 != nil}); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching8T8R-r18: sequence
	{
		if ie.Srs_AntennaSwitching8T8R_r18 != nil {
			c := ie.Srs_AntennaSwitching8T8R_r18
			bandParametersV1810SrsAntennaSwitching8T8RR18Seq := e.NewSequenceEncoder(bandParametersV1810SrsAntennaSwitching8T8RR18Constraints)
			if err := bandParametersV1810SrsAntennaSwitching8T8RR18Seq.EncodePreamble([]bool{c.AntennaSwitch8T8R_r18 != nil, c.DowngradeConfig_r18 != nil, c.EntryNumberAffect_r18 != nil, c.EntryNumberSwitch_r18 != nil}); err != nil {
				return err
			}
			if c.AntennaSwitch8T8R_r18 != nil {
				if err := e.EncodeEnumerated((*c.AntennaSwitch8T8R_r18), bandParametersV1810SrsAntennaSwitching8T8RR18AntennaSwitch8T8RR18Constraints); err != nil {
					return err
				}
			}
			if c.DowngradeConfig_r18 != nil {
				choiceEnc := e.NewChoiceEncoder(bandParametersV1810SrsAntennaSwitching8T8RR18DowngradeConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.DowngradeConfig_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.DowngradeConfig_r18).Choice {
				case BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_DowngradeConfig_r18_Empty_r18:
					if err := e.EncodeNull(); err != nil {
						return err
					}
				case BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_DowngradeConfig_r18_Downgrade_r18:
					if err := e.EncodeBitString((*(*c.DowngradeConfig_r18).Downgrade_r18), per.FixedSize(11)); err != nil {
						return err
					}
				}
			}
			if c.EntryNumberAffect_r18 != nil {
				if err := e.EncodeInteger((*c.EntryNumberAffect_r18), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.EntryNumberSwitch_r18 != nil {
				if err := e.EncodeInteger((*c.EntryNumberSwitch_r18), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1810) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1810Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching8T8R-r18: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_AntennaSwitching8T8R_r18 = &struct {
				AntennaSwitch8T8R_r18 *int64
				DowngradeConfig_r18   *struct {
					Choice        int
					Empty_r18     *struct{}
					Downgrade_r18 *per.BitString
				}
				EntryNumberAffect_r18 *int64
				EntryNumberSwitch_r18 *int64
			}{}
			c := ie.Srs_AntennaSwitching8T8R_r18
			bandParametersV1810SrsAntennaSwitching8T8RR18Seq := d.NewSequenceDecoder(bandParametersV1810SrsAntennaSwitching8T8RR18Constraints)
			if err := bandParametersV1810SrsAntennaSwitching8T8RR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandParametersV1810SrsAntennaSwitching8T8RR18Seq.IsComponentPresent(0) {
				c.AntennaSwitch8T8R_r18 = new(int64)
				v, err := d.DecodeEnumerated(bandParametersV1810SrsAntennaSwitching8T8RR18AntennaSwitch8T8RR18Constraints)
				if err != nil {
					return err
				}
				(*c.AntennaSwitch8T8R_r18) = v
			}
			if bandParametersV1810SrsAntennaSwitching8T8RR18Seq.IsComponentPresent(1) {
				c.DowngradeConfig_r18 = &struct {
					Choice        int
					Empty_r18     *struct{}
					Downgrade_r18 *per.BitString
				}{}
				choiceDec := d.NewChoiceDecoder(bandParametersV1810SrsAntennaSwitching8T8RR18DowngradeConfigR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.DowngradeConfig_r18).Choice = int(index)
				switch index {
				case BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_DowngradeConfig_r18_Empty_r18:
					(*c.DowngradeConfig_r18).Empty_r18 = &struct{}{}
					if err := d.DecodeNull(); err != nil {
						return err
					}
				case BandParameters_v1810_Srs_AntennaSwitching8T8R_r18_DowngradeConfig_r18_Downgrade_r18:
					(*c.DowngradeConfig_r18).Downgrade_r18 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(11))
					if err != nil {
						return err
					}
					(*(*c.DowngradeConfig_r18).Downgrade_r18) = v
				}
			}
			if bandParametersV1810SrsAntennaSwitching8T8RR18Seq.IsComponentPresent(2) {
				c.EntryNumberAffect_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberAffect_r18) = v
			}
			if bandParametersV1810SrsAntennaSwitching8T8RR18Seq.IsComponentPresent(3) {
				c.EntryNumberSwitch_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberSwitch_r18) = v
			}
		}
	}

	return nil
}
