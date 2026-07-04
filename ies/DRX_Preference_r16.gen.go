// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DRX-Preference-r16 (line 2532).

var dRXPreferenceR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "preferredDRX-InactivityTimer-r16", Optional: true},
		{Name: "preferredDRX-LongCycle-r16", Optional: true},
		{Name: "preferredDRX-ShortCycle-r16", Optional: true},
		{Name: "preferredDRX-ShortCycleTimer-r16", Optional: true},
	},
}

const (
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms0    = 0
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms1    = 1
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms2    = 2
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms3    = 3
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms4    = 4
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms5    = 5
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms6    = 6
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms8    = 7
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms10   = 8
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms20   = 9
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms30   = 10
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms40   = 11
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms50   = 12
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms60   = 13
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms80   = 14
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms100  = 15
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms200  = 16
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms300  = 17
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms500  = 18
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms750  = 19
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms1280 = 20
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms1920 = 21
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms2560 = 22
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare9 = 23
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare8 = 24
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare7 = 25
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare6 = 26
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare5 = 27
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare4 = 28
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare3 = 29
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare2 = 30
	DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare1 = 31
)

var dRXPreferenceR16PreferredDRXInactivityTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms0, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms1, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms2, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms3, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms4, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms5, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms6, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms8, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms10, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms20, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms30, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms40, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms50, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms60, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms80, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms100, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms200, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms300, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms500, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms750, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms1280, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms1920, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Ms2560, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare9, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare8, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare7, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare6, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare5, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare4, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare3, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare2, DRX_Preference_r16_PreferredDRX_InactivityTimer_r16_Spare1},
}

const (
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms10    = 0
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms20    = 1
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms32    = 2
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms40    = 3
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms60    = 4
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms64    = 5
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms70    = 6
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms80    = 7
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms128   = 8
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms160   = 9
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms256   = 10
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms320   = 11
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms512   = 12
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms640   = 13
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms1024  = 14
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms1280  = 15
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms2048  = 16
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms2560  = 17
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms5120  = 18
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms10240 = 19
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare12 = 20
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare11 = 21
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare10 = 22
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare9  = 23
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare8  = 24
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare7  = 25
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare6  = 26
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare5  = 27
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare4  = 28
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare3  = 29
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare2  = 30
	DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare1  = 31
)

var dRXPreferenceR16PreferredDRXLongCycleR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms10, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms20, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms32, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms40, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms60, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms64, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms70, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms80, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms128, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms160, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms256, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms320, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms512, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms640, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms1024, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms1280, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms2048, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms2560, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms5120, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Ms10240, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare12, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare11, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare10, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare9, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare8, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare7, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare6, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare5, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare4, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare3, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare2, DRX_Preference_r16_PreferredDRX_LongCycle_r16_Spare1},
}

const (
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms2    = 0
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms3    = 1
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms4    = 2
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms5    = 3
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms6    = 4
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms7    = 5
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms8    = 6
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms10   = 7
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms14   = 8
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms16   = 9
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms20   = 10
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms30   = 11
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms32   = 12
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms35   = 13
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms40   = 14
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms64   = 15
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms80   = 16
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms128  = 17
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms160  = 18
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms256  = 19
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms320  = 20
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms512  = 21
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms640  = 22
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare9 = 23
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare8 = 24
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare7 = 25
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare6 = 26
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare5 = 27
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare4 = 28
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare3 = 29
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare2 = 30
	DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare1 = 31
)

var dRXPreferenceR16PreferredDRXShortCycleR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms2, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms3, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms4, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms5, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms6, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms7, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms8, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms10, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms14, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms16, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms20, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms30, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms32, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms35, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms40, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms64, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms80, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms128, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms160, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms256, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms320, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms512, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Ms640, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare9, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare8, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare7, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare6, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare5, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare4, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare3, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare2, DRX_Preference_r16_PreferredDRX_ShortCycle_r16_Spare1},
}

var dRXPreferenceR16PreferredDRXShortCycleTimerR16Constraints = per.Constrained(1, 16)

type DRX_Preference_r16 struct {
	PreferredDRX_InactivityTimer_r16 *int64
	PreferredDRX_LongCycle_r16       *int64
	PreferredDRX_ShortCycle_r16      *int64
	PreferredDRX_ShortCycleTimer_r16 *int64
}

func (ie *DRX_Preference_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dRXPreferenceR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PreferredDRX_InactivityTimer_r16 != nil, ie.PreferredDRX_LongCycle_r16 != nil, ie.PreferredDRX_ShortCycle_r16 != nil, ie.PreferredDRX_ShortCycleTimer_r16 != nil}); err != nil {
		return err
	}

	// 2. preferredDRX-InactivityTimer-r16: enumerated
	{
		if ie.PreferredDRX_InactivityTimer_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PreferredDRX_InactivityTimer_r16, dRXPreferenceR16PreferredDRXInactivityTimerR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. preferredDRX-LongCycle-r16: enumerated
	{
		if ie.PreferredDRX_LongCycle_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PreferredDRX_LongCycle_r16, dRXPreferenceR16PreferredDRXLongCycleR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. preferredDRX-ShortCycle-r16: enumerated
	{
		if ie.PreferredDRX_ShortCycle_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PreferredDRX_ShortCycle_r16, dRXPreferenceR16PreferredDRXShortCycleR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. preferredDRX-ShortCycleTimer-r16: integer
	{
		if ie.PreferredDRX_ShortCycleTimer_r16 != nil {
			if err := e.EncodeInteger(*ie.PreferredDRX_ShortCycleTimer_r16, dRXPreferenceR16PreferredDRXShortCycleTimerR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DRX_Preference_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dRXPreferenceR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. preferredDRX-InactivityTimer-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(dRXPreferenceR16PreferredDRXInactivityTimerR16Constraints)
			if err != nil {
				return err
			}
			ie.PreferredDRX_InactivityTimer_r16 = &idx
		}
	}

	// 3. preferredDRX-LongCycle-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(dRXPreferenceR16PreferredDRXLongCycleR16Constraints)
			if err != nil {
				return err
			}
			ie.PreferredDRX_LongCycle_r16 = &idx
		}
	}

	// 4. preferredDRX-ShortCycle-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dRXPreferenceR16PreferredDRXShortCycleR16Constraints)
			if err != nil {
				return err
			}
			ie.PreferredDRX_ShortCycle_r16 = &idx
		}
	}

	// 5. preferredDRX-ShortCycleTimer-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(dRXPreferenceR16PreferredDRXShortCycleTimerR16Constraints)
			if err != nil {
				return err
			}
			ie.PreferredDRX_ShortCycleTimer_r16 = &v
		}
	}

	return nil
}
