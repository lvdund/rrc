package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VictimSystemType_r16 struct {
	Gps_r16       *VictimSystemType_r16_gps_r16       `optional`
	Glonass_r16   *VictimSystemType_r16_glonass_r16   `optional`
	Bds_r16       *VictimSystemType_r16_bds_r16       `optional`
	Galileo_r16   *VictimSystemType_r16_galileo_r16   `optional`
	NavIC_r16     *VictimSystemType_r16_navIC_r16     `optional`
	Wlan_r16      *VictimSystemType_r16_wlan_r16      `optional`
	Bluetooth_r16 *VictimSystemType_r16_bluetooth_r16 `optional`
}

func (ie *VictimSystemType_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Gps_r16 != nil, ie.Glonass_r16 != nil, ie.Bds_r16 != nil, ie.Galileo_r16 != nil, ie.NavIC_r16 != nil, ie.Wlan_r16 != nil, ie.Bluetooth_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Gps_r16 != nil {
		if err = ie.Gps_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gps_r16", err)
		}
	}
	if ie.Glonass_r16 != nil {
		if err = ie.Glonass_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Glonass_r16", err)
		}
	}
	if ie.Bds_r16 != nil {
		if err = ie.Bds_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bds_r16", err)
		}
	}
	if ie.Galileo_r16 != nil {
		if err = ie.Galileo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Galileo_r16", err)
		}
	}
	if ie.NavIC_r16 != nil {
		if err = ie.NavIC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NavIC_r16", err)
		}
	}
	if ie.Wlan_r16 != nil {
		if err = ie.Wlan_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Wlan_r16", err)
		}
	}
	if ie.Bluetooth_r16 != nil {
		if err = ie.Bluetooth_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bluetooth_r16", err)
		}
	}
	return nil
}

func (ie *VictimSystemType_r16) Decode(r *uper.UperReader) error {
	var err error
	var Gps_r16Present bool
	if Gps_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Glonass_r16Present bool
	if Glonass_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bds_r16Present bool
	if Bds_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Galileo_r16Present bool
	if Galileo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NavIC_r16Present bool
	if NavIC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Wlan_r16Present bool
	if Wlan_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bluetooth_r16Present bool
	if Bluetooth_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Gps_r16Present {
		ie.Gps_r16 = new(VictimSystemType_r16_gps_r16)
		if err = ie.Gps_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gps_r16", err)
		}
	}
	if Glonass_r16Present {
		ie.Glonass_r16 = new(VictimSystemType_r16_glonass_r16)
		if err = ie.Glonass_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Glonass_r16", err)
		}
	}
	if Bds_r16Present {
		ie.Bds_r16 = new(VictimSystemType_r16_bds_r16)
		if err = ie.Bds_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bds_r16", err)
		}
	}
	if Galileo_r16Present {
		ie.Galileo_r16 = new(VictimSystemType_r16_galileo_r16)
		if err = ie.Galileo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Galileo_r16", err)
		}
	}
	if NavIC_r16Present {
		ie.NavIC_r16 = new(VictimSystemType_r16_navIC_r16)
		if err = ie.NavIC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NavIC_r16", err)
		}
	}
	if Wlan_r16Present {
		ie.Wlan_r16 = new(VictimSystemType_r16_wlan_r16)
		if err = ie.Wlan_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Wlan_r16", err)
		}
	}
	if Bluetooth_r16Present {
		ie.Bluetooth_r16 = new(VictimSystemType_r16_bluetooth_r16)
		if err = ie.Bluetooth_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bluetooth_r16", err)
		}
	}
	return nil
}
