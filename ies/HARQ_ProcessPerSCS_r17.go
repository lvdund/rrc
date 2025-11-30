package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type HARQ_ProcessPerSCS_r17 struct {
	Scs_120kHz_r17 *HARQ_ProcessPerSCS_r17_scs_120kHz_r17 `optional`
	Scs_480kHz_r17 *HARQ_ProcessPerSCS_r17_scs_480kHz_r17 `optional`
	Scs_960kHz_r17 *HARQ_ProcessPerSCS_r17_scs_960kHz_r17 `optional`
}

func (ie *HARQ_ProcessPerSCS_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_120kHz_r17 != nil, ie.Scs_480kHz_r17 != nil, ie.Scs_960kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_120kHz_r17 != nil {
		if err = ie.Scs_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_120kHz_r17", err)
		}
	}
	if ie.Scs_480kHz_r17 != nil {
		if err = ie.Scs_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_480kHz_r17", err)
		}
	}
	if ie.Scs_960kHz_r17 != nil {
		if err = ie.Scs_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_960kHz_r17", err)
		}
	}
	return nil
}

func (ie *HARQ_ProcessPerSCS_r17) Decode(r *aper.AperReader) error {
	var err error
	var Scs_120kHz_r17Present bool
	if Scs_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_480kHz_r17Present bool
	if Scs_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_960kHz_r17Present bool
	if Scs_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_120kHz_r17Present {
		ie.Scs_120kHz_r17 = new(HARQ_ProcessPerSCS_r17_scs_120kHz_r17)
		if err = ie.Scs_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_120kHz_r17", err)
		}
	}
	if Scs_480kHz_r17Present {
		ie.Scs_480kHz_r17 = new(HARQ_ProcessPerSCS_r17_scs_480kHz_r17)
		if err = ie.Scs_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_480kHz_r17", err)
		}
	}
	if Scs_960kHz_r17Present {
		ie.Scs_960kHz_r17 = new(HARQ_ProcessPerSCS_r17_scs_960kHz_r17)
		if err = ie.Scs_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_960kHz_r17", err)
		}
	}
	return nil
}
