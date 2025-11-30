package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type VictimSystemType struct {
	Gps       *VictimSystemType_gps       `optional`
	Glonass   *VictimSystemType_glonass   `optional`
	Bds       *VictimSystemType_bds       `optional`
	Galileo   *VictimSystemType_galileo   `optional`
	Wlan      *VictimSystemType_wlan      `optional`
	Bluetooth *VictimSystemType_bluetooth `optional`
}

func (ie *VictimSystemType) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Gps != nil, ie.Glonass != nil, ie.Bds != nil, ie.Galileo != nil, ie.Wlan != nil, ie.Bluetooth != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Gps != nil {
		if err = ie.Gps.Encode(w); err != nil {
			return utils.WrapError("Encode Gps", err)
		}
	}
	if ie.Glonass != nil {
		if err = ie.Glonass.Encode(w); err != nil {
			return utils.WrapError("Encode Glonass", err)
		}
	}
	if ie.Bds != nil {
		if err = ie.Bds.Encode(w); err != nil {
			return utils.WrapError("Encode Bds", err)
		}
	}
	if ie.Galileo != nil {
		if err = ie.Galileo.Encode(w); err != nil {
			return utils.WrapError("Encode Galileo", err)
		}
	}
	if ie.Wlan != nil {
		if err = ie.Wlan.Encode(w); err != nil {
			return utils.WrapError("Encode Wlan", err)
		}
	}
	if ie.Bluetooth != nil {
		if err = ie.Bluetooth.Encode(w); err != nil {
			return utils.WrapError("Encode Bluetooth", err)
		}
	}
	return nil
}

func (ie *VictimSystemType) Decode(r *aper.AperReader) error {
	var err error
	var GpsPresent bool
	if GpsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var GlonassPresent bool
	if GlonassPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BdsPresent bool
	if BdsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var GalileoPresent bool
	if GalileoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var WlanPresent bool
	if WlanPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BluetoothPresent bool
	if BluetoothPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if GpsPresent {
		ie.Gps = new(VictimSystemType_gps)
		if err = ie.Gps.Decode(r); err != nil {
			return utils.WrapError("Decode Gps", err)
		}
	}
	if GlonassPresent {
		ie.Glonass = new(VictimSystemType_glonass)
		if err = ie.Glonass.Decode(r); err != nil {
			return utils.WrapError("Decode Glonass", err)
		}
	}
	if BdsPresent {
		ie.Bds = new(VictimSystemType_bds)
		if err = ie.Bds.Decode(r); err != nil {
			return utils.WrapError("Decode Bds", err)
		}
	}
	if GalileoPresent {
		ie.Galileo = new(VictimSystemType_galileo)
		if err = ie.Galileo.Decode(r); err != nil {
			return utils.WrapError("Decode Galileo", err)
		}
	}
	if WlanPresent {
		ie.Wlan = new(VictimSystemType_wlan)
		if err = ie.Wlan.Decode(r); err != nil {
			return utils.WrapError("Decode Wlan", err)
		}
	}
	if BluetoothPresent {
		ie.Bluetooth = new(VictimSystemType_bluetooth)
		if err = ie.Bluetooth.Decode(r); err != nil {
			return utils.WrapError("Decode Bluetooth", err)
		}
	}
	return nil
}
