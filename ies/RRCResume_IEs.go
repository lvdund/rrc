package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCResume_IEs struct {
	RadioBearerConfig        *RadioBearerConfig        `optional`
	MasterCellGroup          *[]byte                   `optional`
	MeasConfig               *MeasConfig               `optional`
	FullConfig               *RRCResume_IEs_fullConfig `optional`
	LateNonCriticalExtension *[]byte                   `optional`
	NonCriticalExtension     *RRCResume_v1560_IEs      `optional`
}

func (ie *RRCResume_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RadioBearerConfig != nil, ie.MasterCellGroup != nil, ie.MeasConfig != nil, ie.FullConfig != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
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
	if ie.MasterCellGroup != nil {
		if err = w.WriteOctetString(*ie.MasterCellGroup, nil, false); err != nil {
			return utils.WrapError("Encode MasterCellGroup", err)
		}
	}
	if ie.MeasConfig != nil {
		if err = ie.MeasConfig.Encode(w); err != nil {
			return utils.WrapError("Encode MeasConfig", err)
		}
	}
	if ie.FullConfig != nil {
		if err = ie.FullConfig.Encode(w); err != nil {
			return utils.WrapError("Encode FullConfig", err)
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

func (ie *RRCResume_IEs) Decode(r *aper.AperReader) error {
	var err error
	var RadioBearerConfigPresent bool
	if RadioBearerConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MasterCellGroupPresent bool
	if MasterCellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasConfigPresent bool
	if MeasConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FullConfigPresent bool
	if FullConfigPresent, err = r.ReadBool(); err != nil {
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
	if MasterCellGroupPresent {
		var tmp_os_MasterCellGroup []byte
		if tmp_os_MasterCellGroup, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode MasterCellGroup", err)
		}
		ie.MasterCellGroup = &tmp_os_MasterCellGroup
	}
	if MeasConfigPresent {
		ie.MeasConfig = new(MeasConfig)
		if err = ie.MeasConfig.Decode(r); err != nil {
			return utils.WrapError("Decode MeasConfig", err)
		}
	}
	if FullConfigPresent {
		ie.FullConfig = new(RRCResume_IEs_fullConfig)
		if err = ie.FullConfig.Decode(r); err != nil {
			return utils.WrapError("Decode FullConfig", err)
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
		ie.NonCriticalExtension = new(RRCResume_v1560_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
