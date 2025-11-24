package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NZP_CSI_RS_Resource struct {
	Nzp_CSI_RS_ResourceId  NZP_CSI_RS_ResourceId                     `madatory`
	ResourceMapping        CSI_RS_ResourceMapping                    `madatory`
	PowerControlOffset     int64                                     `lb:-8,ub:15,madatory`
	PowerControlOffsetSS   *NZP_CSI_RS_Resource_powerControlOffsetSS `optional`
	ScramblingID           ScramblingId                              `madatory`
	PeriodicityAndOffset   *CSI_ResourcePeriodicityAndOffset         `optional`
	Qcl_InfoPeriodicCSI_RS *TCI_StateId                              `optional`
}

func (ie *NZP_CSI_RS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PowerControlOffsetSS != nil, ie.PeriodicityAndOffset != nil, ie.Qcl_InfoPeriodicCSI_RS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Nzp_CSI_RS_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode Nzp_CSI_RS_ResourceId", err)
	}
	if err = ie.ResourceMapping.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceMapping", err)
	}
	if err = w.WriteInteger(ie.PowerControlOffset, &uper.Constraint{Lb: -8, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger PowerControlOffset", err)
	}
	if ie.PowerControlOffsetSS != nil {
		if err = ie.PowerControlOffsetSS.Encode(w); err != nil {
			return utils.WrapError("Encode PowerControlOffsetSS", err)
		}
	}
	if err = ie.ScramblingID.Encode(w); err != nil {
		return utils.WrapError("Encode ScramblingID", err)
	}
	if ie.PeriodicityAndOffset != nil {
		if err = ie.PeriodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicityAndOffset", err)
		}
	}
	if ie.Qcl_InfoPeriodicCSI_RS != nil {
		if err = ie.Qcl_InfoPeriodicCSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode Qcl_InfoPeriodicCSI_RS", err)
		}
	}
	return nil
}

func (ie *NZP_CSI_RS_Resource) Decode(r *uper.UperReader) error {
	var err error
	var PowerControlOffsetSSPresent bool
	if PowerControlOffsetSSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PeriodicityAndOffsetPresent bool
	if PeriodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Qcl_InfoPeriodicCSI_RSPresent bool
	if Qcl_InfoPeriodicCSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Nzp_CSI_RS_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode Nzp_CSI_RS_ResourceId", err)
	}
	if err = ie.ResourceMapping.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceMapping", err)
	}
	var tmp_int_PowerControlOffset int64
	if tmp_int_PowerControlOffset, err = r.ReadInteger(&uper.Constraint{Lb: -8, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger PowerControlOffset", err)
	}
	ie.PowerControlOffset = tmp_int_PowerControlOffset
	if PowerControlOffsetSSPresent {
		ie.PowerControlOffsetSS = new(NZP_CSI_RS_Resource_powerControlOffsetSS)
		if err = ie.PowerControlOffsetSS.Decode(r); err != nil {
			return utils.WrapError("Decode PowerControlOffsetSS", err)
		}
	}
	if err = ie.ScramblingID.Decode(r); err != nil {
		return utils.WrapError("Decode ScramblingID", err)
	}
	if PeriodicityAndOffsetPresent {
		ie.PeriodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
		if err = ie.PeriodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicityAndOffset", err)
		}
	}
	if Qcl_InfoPeriodicCSI_RSPresent {
		ie.Qcl_InfoPeriodicCSI_RS = new(TCI_StateId)
		if err = ie.Qcl_InfoPeriodicCSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode Qcl_InfoPeriodicCSI_RS", err)
		}
	}
	return nil
}
