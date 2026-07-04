// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxResourceReqDisc-r17 (line 2203).

var sLTxResourceReqDiscR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIdentityDisc-r17"},
		{Name: "sl-SourceIdentityRelayUE-r17", Optional: true},
		{Name: "sl-CastTypeDisc-r17"},
		{Name: "sl-TxInterestedFreqListDisc-r17"},
		{Name: "sl-TypeTxSyncListDisc-r17"},
		{Name: "sl-DiscoveryType-r17"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Broadcast = 0
	SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Groupcast = 1
	SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Unicast   = 2
	SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Spare1    = 3
)

var sLTxResourceReqDiscR17SlCastTypeDiscR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Broadcast, SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Groupcast, SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Unicast, SL_TxResourceReqDisc_r17_Sl_CastTypeDisc_r17_Spare1},
}

var sLTxResourceReqDiscR17SlTypeTxSyncListDiscR17Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

const (
	SL_TxResourceReqDisc_r17_Sl_DiscoveryType_r17_Relay     = 0
	SL_TxResourceReqDisc_r17_Sl_DiscoveryType_r17_Non_Relay = 1
)

var sLTxResourceReqDiscR17SlDiscoveryTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxResourceReqDisc_r17_Sl_DiscoveryType_r17_Relay, SL_TxResourceReqDisc_r17_Sl_DiscoveryType_r17_Non_Relay},
}

const (
	SL_TxResourceReqDisc_r17_Ext_Ue_TypeU2U_r18_RelayUE  = 0
	SL_TxResourceReqDisc_r17_Ext_Ue_TypeU2U_r18_RemoteUE = 1
)

var sLTxResourceReqDiscR17ExtUeTypeU2UR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxResourceReqDisc_r17_Ext_Ue_TypeU2U_r18_RelayUE, SL_TxResourceReqDisc_r17_Ext_Ue_TypeU2U_r18_RemoteUE},
}

type SL_TxResourceReqDisc_r17 struct {
	Sl_DestinationIdentityDisc_r17  SL_DestinationIdentity_r16
	Sl_SourceIdentityRelayUE_r17    *SL_SourceIdentity_r17
	Sl_CastTypeDisc_r17             int64
	Sl_TxInterestedFreqListDisc_r17 SL_TxInterestedFreqList_r16
	Sl_TypeTxSyncListDisc_r17       []SL_TypeTxSync_r16
	Sl_DiscoveryType_r17            int64
	Ue_TypeU2U_r18                  *int64
}

func (ie *SL_TxResourceReqDisc_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxResourceReqDiscR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ue_TypeU2U_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_SourceIdentityRelayUE_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DestinationIdentityDisc-r17: ref
	{
		if err := ie.Sl_DestinationIdentityDisc_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-SourceIdentityRelayUE-r17: ref
	{
		if ie.Sl_SourceIdentityRelayUE_r17 != nil {
			if err := ie.Sl_SourceIdentityRelayUE_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-CastTypeDisc-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_CastTypeDisc_r17, sLTxResourceReqDiscR17SlCastTypeDiscR17Constraints); err != nil {
			return err
		}
	}

	// 6. sl-TxInterestedFreqListDisc-r17: ref
	{
		if err := ie.Sl_TxInterestedFreqListDisc_r17.Encode(e); err != nil {
			return err
		}
	}

	// 7. sl-TypeTxSyncListDisc-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLTxResourceReqDiscR17SlTypeTxSyncListDiscR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_TypeTxSyncListDisc_r17))); err != nil {
			return err
		}
		for i := range ie.Sl_TypeTxSyncListDisc_r17 {
			if err := ie.Sl_TypeTxSyncListDisc_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. sl-DiscoveryType-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_DiscoveryType_r17, sLTxResourceReqDiscR17SlDiscoveryTypeR17Constraints); err != nil {
			return err
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
					{Name: "ue-TypeU2U-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ue_TypeU2U_r18 != nil}); err != nil {
				return err
			}

			if ie.Ue_TypeU2U_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ue_TypeU2U_r18, sLTxResourceReqDiscR17ExtUeTypeU2UR18Constraints); err != nil {
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

func (ie *SL_TxResourceReqDisc_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxResourceReqDiscR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DestinationIdentityDisc-r17: ref
	{
		if err := ie.Sl_DestinationIdentityDisc_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-SourceIdentityRelayUE-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_SourceIdentityRelayUE_r17 = new(SL_SourceIdentity_r17)
			if err := ie.Sl_SourceIdentityRelayUE_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-CastTypeDisc-r17: enumerated
	{
		v2, err := d.DecodeEnumerated(sLTxResourceReqDiscR17SlCastTypeDiscR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_CastTypeDisc_r17 = v2
	}

	// 6. sl-TxInterestedFreqListDisc-r17: ref
	{
		if err := ie.Sl_TxInterestedFreqListDisc_r17.Decode(d); err != nil {
			return err
		}
	}

	// 7. sl-TypeTxSyncListDisc-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLTxResourceReqDiscR17SlTypeTxSyncListDiscR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_TypeTxSyncListDisc_r17 = make([]SL_TypeTxSync_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_TypeTxSyncListDisc_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. sl-DiscoveryType-r17: enumerated
	{
		v5, err := d.DecodeEnumerated(sLTxResourceReqDiscR17SlDiscoveryTypeR17Constraints)
		if err != nil {
			return err
		}
		ie.Sl_DiscoveryType_r17 = v5
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
				{Name: "ue-TypeU2U-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sLTxResourceReqDiscR17ExtUeTypeU2UR18Constraints)
			if err != nil {
				return err
			}
			ie.Ue_TypeU2U_r18 = &v
		}
	}

	return nil
}
