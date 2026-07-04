// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParameters-v1900 (line 17108).

var bandParametersV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-AntennaSwitching3T6R-r19", Optional: true},
		{Name: "srs-AntennaSwitching3T3R-r19", Optional: true},
		{Name: "srs-SwitchingSimultaneousList-r19", Optional: true},
	},
}

var bandParametersV1900SrsAntennaSwitching3T6RR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "entryNumberAffect-r19", Optional: true},
		{Name: "entryNumberSwitch-r19", Optional: true},
	},
}

var bandParametersV1900SrsAntennaSwitching3T3RR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "entryNumberAffect-r19", Optional: true},
		{Name: "entryNumberSwitch-r19", Optional: true},
	},
}

type BandParameters_v1900 struct {
	Srs_AntennaSwitching3T6R_r19 *struct {
		EntryNumberAffect_r19 *int64
		EntryNumberSwitch_r19 *int64
	}
	Srs_AntennaSwitching3T3R_r19 *struct {
		EntryNumberAffect_r19 *int64
		EntryNumberSwitch_r19 *int64
	}
	Srs_SwitchingSimultaneousList_r19 *SRS_SwitchingSimultaneousBandsNR_r19
}

func (ie *BandParameters_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_AntennaSwitching3T6R_r19 != nil, ie.Srs_AntennaSwitching3T3R_r19 != nil, ie.Srs_SwitchingSimultaneousList_r19 != nil}); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching3T6R-r19: sequence
	{
		if ie.Srs_AntennaSwitching3T6R_r19 != nil {
			c := ie.Srs_AntennaSwitching3T6R_r19
			bandParametersV1900SrsAntennaSwitching3T6RR19Seq := e.NewSequenceEncoder(bandParametersV1900SrsAntennaSwitching3T6RR19Constraints)
			if err := bandParametersV1900SrsAntennaSwitching3T6RR19Seq.EncodePreamble([]bool{c.EntryNumberAffect_r19 != nil, c.EntryNumberSwitch_r19 != nil}); err != nil {
				return err
			}
			if c.EntryNumberAffect_r19 != nil {
				if err := e.EncodeInteger((*c.EntryNumberAffect_r19), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.EntryNumberSwitch_r19 != nil {
				if err := e.EncodeInteger((*c.EntryNumberSwitch_r19), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		}
	}

	// 3. srs-AntennaSwitching3T3R-r19: sequence
	{
		if ie.Srs_AntennaSwitching3T3R_r19 != nil {
			c := ie.Srs_AntennaSwitching3T3R_r19
			bandParametersV1900SrsAntennaSwitching3T3RR19Seq := e.NewSequenceEncoder(bandParametersV1900SrsAntennaSwitching3T3RR19Constraints)
			if err := bandParametersV1900SrsAntennaSwitching3T3RR19Seq.EncodePreamble([]bool{c.EntryNumberAffect_r19 != nil, c.EntryNumberSwitch_r19 != nil}); err != nil {
				return err
			}
			if c.EntryNumberAffect_r19 != nil {
				if err := e.EncodeInteger((*c.EntryNumberAffect_r19), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			if c.EntryNumberSwitch_r19 != nil {
				if err := e.EncodeInteger((*c.EntryNumberSwitch_r19), per.Constrained(1, 32)); err != nil {
					return err
				}
			}
		}
	}

	// 4. srs-SwitchingSimultaneousList-r19: ref
	{
		if ie.Srs_SwitchingSimultaneousList_r19 != nil {
			if err := ie.Srs_SwitchingSimultaneousList_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-AntennaSwitching3T6R-r19: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_AntennaSwitching3T6R_r19 = &struct {
				EntryNumberAffect_r19 *int64
				EntryNumberSwitch_r19 *int64
			}{}
			c := ie.Srs_AntennaSwitching3T6R_r19
			bandParametersV1900SrsAntennaSwitching3T6RR19Seq := d.NewSequenceDecoder(bandParametersV1900SrsAntennaSwitching3T6RR19Constraints)
			if err := bandParametersV1900SrsAntennaSwitching3T6RR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandParametersV1900SrsAntennaSwitching3T6RR19Seq.IsComponentPresent(0) {
				c.EntryNumberAffect_r19 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberAffect_r19) = v
			}
			if bandParametersV1900SrsAntennaSwitching3T6RR19Seq.IsComponentPresent(1) {
				c.EntryNumberSwitch_r19 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberSwitch_r19) = v
			}
		}
	}

	// 3. srs-AntennaSwitching3T3R-r19: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Srs_AntennaSwitching3T3R_r19 = &struct {
				EntryNumberAffect_r19 *int64
				EntryNumberSwitch_r19 *int64
			}{}
			c := ie.Srs_AntennaSwitching3T3R_r19
			bandParametersV1900SrsAntennaSwitching3T3RR19Seq := d.NewSequenceDecoder(bandParametersV1900SrsAntennaSwitching3T3RR19Constraints)
			if err := bandParametersV1900SrsAntennaSwitching3T3RR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandParametersV1900SrsAntennaSwitching3T3RR19Seq.IsComponentPresent(0) {
				c.EntryNumberAffect_r19 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberAffect_r19) = v
			}
			if bandParametersV1900SrsAntennaSwitching3T3RR19Seq.IsComponentPresent(1) {
				c.EntryNumberSwitch_r19 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				(*c.EntryNumberSwitch_r19) = v
			}
		}
	}

	// 4. srs-SwitchingSimultaneousList-r19: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Srs_SwitchingSimultaneousList_r19 = new(SRS_SwitchingSimultaneousBandsNR_r19)
			if err := ie.Srs_SwitchingSimultaneousList_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
