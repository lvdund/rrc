package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentCell struct {
	ServCellIndex              ServCellIndex              `madatory`
	UplinkDirectCurrentBWP     []UplinkTxDirectCurrentBWP `lb:1,ub:maxNrofBWPs,madatory`
	UplinkDirectCurrentBWP_SUL []UplinkTxDirectCurrentBWP `lb:1,ub:maxNrofBWPs,optional,ext-1`
}

func (ie *UplinkTxDirectCurrentCell) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := len(ie.UplinkDirectCurrentBWP_SUL) > 0
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ServCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellIndex", err)
	}
	tmp_UplinkDirectCurrentBWP := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, aper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
	for _, i := range ie.UplinkDirectCurrentBWP {
		tmp_UplinkDirectCurrentBWP.Value = append(tmp_UplinkDirectCurrentBWP.Value, &i)
	}
	if err = tmp_UplinkDirectCurrentBWP.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkDirectCurrentBWP", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.UplinkDirectCurrentBWP_SUL) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UplinkTxDirectCurrentCell", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.UplinkDirectCurrentBWP_SUL) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode UplinkDirectCurrentBWP_SUL optional
			if len(ie.UplinkDirectCurrentBWP_SUL) > 0 {
				tmp_UplinkDirectCurrentBWP_SUL := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, aper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
				for _, i := range ie.UplinkDirectCurrentBWP_SUL {
					tmp_UplinkDirectCurrentBWP_SUL.Value = append(tmp_UplinkDirectCurrentBWP_SUL.Value, &i)
				}
				if err = tmp_UplinkDirectCurrentBWP_SUL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkDirectCurrentBWP_SUL", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *UplinkTxDirectCurrentCell) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ServCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode ServCellIndex", err)
	}
	tmp_UplinkDirectCurrentBWP := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, aper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
	fn_UplinkDirectCurrentBWP := func() *UplinkTxDirectCurrentBWP {
		return new(UplinkTxDirectCurrentBWP)
	}
	if err = tmp_UplinkDirectCurrentBWP.Decode(r, fn_UplinkDirectCurrentBWP); err != nil {
		return utils.WrapError("Decode UplinkDirectCurrentBWP", err)
	}
	ie.UplinkDirectCurrentBWP = []UplinkTxDirectCurrentBWP{}
	for _, i := range tmp_UplinkDirectCurrentBWP.Value {
		ie.UplinkDirectCurrentBWP = append(ie.UplinkDirectCurrentBWP, *i)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			UplinkDirectCurrentBWP_SULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode UplinkDirectCurrentBWP_SUL optional
			if UplinkDirectCurrentBWP_SULPresent {
				tmp_UplinkDirectCurrentBWP_SUL := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, aper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
				fn_UplinkDirectCurrentBWP_SUL := func() *UplinkTxDirectCurrentBWP {
					return new(UplinkTxDirectCurrentBWP)
				}
				if err = tmp_UplinkDirectCurrentBWP_SUL.Decode(extReader, fn_UplinkDirectCurrentBWP_SUL); err != nil {
					return utils.WrapError("Decode UplinkDirectCurrentBWP_SUL", err)
				}
				ie.UplinkDirectCurrentBWP_SUL = []UplinkTxDirectCurrentBWP{}
				for _, i := range tmp_UplinkDirectCurrentBWP_SUL.Value {
					ie.UplinkDirectCurrentBWP_SUL = append(ie.UplinkDirectCurrentBWP_SUL, *i)
				}
			}
		}
	}
	return nil
}
