// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RAN-VisibleMeasurements-r17 (line 781).

var rANVisibleMeasurementsR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "appLayerBufferLevelList-r17", Optional: true},
		{Name: "playoutDelayForMediaStartup-r17", Optional: true},
		{Name: "pdu-SessionIdList-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var rANVisibleMeasurementsR17AppLayerBufferLevelListR17Constraints = per.SizeRange(1, 8)

var rANVisibleMeasurementsR17PlayoutDelayForMediaStartupR17Constraints = per.Constrained(0, 30000)

var rANVisibleMeasurementsR17PduSessionIdListR17Constraints = per.SizeRange(1, common.MaxNrofPDU_Sessions_r17)

var rANVisibleMeasurementsR17ExtPduSessionIdListExtV1800Constraints = per.SizeRange(1, common.MaxNrofPDU_Sessions_r17)

type RAN_VisibleMeasurements_r17 struct {
	AppLayerBufferLevelList_r17     []AppLayerBufferLevel_r17
	PlayoutDelayForMediaStartup_r17 *int64
	Pdu_SessionIdList_r17           []PDU_SessionID
	Pdu_SessionIdListExt_v1800      []QFI_List_r18
}

func (ie *RAN_VisibleMeasurements_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rANVisibleMeasurementsR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Pdu_SessionIdListExt_v1800 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AppLayerBufferLevelList_r17 != nil, ie.PlayoutDelayForMediaStartup_r17 != nil, ie.Pdu_SessionIdList_r17 != nil}); err != nil {
		return err
	}

	// 3. appLayerBufferLevelList-r17: sequence-of
	{
		if ie.AppLayerBufferLevelList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(rANVisibleMeasurementsR17AppLayerBufferLevelListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AppLayerBufferLevelList_r17))); err != nil {
				return err
			}
			for i := range ie.AppLayerBufferLevelList_r17 {
				if err := ie.AppLayerBufferLevelList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. playoutDelayForMediaStartup-r17: integer
	{
		if ie.PlayoutDelayForMediaStartup_r17 != nil {
			if err := e.EncodeInteger(*ie.PlayoutDelayForMediaStartup_r17, rANVisibleMeasurementsR17PlayoutDelayForMediaStartupR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. pdu-SessionIdList-r17: sequence-of
	{
		if ie.Pdu_SessionIdList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(rANVisibleMeasurementsR17PduSessionIdListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdu_SessionIdList_r17))); err != nil {
				return err
			}
			for i := range ie.Pdu_SessionIdList_r17 {
				if err := ie.Pdu_SessionIdList_r17[i].Encode(e); err != nil {
					return err
				}
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
					{Name: "pdu-SessionIdListExt-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdu_SessionIdListExt_v1800 != nil}); err != nil {
				return err
			}

			if ie.Pdu_SessionIdListExt_v1800 != nil {
				seqOf := ex.NewSequenceOfEncoder(rANVisibleMeasurementsR17ExtPduSessionIdListExtV1800Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pdu_SessionIdListExt_v1800))); err != nil {
					return err
				}
				for i := range ie.Pdu_SessionIdListExt_v1800 {
					if err := ie.Pdu_SessionIdListExt_v1800[i].Encode(ex); err != nil {
						return err
					}
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

func (ie *RAN_VisibleMeasurements_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rANVisibleMeasurementsR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. appLayerBufferLevelList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(rANVisibleMeasurementsR17AppLayerBufferLevelListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AppLayerBufferLevelList_r17 = make([]AppLayerBufferLevel_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AppLayerBufferLevelList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. playoutDelayForMediaStartup-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(rANVisibleMeasurementsR17PlayoutDelayForMediaStartupR17Constraints)
			if err != nil {
				return err
			}
			ie.PlayoutDelayForMediaStartup_r17 = &v
		}
	}

	// 5. pdu-SessionIdList-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(rANVisibleMeasurementsR17PduSessionIdListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdu_SessionIdList_r17 = make([]PDU_SessionID, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdu_SessionIdList_r17[i].Decode(d); err != nil {
					return err
				}
			}
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
				{Name: "pdu-SessionIdListExt-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(rANVisibleMeasurementsR17ExtPduSessionIdListExtV1800Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdu_SessionIdListExt_v1800 = make([]QFI_List_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdu_SessionIdListExt_v1800[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
