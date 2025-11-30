package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkTransformPrecoding_r16 struct {
	Pi2BPSK_ScramblingID0 *int64 `lb:0,ub:65535,optional`
	Pi2BPSK_ScramblingID1 *int64 `lb:0,ub:65535,optional`
}

func (ie *DMRS_UplinkTransformPrecoding_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pi2BPSK_ScramblingID0 != nil, ie.Pi2BPSK_ScramblingID1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pi2BPSK_ScramblingID0 != nil {
		if err = w.WriteInteger(*ie.Pi2BPSK_ScramblingID0, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode Pi2BPSK_ScramblingID0", err)
		}
	}
	if ie.Pi2BPSK_ScramblingID1 != nil {
		if err = w.WriteInteger(*ie.Pi2BPSK_ScramblingID1, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode Pi2BPSK_ScramblingID1", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkTransformPrecoding_r16) Decode(r *aper.AperReader) error {
	var err error
	var Pi2BPSK_ScramblingID0Present bool
	if Pi2BPSK_ScramblingID0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pi2BPSK_ScramblingID1Present bool
	if Pi2BPSK_ScramblingID1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pi2BPSK_ScramblingID0Present {
		var tmp_int_Pi2BPSK_ScramblingID0 int64
		if tmp_int_Pi2BPSK_ScramblingID0, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode Pi2BPSK_ScramblingID0", err)
		}
		ie.Pi2BPSK_ScramblingID0 = &tmp_int_Pi2BPSK_ScramblingID0
	}
	if Pi2BPSK_ScramblingID1Present {
		var tmp_int_Pi2BPSK_ScramblingID1 int64
		if tmp_int_Pi2BPSK_ScramblingID1, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode Pi2BPSK_ScramblingID1", err)
		}
		ie.Pi2BPSK_ScramblingID1 = &tmp_int_Pi2BPSK_ScramblingID1
	}
	return nil
}
