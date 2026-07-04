// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: VictimSystemType-r16 (line 2518).

var victimSystemTypeR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gps-r16", Optional: true},
		{Name: "glonass-r16", Optional: true},
		{Name: "bds-r16", Optional: true},
		{Name: "galileo-r16", Optional: true},
		{Name: "navIC-r16", Optional: true},
		{Name: "wlan-r16", Optional: true},
		{Name: "bluetooth-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	VictimSystemType_r16_Gps_r16_True = 0
)

var victimSystemTypeR16GpsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_Gps_r16_True},
}

const (
	VictimSystemType_r16_Glonass_r16_True = 0
)

var victimSystemTypeR16GlonassR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_Glonass_r16_True},
}

const (
	VictimSystemType_r16_Bds_r16_True = 0
)

var victimSystemTypeR16BdsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_Bds_r16_True},
}

const (
	VictimSystemType_r16_Galileo_r16_True = 0
)

var victimSystemTypeR16GalileoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_Galileo_r16_True},
}

const (
	VictimSystemType_r16_NavIC_r16_True = 0
)

var victimSystemTypeR16NavICR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_NavIC_r16_True},
}

const (
	VictimSystemType_r16_Wlan_r16_True = 0
)

var victimSystemTypeR16WlanR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_Wlan_r16_True},
}

const (
	VictimSystemType_r16_Bluetooth_r16_True = 0
)

var victimSystemTypeR16BluetoothR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_Bluetooth_r16_True},
}

const (
	VictimSystemType_r16_Ext_Uwb_r18_True = 0
)

var victimSystemTypeR16ExtUwbR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{VictimSystemType_r16_Ext_Uwb_r18_True},
}

type VictimSystemType_r16 struct {
	Gps_r16       *int64
	Glonass_r16   *int64
	Bds_r16       *int64
	Galileo_r16   *int64
	NavIC_r16     *int64
	Wlan_r16      *int64
	Bluetooth_r16 *int64
	Uwb_r18       *int64
}

func (ie *VictimSystemType_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(victimSystemTypeR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Uwb_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Gps_r16 != nil, ie.Glonass_r16 != nil, ie.Bds_r16 != nil, ie.Galileo_r16 != nil, ie.NavIC_r16 != nil, ie.Wlan_r16 != nil, ie.Bluetooth_r16 != nil}); err != nil {
		return err
	}

	// 3. gps-r16: enumerated
	{
		if ie.Gps_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Gps_r16, victimSystemTypeR16GpsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. glonass-r16: enumerated
	{
		if ie.Glonass_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Glonass_r16, victimSystemTypeR16GlonassR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. bds-r16: enumerated
	{
		if ie.Bds_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Bds_r16, victimSystemTypeR16BdsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. galileo-r16: enumerated
	{
		if ie.Galileo_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Galileo_r16, victimSystemTypeR16GalileoR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. navIC-r16: enumerated
	{
		if ie.NavIC_r16 != nil {
			if err := e.EncodeEnumerated(*ie.NavIC_r16, victimSystemTypeR16NavICR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. wlan-r16: enumerated
	{
		if ie.Wlan_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Wlan_r16, victimSystemTypeR16WlanR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. bluetooth-r16: enumerated
	{
		if ie.Bluetooth_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Bluetooth_r16, victimSystemTypeR16BluetoothR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "uwb-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Uwb_r18 != nil}); err != nil {
				return err
			}

			if ie.Uwb_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Uwb_r18, victimSystemTypeR16ExtUwbR18Constraints); err != nil {
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

func (ie *VictimSystemType_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(victimSystemTypeR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. gps-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(victimSystemTypeR16GpsR16Constraints)
			if err != nil {
				return err
			}
			ie.Gps_r16 = &idx
		}
	}

	// 4. glonass-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(victimSystemTypeR16GlonassR16Constraints)
			if err != nil {
				return err
			}
			ie.Glonass_r16 = &idx
		}
	}

	// 5. bds-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(victimSystemTypeR16BdsR16Constraints)
			if err != nil {
				return err
			}
			ie.Bds_r16 = &idx
		}
	}

	// 6. galileo-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(victimSystemTypeR16GalileoR16Constraints)
			if err != nil {
				return err
			}
			ie.Galileo_r16 = &idx
		}
	}

	// 7. navIC-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(victimSystemTypeR16NavICR16Constraints)
			if err != nil {
				return err
			}
			ie.NavIC_r16 = &idx
		}
	}

	// 8. wlan-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(victimSystemTypeR16WlanR16Constraints)
			if err != nil {
				return err
			}
			ie.Wlan_r16 = &idx
		}
	}

	// 9. bluetooth-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(victimSystemTypeR16BluetoothR16Constraints)
			if err != nil {
				return err
			}
			ie.Bluetooth_r16 = &idx
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
				{Name: "uwb-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(victimSystemTypeR16ExtUwbR18Constraints)
			if err != nil {
				return err
			}
			ie.Uwb_r18 = &v
		}
	}

	return nil
}
