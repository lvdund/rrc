package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasResultWLAN_r16 struct {
	Wlan_Identifiers_r16 WLAN_Identifiers_r16 `madatory`
	RssiWLAN_r16         *WLAN_RSSI_Range_r16 `optional`
	Rtt_WLAN_r16         *WLAN_RTT_r16        `optional`
}

func (ie *LogMeasResultWLAN_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RssiWLAN_r16 != nil, ie.Rtt_WLAN_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Wlan_Identifiers_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Wlan_Identifiers_r16", err)
	}
	if ie.RssiWLAN_r16 != nil {
		if err = ie.RssiWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RssiWLAN_r16", err)
		}
	}
	if ie.Rtt_WLAN_r16 != nil {
		if err = ie.Rtt_WLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rtt_WLAN_r16", err)
		}
	}
	return nil
}

func (ie *LogMeasResultWLAN_r16) Decode(r *uper.UperReader) error {
	var err error
	var RssiWLAN_r16Present bool
	if RssiWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rtt_WLAN_r16Present bool
	if Rtt_WLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Wlan_Identifiers_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Wlan_Identifiers_r16", err)
	}
	if RssiWLAN_r16Present {
		ie.RssiWLAN_r16 = new(WLAN_RSSI_Range_r16)
		if err = ie.RssiWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RssiWLAN_r16", err)
		}
	}
	if Rtt_WLAN_r16Present {
		ie.Rtt_WLAN_r16 = new(WLAN_RTT_r16)
		if err = ie.Rtt_WLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rtt_WLAN_r16", err)
		}
	}
	return nil
}
