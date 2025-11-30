package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_IEs struct {
	RadioBearerConfig        *RadioBearerConfig            `optional`
	SecondaryCellGroup       *[]byte                       `optional`
	MeasConfig               *MeasConfig                   `optional`
	LateNonCriticalExtension *[]byte                       `optional`
	NonCriticalExtension     *RRCReconfiguration_v1530_IEs `optional`
}

func (ie *RRCReconfiguration_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RadioBearerConfig != nil, ie.SecondaryCellGroup != nil, ie.MeasConfig != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RadioBearerConfig != nil {
		if err = ie.RadioBearerConfig.Encode(w); err != nil {
			return utils.WrapError("Encode RadioBearerConfig", err)
		}
	}
	if ie.SecondaryCellGroup != nil {
		if err = w.WriteOctetString(*ie.SecondaryCellGroup, nil, false); err != nil {
			return utils.WrapError("Encode SecondaryCellGroup", err)
		}
	}
	if ie.MeasConfig != nil {
		if err = ie.MeasConfig.Encode(w); err != nil {
			return utils.WrapError("Encode MeasConfig", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_IEs) Decode(r *aper.AperReader) error {
	var err error
	var RadioBearerConfigPresent bool
	if RadioBearerConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SecondaryCellGroupPresent bool
	if SecondaryCellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasConfigPresent bool
	if MeasConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if RadioBearerConfigPresent {
		ie.RadioBearerConfig = new(RadioBearerConfig)
		if err = ie.RadioBearerConfig.Decode(r); err != nil {
			return utils.WrapError("Decode RadioBearerConfig", err)
		}
	}
	if SecondaryCellGroupPresent {
		var tmp_os_SecondaryCellGroup []byte
		if tmp_os_SecondaryCellGroup, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode SecondaryCellGroup", err)
		}
		ie.SecondaryCellGroup = &tmp_os_SecondaryCellGroup
	}
	if MeasConfigPresent {
		ie.MeasConfig = new(MeasConfig)
		if err = ie.MeasConfig.Decode(r); err != nil {
			return utils.WrapError("Decode MeasConfig", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfiguration_v1530_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
