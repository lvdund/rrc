// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CommonLocationInfo-r16 (line 6446).

var commonLocationInfoR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gnss-TOD-msec-r16", Optional: true},
		{Name: "locationTimestamp-r16", Optional: true},
		{Name: "locationCoordinate-r16", Optional: true},
		{Name: "locationError-r16", Optional: true},
		{Name: "locationSource-r16", Optional: true},
		{Name: "velocityEstimate-r16", Optional: true},
	},
}

var commonLocationInfoR16GnssTODMsecR16Constraints = per.SizeConstraints{}

var commonLocationInfoR16LocationTimestampR16Constraints = per.SizeConstraints{}

var commonLocationInfoR16LocationCoordinateR16Constraints = per.SizeConstraints{}

var commonLocationInfoR16LocationErrorR16Constraints = per.SizeConstraints{}

var commonLocationInfoR16LocationSourceR16Constraints = per.SizeConstraints{}

var commonLocationInfoR16VelocityEstimateR16Constraints = per.SizeConstraints{}

type CommonLocationInfo_r16 struct {
	Gnss_TOD_Msec_r16      []byte
	LocationTimestamp_r16  []byte
	LocationCoordinate_r16 []byte
	LocationError_r16      []byte
	LocationSource_r16     []byte
	VelocityEstimate_r16   []byte
}

func (ie *CommonLocationInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(commonLocationInfoR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Gnss_TOD_Msec_r16 != nil, ie.LocationTimestamp_r16 != nil, ie.LocationCoordinate_r16 != nil, ie.LocationError_r16 != nil, ie.LocationSource_r16 != nil, ie.VelocityEstimate_r16 != nil}); err != nil {
		return err
	}

	// 2. gnss-TOD-msec-r16: octet-string
	{
		if ie.Gnss_TOD_Msec_r16 != nil {
			if err := e.EncodeOctetString(ie.Gnss_TOD_Msec_r16, commonLocationInfoR16GnssTODMsecR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. locationTimestamp-r16: octet-string
	{
		if ie.LocationTimestamp_r16 != nil {
			if err := e.EncodeOctetString(ie.LocationTimestamp_r16, commonLocationInfoR16LocationTimestampR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. locationCoordinate-r16: octet-string
	{
		if ie.LocationCoordinate_r16 != nil {
			if err := e.EncodeOctetString(ie.LocationCoordinate_r16, commonLocationInfoR16LocationCoordinateR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. locationError-r16: octet-string
	{
		if ie.LocationError_r16 != nil {
			if err := e.EncodeOctetString(ie.LocationError_r16, commonLocationInfoR16LocationErrorR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. locationSource-r16: octet-string
	{
		if ie.LocationSource_r16 != nil {
			if err := e.EncodeOctetString(ie.LocationSource_r16, commonLocationInfoR16LocationSourceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. velocityEstimate-r16: octet-string
	{
		if ie.VelocityEstimate_r16 != nil {
			if err := e.EncodeOctetString(ie.VelocityEstimate_r16, commonLocationInfoR16VelocityEstimateR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CommonLocationInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(commonLocationInfoR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. gnss-TOD-msec-r16: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(commonLocationInfoR16GnssTODMsecR16Constraints)
			if err != nil {
				return err
			}
			ie.Gnss_TOD_Msec_r16 = v
		}
	}

	// 3. locationTimestamp-r16: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(commonLocationInfoR16LocationTimestampR16Constraints)
			if err != nil {
				return err
			}
			ie.LocationTimestamp_r16 = v
		}
	}

	// 4. locationCoordinate-r16: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(commonLocationInfoR16LocationCoordinateR16Constraints)
			if err != nil {
				return err
			}
			ie.LocationCoordinate_r16 = v
		}
	}

	// 5. locationError-r16: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(commonLocationInfoR16LocationErrorR16Constraints)
			if err != nil {
				return err
			}
			ie.LocationError_r16 = v
		}
	}

	// 6. locationSource-r16: octet-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeOctetString(commonLocationInfoR16LocationSourceR16Constraints)
			if err != nil {
				return err
			}
			ie.LocationSource_r16 = v
		}
	}

	// 7. velocityEstimate-r16: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(commonLocationInfoR16VelocityEstimateR16Constraints)
			if err != nil {
				return err
			}
			ie.VelocityEstimate_r16 = v
		}
	}

	return nil
}
