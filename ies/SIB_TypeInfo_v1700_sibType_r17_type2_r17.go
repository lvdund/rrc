package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB_TypeInfo_v1700_sibType_r17_type2_r17 struct {
	PosSibType_r17 SIB_TypeInfo_v1700_sibType_r17_type2_r17_posSibType_r17 `madatory`
	Encrypted_r17  *SIB_TypeInfo_v1700_sibType_r17_type2_r17_encrypted_r17 `optional`
	Gnss_id_r17    *GNSS_ID_r16                                            `optional`
	Sbas_id_r17    *SBAS_ID_r16                                            `optional`
}

func (ie *SIB_TypeInfo_v1700_sibType_r17_type2_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Encrypted_r17 != nil, ie.Gnss_id_r17 != nil, ie.Sbas_id_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PosSibType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PosSibType_r17", err)
	}
	if ie.Encrypted_r17 != nil {
		if err = ie.Encrypted_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Encrypted_r17", err)
		}
	}
	if ie.Gnss_id_r17 != nil {
		if err = ie.Gnss_id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Gnss_id_r17", err)
		}
	}
	if ie.Sbas_id_r17 != nil {
		if err = ie.Sbas_id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sbas_id_r17", err)
		}
	}
	return nil
}

func (ie *SIB_TypeInfo_v1700_sibType_r17_type2_r17) Decode(r *aper.AperReader) error {
	var err error
	var Encrypted_r17Present bool
	if Encrypted_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Gnss_id_r17Present bool
	if Gnss_id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sbas_id_r17Present bool
	if Sbas_id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PosSibType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PosSibType_r17", err)
	}
	if Encrypted_r17Present {
		ie.Encrypted_r17 = new(SIB_TypeInfo_v1700_sibType_r17_type2_r17_encrypted_r17)
		if err = ie.Encrypted_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Encrypted_r17", err)
		}
	}
	if Gnss_id_r17Present {
		ie.Gnss_id_r17 = new(GNSS_ID_r16)
		if err = ie.Gnss_id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Gnss_id_r17", err)
		}
	}
	if Sbas_id_r17Present {
		ie.Sbas_id_r17 = new(SBAS_ID_r16)
		if err = ie.Sbas_id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sbas_id_r17", err)
		}
	}
	return nil
}
