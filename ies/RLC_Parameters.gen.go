// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLC-Parameters (line 24815).

var rLCParametersConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "am-WithShortSN", Optional: true},
		{Name: "um-WithShortSN", Optional: true},
		{Name: "um-WithLongSN", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

const (
	RLC_Parameters_Am_WithShortSN_Supported = 0
)

var rLCParametersAmWithShortSNConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Am_WithShortSN_Supported},
}

const (
	RLC_Parameters_Um_WithShortSN_Supported = 0
)

var rLCParametersUmWithShortSNConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Um_WithShortSN_Supported},
}

const (
	RLC_Parameters_Um_WithLongSN_Supported = 0
)

var rLCParametersUmWithLongSNConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Um_WithLongSN_Supported},
}

const (
	RLC_Parameters_Ext_ExtendedT_PollRetransmit_r16_Supported = 0
)

var rLCParametersExtExtendedTPollRetransmitR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_ExtendedT_PollRetransmit_r16_Supported},
}

const (
	RLC_Parameters_Ext_ExtendedT_StatusProhibit_r16_Supported = 0
)

var rLCParametersExtExtendedTStatusProhibitR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_ExtendedT_StatusProhibit_r16_Supported},
}

const (
	RLC_Parameters_Ext_Am_WithLongSN_RedCap_r17_Supported = 0
)

var rLCParametersExtAmWithLongSNRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_Am_WithLongSN_RedCap_r17_Supported},
}

const (
	RLC_Parameters_Ext_Am_WithLongSN_NCR_r18_Supported = 0
)

var rLCParametersExtAmWithLongSNNCRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_Am_WithLongSN_NCR_r18_Supported},
}

const (
	RLC_Parameters_Ext_RemainingTimeBasedRetransmission_r19_Supported = 0
)

var rLCParametersExtRemainingTimeBasedRetransmissionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_RemainingTimeBasedRetransmission_r19_Supported},
}

const (
	RLC_Parameters_Ext_RemainingTimeBasedPolling_r19_Supported = 0
)

var rLCParametersExtRemainingTimeBasedPollingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_RemainingTimeBasedPolling_r19_Supported},
}

const (
	RLC_Parameters_Ext_RxRLC_Discard_r19_Supported = 0
)

var rLCParametersExtRxRLCDiscardR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_RxRLC_Discard_r19_Supported},
}

const (
	RLC_Parameters_Ext_TxRLC_StopReTxDiscardedSDU_r19_Supported = 0
)

var rLCParametersExtTxRLCStopReTxDiscardedSDUR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLC_Parameters_Ext_TxRLC_StopReTxDiscardedSDU_r19_Supported},
}

type RLC_Parameters struct {
	Am_WithShortSN                       *int64
	Um_WithShortSN                       *int64
	Um_WithLongSN                        *int64
	ExtendedT_PollRetransmit_r16         *int64
	ExtendedT_StatusProhibit_r16         *int64
	Am_WithLongSN_RedCap_r17             *int64
	Am_WithLongSN_NCR_r18                *int64
	RemainingTimeBasedRetransmission_r19 *int64
	RemainingTimeBasedPolling_r19        *int64
	RxRLC_Discard_r19                    *int64
	TxRLC_StopReTxDiscardedSDU_r19       *int64
}

