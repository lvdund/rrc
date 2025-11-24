package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetup_IEs struct {
	RadioBearerConfig        RadioBearerConfig   `madatory`
	MasterCellGroup          []byte              `madatory`
	LateNonCriticalExtension *[]byte             `optional`
	NonCriticalExtension     *RRCSetup_v1700_IEs `optional`
}

func (ie *RRCSetup_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.RadioBearerConfig.Encode(w); err != nil {
		return utils.WrapError("Encode RadioBearerConfig", err)
	}
	if err = w.WriteOctetString(ie.MasterCellGroup, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString MasterCellGroup", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
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

func (ie *RRCSetup_IEs) Decode(r *uper.UperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.RadioBearerConfig.Decode(r); err != nil {
		return utils.WrapError("Decode RadioBearerConfig", err)
	}
	var tmp_os_MasterCellGroup []byte
	if tmp_os_MasterCellGroup, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString MasterCellGroup", err)
	}
	ie.MasterCellGroup = tmp_os_MasterCellGroup
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCSetup_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
