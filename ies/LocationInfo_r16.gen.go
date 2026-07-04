// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LocationInfo-r16 (line 8549).

var locationInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "commonLocationInfo-r16", Optional: true},
		{Name: "bt-LocationInfo-r16", Optional: true},
		{Name: "wlan-LocationInfo-r16", Optional: true},
		{Name: "sensor-LocationInfo-r16", Optional: true},
	},
}

type LocationInfo_r16 struct {
	CommonLocationInfo_r16  *CommonLocationInfo_r16
	Bt_LocationInfo_r16     *LogMeasResultListBT_r16
	Wlan_LocationInfo_r16   *LogMeasResultListWLAN_r16
	Sensor_LocationInfo_r16 *Sensor_LocationInfo_r16
}

func (ie *LocationInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(locationInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CommonLocationInfo_r16 != nil, ie.Bt_LocationInfo_r16 != nil, ie.Wlan_LocationInfo_r16 != nil, ie.Sensor_LocationInfo_r16 != nil}); err != nil {
		return err
	}

	// 3. commonLocationInfo-r16: ref
	{
		if ie.CommonLocationInfo_r16 != nil {
			if err := ie.CommonLocationInfo_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. bt-LocationInfo-r16: ref
	{
		if ie.Bt_LocationInfo_r16 != nil {
			if err := ie.Bt_LocationInfo_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. wlan-LocationInfo-r16: ref
	{
		if ie.Wlan_LocationInfo_r16 != nil {
			if err := ie.Wlan_LocationInfo_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sensor-LocationInfo-r16: ref
	{
		if ie.Sensor_LocationInfo_r16 != nil {
			if err := ie.Sensor_LocationInfo_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LocationInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(locationInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. commonLocationInfo-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.CommonLocationInfo_r16 = new(CommonLocationInfo_r16)
			if err := ie.CommonLocationInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. bt-LocationInfo-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Bt_LocationInfo_r16 = new(LogMeasResultListBT_r16)
			if err := ie.Bt_LocationInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. wlan-LocationInfo-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Wlan_LocationInfo_r16 = new(LogMeasResultListWLAN_r16)
			if err := ie.Wlan_LocationInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sensor-LocationInfo-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Sensor_LocationInfo_r16 = new(Sensor_LocationInfo_r16)
			if err := ie.Sensor_LocationInfo_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