func (ie *RLC_Parameters) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLCParametersConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ExtendedT_PollRetransmit_r16 != nil || ie.ExtendedT_StatusProhibit_r16 != nil
	hasExtGroup1 := ie.Am_WithLongSN_RedCap_r17 != nil
	hasExtGroup2 := ie.Am_WithLongSN_NCR_r18 != nil
	hasExtGroup3 := ie.RemainingTimeBasedRetransmission_r19 != nil || ie.RemainingTimeBasedPolling_r19 != nil || ie.RxRLC_Discard_r19 != nil || ie.TxRLC_StopReTxDiscardedSDU_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Am_WithShortSN != nil, ie.Um_WithShortSN != nil, ie.Um_WithLongSN != nil}); err != nil {
		return err
	}

	// 3. am-WithShortSN: enumerated
	{
		if ie.Am_WithShortSN != nil {
			if err := e.EncodeEnumerated(*ie.Am_WithShortSN, rLCParametersAmWithShortSNConstraints); err != nil {
				return err
			}
		}
	}

	// 4. um-WithShortSN: enumerated
	{
		if ie.Um_WithShortSN != nil {
			if err := e.EncodeEnumerated(*ie.Um_WithShortSN, rLCParametersUmWithShortSNConstraints); err != nil {
				return err
			}
		}
	}

	// 5. um-WithLongSN: enumerated
	{
		if ie.Um_WithLongSN != nil {
			if err := e.EncodeEnumerated(*ie.Um_WithLongSN, rLCParametersUmWithLongSNConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "extendedT-PollRetransmit-r16", Optional: true},
					{Name: "extendedT-StatusProhibit-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ExtendedT_PollRetransmit_r16 != nil, ie.ExtendedT_StatusProhibit_r16 != nil}); err != nil {
				return err
			}

			if ie.ExtendedT_PollRetransmit_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedT_PollRetransmit_r16, rLCParametersExtExtendedTPollRetransmitR16Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedT_StatusProhibit_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedT_StatusProhibit_r16, rLCParametersExtExtendedTStatusProhibitR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "am-WithLongSN-RedCap-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Am_WithLongSN_RedCap_r17 != nil}); err != nil {
				return err
			}

			if ie.Am_WithLongSN_RedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Am_WithLongSN_RedCap_r17, rLCParametersExtAmWithLongSNRedCapR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "am-WithLongSN-NCR-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Am_WithLongSN_NCR_r18 != nil}); err != nil {
				return err
			}

			if ie.Am_WithLongSN_NCR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Am_WithLongSN_NCR_r18, rLCParametersExtAmWithLongSNNCRR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "remainingTimeBasedRetransmission-r19", Optional: true},
					{Name: "remainingTimeBasedPolling-r19", Optional: true},
					{Name: "rxRLC-Discard-r19", Optional: true},
					{Name: "txRLC-StopReTxDiscardedSDU-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RemainingTimeBasedRetransmission_r19 != nil, ie.RemainingTimeBasedPolling_r19 != nil, ie.RxRLC_Discard_r19 != nil, ie.TxRLC_StopReTxDiscardedSDU_r19 != nil}); err != nil {
				return err
			}

			if ie.RemainingTimeBasedRetransmission_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.RemainingTimeBasedRetransmission_r19, rLCParametersExtRemainingTimeBasedRetransmissionR19Constraints); err != nil {
					return err
				}
			}

			if ie.RemainingTimeBasedPolling_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.RemainingTimeBasedPolling_r19, rLCParametersExtRemainingTimeBasedPollingR19Constraints); err != nil {
					return err
				}
			}

			if ie.RxRLC_Discard_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.RxRLC_Discard_r19, rLCParametersExtRxRLCDiscardR19Constraints); err != nil {
					return err
				}
			}

			if ie.TxRLC_StopReTxDiscardedSDU_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TxRLC_StopReTxDiscardedSDU_r19, rLCParametersExtTxRLCStopReTxDiscardedSDUR19Constraints); err != nil {
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

func (ie *RLC_Parameters) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLCParametersConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. am-WithShortSN: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rLCParametersAmWithShortSNConstraints)
			if err != nil {
				return err
			}
			ie.Am_WithShortSN = &idx
		}
	}

	// 4. um-WithShortSN: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(rLCParametersUmWithShortSNConstraints)
			if err != nil {
				return err
			}
			ie.Um_WithShortSN = &idx
		}
	}

	// 5. um-WithLongSN: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rLCParametersUmWithLongSNConstraints)
			if err != nil {
				return err
			}
			ie.Um_WithLongSN = &idx
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
				{Name: "extendedT-PollRetransmit-r16", Optional: true},
				{Name: "extendedT-StatusProhibit-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rLCParametersExtExtendedTPollRetransmitR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedT_PollRetransmit_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(rLCParametersExtExtendedTStatusProhibitR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedT_StatusProhibit_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "am-WithLongSN-RedCap-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rLCParametersExtAmWithLongSNRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.Am_WithLongSN_RedCap_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "am-WithLongSN-NCR-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rLCParametersExtAmWithLongSNNCRR18Constraints)
			if err != nil {
				return err
			}
			ie.Am_WithLongSN_NCR_r18 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "remainingTimeBasedRetransmission-r19", Optional: true},
				{Name: "remainingTimeBasedPolling-r19", Optional: true},
				{Name: "rxRLC-Discard-r19", Optional: true},
				{Name: "txRLC-StopReTxDiscardedSDU-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rLCParametersExtRemainingTimeBasedRetransmissionR19Constraints)
			if err != nil {
				return err
			}
			ie.RemainingTimeBasedRetransmission_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(rLCParametersExtRemainingTimeBasedPollingR19Constraints)
			if err != nil {
				return err
			}
			ie.RemainingTimeBasedPolling_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(rLCParametersExtRxRLCDiscardR19Constraints)
			if err != nil {
				return err
			}
			ie.RxRLC_Discard_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(rLCParametersExtTxRLCStopReTxDiscardedSDUR19Constraints)
			if err != nil {
				return err
			}
			ie.TxRLC_StopReTxDiscardedSDU_r19 = &v
		}
	}

	return nil
}
