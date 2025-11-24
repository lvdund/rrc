package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type WLAN_Identifiers_r16 struct {
	Ssid_r16   *[]byte `lb:1,ub:32,optional`
	Bssid_r16  *[]byte `lb:6,ub:6,optional`
	Hessid_r16 *[]byte `lb:6,ub:6,optional`
}

func (ie *WLAN_Identifiers_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssid_r16 != nil, ie.Bssid_r16 != nil, ie.Hessid_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ssid_r16 != nil {
		if err = w.WriteOctetString(*ie.Ssid_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Ssid_r16", err)
		}
	}
	if ie.Bssid_r16 != nil {
		if err = w.WriteOctetString(*ie.Bssid_r16, &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode Bssid_r16", err)
		}
	}
	if ie.Hessid_r16 != nil {
		if err = w.WriteOctetString(*ie.Hessid_r16, &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode Hessid_r16", err)
		}
	}
	return nil
}

func (ie *WLAN_Identifiers_r16) Decode(r *uper.UperReader) error {
	var err error
	var Ssid_r16Present bool
	if Ssid_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bssid_r16Present bool
	if Bssid_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Hessid_r16Present bool
	if Hessid_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ssid_r16Present {
		var tmp_os_Ssid_r16 []byte
		if tmp_os_Ssid_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Ssid_r16", err)
		}
		ie.Ssid_r16 = &tmp_os_Ssid_r16
	}
	if Bssid_r16Present {
		var tmp_os_Bssid_r16 []byte
		if tmp_os_Bssid_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode Bssid_r16", err)
		}
		ie.Bssid_r16 = &tmp_os_Bssid_r16
	}
	if Hessid_r16Present {
		var tmp_os_Hessid_r16 []byte
		if tmp_os_Hessid_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode Hessid_r16", err)
		}
		ie.Hessid_r16 = &tmp_os_Hessid_r16
	}
	return nil
}
