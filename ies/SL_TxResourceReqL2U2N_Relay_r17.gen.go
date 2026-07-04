// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-TxResourceReqL2U2N-Relay-r17 (line 2228).

var sLTxResourceReqL2U2NRelayR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DestinationIdentityL2U2N-r17", Optional: true},
		{Name: "sl-TxInterestedFreqListL2U2N-r17"},
		{Name: "sl-TypeTxSyncListL2U2N-r17"},
		{Name: "sl-LocalID-Request-r17", Optional: true},
		{Name: "sl-PagingIdentityRemoteUE-r17", Optional: true},
		{Name: "sl-CapabilityInformationSidelink-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var sLTxResourceReqL2U2NRelayR17SlTypeTxSyncListL2U2NR17Constraints = per.SizeRange(1, common.MaxNrofFreqSL_r16)

const (
	SL_TxResourceReqL2U2N_Relay_r17_Sl_LocalID_Request_r17_True = 0
)

var sLTxResourceReqL2U2NRelayR17SlLocalIDRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_TxResourceReqL2U2N_Relay_r17_Sl_LocalID_Request_r17_True},
}

var sLTxResourceReqL2U2NRelayR17SlCapabilityInformationSidelinkR17Constraints = per.SizeConstraints{}

var sLTxResourceReqL2U2NRelayR17ExtSlPagingIdentityRemoteUEListR19Constraints = per.SizeRange(1, common.MaxNrofRemoteUE_r17)

type SL_TxResourceReqL2U2N_Relay_r17 struct {
	Sl_DestinationIdentityL2U2N_r17      *SL_DestinationIdentity_r16
	Sl_TxInterestedFreqListL2U2N_r17     SL_TxInterestedFreqList_r16
	Sl_TypeTxSyncListL2U2N_r17           []SL_TypeTxSync_r16
	Sl_LocalID_Request_r17               *int64
	Sl_PagingIdentityRemoteUE_r17        *SL_PagingIdentityRemoteUE_r17
	Sl_CapabilityInformationSidelink_r17 []byte
	Sl_PagingIdentityRemoteUEList_r19    []SL_PagingIdentityRemoteUE_r17
}

func (ie *SL_TxResourceReqL2U2N_Relay_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLTxResourceReqL2U2NRelayR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_PagingIdentityRemoteUEList_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DestinationIdentityL2U2N_r17 != nil, ie.Sl_LocalID_Request_r17 != nil, ie.Sl_PagingIdentityRemoteUE_r17 != nil, ie.Sl_CapabilityInformationSidelink_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DestinationIdentityL2U2N-r17: ref
	{
		if ie.Sl_DestinationIdentityL2U2N_r17 != nil {
			if err := ie.Sl_DestinationIdentityL2U2N_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-TxInterestedFreqListL2U2N-r17: ref
	{
		if err := ie.Sl_TxInterestedFreqListL2U2N_r17.Encode(e); err != nil {
			return err
		}
	}

	// 5. sl-TypeTxSyncListL2U2N-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sLTxResourceReqL2U2NRelayR17SlTypeTxSyncListL2U2NR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sl_TypeTxSyncListL2U2N_r17))); err != nil {
			return err
		}
		for i := range ie.Sl_TypeTxSyncListL2U2N_r17 {
			if err := ie.Sl_TypeTxSyncListL2U2N_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. sl-LocalID-Request-r17: enumerated
	{
		if ie.Sl_LocalID_Request_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_LocalID_Request_r17, sLTxResourceReqL2U2NRelayR17SlLocalIDRequestR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-PagingIdentityRemoteUE-r17: ref
	{
		if ie.Sl_PagingIdentityRemoteUE_r17 != nil {
			if err := ie.Sl_PagingIdentityRemoteUE_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. sl-CapabilityInformationSidelink-r17: octet-string
	{
		if ie.Sl_CapabilityInformationSidelink_r17 != nil {
			if err := e.EncodeOctetString(ie.Sl_CapabilityInformationSidelink_r17, sLTxResourceReqL2U2NRelayR17SlCapabilityInformationSidelinkR17Constraints); err != nil {
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
					{Name: "sl-PagingIdentityRemoteUEList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PagingIdentityRemoteUEList_r19 != nil}); err != nil {
				return err
			}

			if ie.Sl_PagingIdentityRemoteUEList_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(sLTxResourceReqL2U2NRelayR17ExtSlPagingIdentityRemoteUEListR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Sl_PagingIdentityRemoteUEList_r19))); err != nil {
					return err
				}
				for i := range ie.Sl_PagingIdentityRemoteUEList_r19 {
					if err := ie.Sl_PagingIdentityRemoteUEList_r19[i].Encode(ex); err != nil {
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

func (ie *SL_TxResourceReqL2U2N_Relay_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLTxResourceReqL2U2NRelayR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DestinationIdentityL2U2N-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_DestinationIdentityL2U2N_r17 = new(SL_DestinationIdentity_r16)
			if err := ie.Sl_DestinationIdentityL2U2N_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-TxInterestedFreqListL2U2N-r17: ref
	{
		if err := ie.Sl_TxInterestedFreqListL2U2N_r17.Decode(d); err != nil {
			return err
		}
	}

	// 5. sl-TypeTxSyncListL2U2N-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sLTxResourceReqL2U2NRelayR17SlTypeTxSyncListL2U2NR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sl_TypeTxSyncListL2U2N_r17 = make([]SL_TypeTxSync_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Sl_TypeTxSyncListL2U2N_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. sl-LocalID-Request-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLTxResourceReqL2U2NRelayR17SlLocalIDRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_LocalID_Request_r17 = &idx
		}
	}

	// 7. sl-PagingIdentityRemoteUE-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Sl_PagingIdentityRemoteUE_r17 = new(SL_PagingIdentityRemoteUE_r17)
			if err := ie.Sl_PagingIdentityRemoteUE_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. sl-CapabilityInformationSidelink-r17: octet-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeOctetString(sLTxResourceReqL2U2NRelayR17SlCapabilityInformationSidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CapabilityInformationSidelink_r17 = v
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
				{Name: "sl-PagingIdentityRemoteUEList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(sLTxResourceReqL2U2NRelayR17ExtSlPagingIdentityRemoteUEListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PagingIdentityRemoteUEList_r19 = make([]SL_PagingIdentityRemoteUE_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PagingIdentityRemoteUEList_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
