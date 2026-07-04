// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedBandUTRA-FDD-r16 (line 20932).

const (
	SupportedBandUTRA_FDD_r16_BandI      = 0
	SupportedBandUTRA_FDD_r16_BandII     = 1
	SupportedBandUTRA_FDD_r16_BandIII    = 2
	SupportedBandUTRA_FDD_r16_BandIV     = 3
	SupportedBandUTRA_FDD_r16_BandV      = 4
	SupportedBandUTRA_FDD_r16_BandVI     = 5
	SupportedBandUTRA_FDD_r16_BandVII    = 6
	SupportedBandUTRA_FDD_r16_BandVIII   = 7
	SupportedBandUTRA_FDD_r16_BandIX     = 8
	SupportedBandUTRA_FDD_r16_BandX      = 9
	SupportedBandUTRA_FDD_r16_BandXI     = 10
	SupportedBandUTRA_FDD_r16_BandXII    = 11
	SupportedBandUTRA_FDD_r16_BandXIII   = 12
	SupportedBandUTRA_FDD_r16_BandXIV    = 13
	SupportedBandUTRA_FDD_r16_BandXV     = 14
	SupportedBandUTRA_FDD_r16_BandXVI    = 15
	SupportedBandUTRA_FDD_r16_BandXVII   = 16
	SupportedBandUTRA_FDD_r16_BandXVIII  = 17
	SupportedBandUTRA_FDD_r16_BandXIX    = 18
	SupportedBandUTRA_FDD_r16_BandXX     = 19
	SupportedBandUTRA_FDD_r16_BandXXI    = 20
	SupportedBandUTRA_FDD_r16_BandXXII   = 21
	SupportedBandUTRA_FDD_r16_BandXXIII  = 22
	SupportedBandUTRA_FDD_r16_BandXXIV   = 23
	SupportedBandUTRA_FDD_r16_BandXXV    = 24
	SupportedBandUTRA_FDD_r16_BandXXVI   = 25
	SupportedBandUTRA_FDD_r16_BandXXVII  = 26
	SupportedBandUTRA_FDD_r16_BandXXVIII = 27
	SupportedBandUTRA_FDD_r16_BandXXIX   = 28
	SupportedBandUTRA_FDD_r16_BandXXX    = 29
	SupportedBandUTRA_FDD_r16_BandXXXI   = 30
	SupportedBandUTRA_FDD_r16_BandXXXII  = 31
)

var supportedBandUTRAFDDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedBandUTRA_FDD_r16_BandI, SupportedBandUTRA_FDD_r16_BandII, SupportedBandUTRA_FDD_r16_BandIII, SupportedBandUTRA_FDD_r16_BandIV, SupportedBandUTRA_FDD_r16_BandV, SupportedBandUTRA_FDD_r16_BandVI, SupportedBandUTRA_FDD_r16_BandVII, SupportedBandUTRA_FDD_r16_BandVIII, SupportedBandUTRA_FDD_r16_BandIX, SupportedBandUTRA_FDD_r16_BandX, SupportedBandUTRA_FDD_r16_BandXI, SupportedBandUTRA_FDD_r16_BandXII, SupportedBandUTRA_FDD_r16_BandXIII, SupportedBandUTRA_FDD_r16_BandXIV, SupportedBandUTRA_FDD_r16_BandXV, SupportedBandUTRA_FDD_r16_BandXVI, SupportedBandUTRA_FDD_r16_BandXVII, SupportedBandUTRA_FDD_r16_BandXVIII, SupportedBandUTRA_FDD_r16_BandXIX, SupportedBandUTRA_FDD_r16_BandXX, SupportedBandUTRA_FDD_r16_BandXXI, SupportedBandUTRA_FDD_r16_BandXXII, SupportedBandUTRA_FDD_r16_BandXXIII, SupportedBandUTRA_FDD_r16_BandXXIV, SupportedBandUTRA_FDD_r16_BandXXV, SupportedBandUTRA_FDD_r16_BandXXVI, SupportedBandUTRA_FDD_r16_BandXXVII, SupportedBandUTRA_FDD_r16_BandXXVIII, SupportedBandUTRA_FDD_r16_BandXXIX, SupportedBandUTRA_FDD_r16_BandXXX, SupportedBandUTRA_FDD_r16_BandXXXI, SupportedBandUTRA_FDD_r16_BandXXXII},
}

type SupportedBandUTRA_FDD_r16 struct {
	Value int64
}

func (v *SupportedBandUTRA_FDD_r16) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, supportedBandUTRAFDDR16Constraints)
}

func (v *SupportedBandUTRA_FDD_r16) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(supportedBandUTRAFDDR16Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
