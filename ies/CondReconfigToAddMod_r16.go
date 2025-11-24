package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondReconfigToAddMod_r16 struct {
	CondReconfigId_r16       CondReconfigId_r16 `madatory`
	CondExecutionCond_r16    []MeasId           `lb:1,ub:2,optional`
	CondRRCReconfig_r16      *[]byte            `optional`
	CondExecutionCondSCG_r17 *[]byte            `optional,ext-1`
}

func (ie *CondReconfigToAddMod_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.CondExecutionCondSCG_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.CondExecutionCond_r16) > 0, ie.CondRRCReconfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CondReconfigId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CondReconfigId_r16", err)
	}
	if len(ie.CondExecutionCond_r16) > 0 {
		tmp_CondExecutionCond_r16 := utils.NewSequence[*MeasId]([]*MeasId{}, uper.Constraint{Lb: 1, Ub: 2}, false)
		for _, i := range ie.CondExecutionCond_r16 {
			tmp_CondExecutionCond_r16.Value = append(tmp_CondExecutionCond_r16.Value, &i)
		}
		if err = tmp_CondExecutionCond_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondExecutionCond_r16", err)
		}
	}
	if ie.CondRRCReconfig_r16 != nil {
		if err = w.WriteOctetString(*ie.CondRRCReconfig_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode CondRRCReconfig_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.CondExecutionCondSCG_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CondReconfigToAddMod_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.CondExecutionCondSCG_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CondExecutionCondSCG_r17 optional
			if ie.CondExecutionCondSCG_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.CondExecutionCondSCG_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode CondExecutionCondSCG_r17", err)
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

func (ie *CondReconfigToAddMod_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CondExecutionCond_r16Present bool
	if CondExecutionCond_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondRRCReconfig_r16Present bool
	if CondRRCReconfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CondReconfigId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CondReconfigId_r16", err)
	}
	if CondExecutionCond_r16Present {
		tmp_CondExecutionCond_r16 := utils.NewSequence[*MeasId]([]*MeasId{}, uper.Constraint{Lb: 1, Ub: 2}, false)
		fn_CondExecutionCond_r16 := func() *MeasId {
			return new(MeasId)
		}
		if err = tmp_CondExecutionCond_r16.Decode(r, fn_CondExecutionCond_r16); err != nil {
			return utils.WrapError("Decode CondExecutionCond_r16", err)
		}
		ie.CondExecutionCond_r16 = []MeasId{}
		for _, i := range tmp_CondExecutionCond_r16.Value {
			ie.CondExecutionCond_r16 = append(ie.CondExecutionCond_r16, *i)
		}
	}
	if CondRRCReconfig_r16Present {
		var tmp_os_CondRRCReconfig_r16 []byte
		if tmp_os_CondRRCReconfig_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode CondRRCReconfig_r16", err)
		}
		ie.CondRRCReconfig_r16 = &tmp_os_CondRRCReconfig_r16
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			CondExecutionCondSCG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CondExecutionCondSCG_r17 optional
			if CondExecutionCondSCG_r17Present {
				var tmp_os_CondExecutionCondSCG_r17 []byte
				if tmp_os_CondExecutionCondSCG_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode CondExecutionCondSCG_r17", err)
				}
				ie.CondExecutionCondSCG_r17 = &tmp_os_CondExecutionCondSCG_r17
			}
		}
	}
	return nil
}
