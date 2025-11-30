package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_r16_eutra struct {
	BandParametersSidelinkEUTRA1_r16 *[]byte `optional`
	BandParametersSidelinkEUTRA2_r16 *[]byte `optional`
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_eutra) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.BandParametersSidelinkEUTRA1_r16 != nil, ie.BandParametersSidelinkEUTRA2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BandParametersSidelinkEUTRA1_r16 != nil {
		if err = w.WriteOctetString(*ie.BandParametersSidelinkEUTRA1_r16, nil, false); err != nil {
			return utils.WrapError("Encode BandParametersSidelinkEUTRA1_r16", err)
		}
	}
	if ie.BandParametersSidelinkEUTRA2_r16 != nil {
		if err = w.WriteOctetString(*ie.BandParametersSidelinkEUTRA2_r16, nil, false); err != nil {
			return utils.WrapError("Encode BandParametersSidelinkEUTRA2_r16", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_eutra) Decode(r *aper.AperReader) error {
	var err error
	var BandParametersSidelinkEUTRA1_r16Present bool
	if BandParametersSidelinkEUTRA1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandParametersSidelinkEUTRA2_r16Present bool
	if BandParametersSidelinkEUTRA2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandParametersSidelinkEUTRA1_r16Present {
		var tmp_os_BandParametersSidelinkEUTRA1_r16 []byte
		if tmp_os_BandParametersSidelinkEUTRA1_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode BandParametersSidelinkEUTRA1_r16", err)
		}
		ie.BandParametersSidelinkEUTRA1_r16 = &tmp_os_BandParametersSidelinkEUTRA1_r16
	}
	if BandParametersSidelinkEUTRA2_r16Present {
		var tmp_os_BandParametersSidelinkEUTRA2_r16 []byte
		if tmp_os_BandParametersSidelinkEUTRA2_r16, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode BandParametersSidelinkEUTRA2_r16", err)
		}
		ie.BandParametersSidelinkEUTRA2_r16 = &tmp_os_BandParametersSidelinkEUTRA2_r16
	}
	return nil
}
