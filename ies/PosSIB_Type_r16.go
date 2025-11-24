package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSIB_Type_r16 struct {
	Encrypted_r16  *PosSIB_Type_r16_encrypted_r16 `optional`
	Gnss_id_r16    *GNSS_ID_r16                   `optional`
	Sbas_id_r16    *SBAS_ID_r16                   `optional`
	PosSibType_r16 PosSIB_Type_r16_posSibType_r16 `madatory`
	AreaScope_r16  *PosSIB_Type_r16_areaScope_r16 `optional`
}

func (ie *PosSIB_Type_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Encrypted_r16 != nil, ie.Gnss_id_r16 != nil, ie.Sbas_id_r16 != nil, ie.AreaScope_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Encrypted_r16 != nil {
		if err = ie.Encrypted_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Encrypted_r16", err)
		}
	}
	if ie.Gnss_id_r16 != nil {
		if err = ie.Gnss_id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Gnss_id_r16", err)
		}
	}
	if ie.Sbas_id_r16 != nil {
		if err = ie.Sbas_id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sbas_id_r16", err)
		}
	}
	if err = ie.PosSibType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PosSibType_r16", err)
	}
	if ie.AreaScope_r16 != nil {
		if err = ie.AreaScope_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AreaScope_r16", err)
		}
	}
	return nil
}

func (ie *PosSIB_Type_r16) Decode(r *uper.UperReader) error {
	var err error
	var Encrypted_r16Present bool
	if Encrypted_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Gnss_id_r16Present bool
	if Gnss_id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sbas_id_r16Present bool
	if Sbas_id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AreaScope_r16Present bool
	if AreaScope_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Encrypted_r16Present {
		ie.Encrypted_r16 = new(PosSIB_Type_r16_encrypted_r16)
		if err = ie.Encrypted_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Encrypted_r16", err)
		}
	}
	if Gnss_id_r16Present {
		ie.Gnss_id_r16 = new(GNSS_ID_r16)
		if err = ie.Gnss_id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Gnss_id_r16", err)
		}
	}
	if Sbas_id_r16Present {
		ie.Sbas_id_r16 = new(SBAS_ID_r16)
		if err = ie.Sbas_id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sbas_id_r16", err)
		}
	}
	if err = ie.PosSibType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PosSibType_r16", err)
	}
	if AreaScope_r16Present {
		ie.AreaScope_r16 = new(PosSIB_Type_r16_areaScope_r16)
		if err = ie.AreaScope_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AreaScope_r16", err)
		}
	}
	return nil
}
