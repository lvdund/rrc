package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2NR_r16 struct {
	SsbFrequency_r16   *ARFCN_ValueNR   `optional`
	RefFreqCSI_RS_r16  *ARFCN_ValueNR   `optional`
	MeasResultList_r16 MeasResultListNR `madatory`
}

func (ie *MeasResult2NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SsbFrequency_r16 != nil, ie.RefFreqCSI_RS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SsbFrequency_r16 != nil {
		if err = ie.SsbFrequency_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SsbFrequency_r16", err)
		}
	}
	if ie.RefFreqCSI_RS_r16 != nil {
		if err = ie.RefFreqCSI_RS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RefFreqCSI_RS_r16", err)
		}
	}
	if err = ie.MeasResultList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultList_r16", err)
	}
	return nil
}

func (ie *MeasResult2NR_r16) Decode(r *aper.AperReader) error {
	var err error
	var SsbFrequency_r16Present bool
	if SsbFrequency_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RefFreqCSI_RS_r16Present bool
	if RefFreqCSI_RS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SsbFrequency_r16Present {
		ie.SsbFrequency_r16 = new(ARFCN_ValueNR)
		if err = ie.SsbFrequency_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SsbFrequency_r16", err)
		}
	}
	if RefFreqCSI_RS_r16Present {
		ie.RefFreqCSI_RS_r16 = new(ARFCN_ValueNR)
		if err = ie.RefFreqCSI_RS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RefFreqCSI_RS_r16", err)
		}
	}
	if err = ie.MeasResultList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultList_r16", err)
	}
	return nil
}
