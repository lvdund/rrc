package common

import (
	"errors"

	"github.com/lvdund/asn1go/per"
)

type NULL struct{}

type GenericSetuprelease interface {
	Encode(e *per.Encoder) error
	Decode(d *per.Decoder) error
}

type Release_Setup_Choice[T GenericSetuprelease] struct {
	Choice  int
	Release *NULL
	Setup   T
}

func (ie *Release_Setup_Choice[T]) Encode(e *per.Encoder) error {
	choice := e.NewChoiceEncoder(per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			{Name: "release"},
			{Name: "setup"},
		},
		ExtAlternatives: nil,
	})
	switch ie.Choice {
	case 0:
		if err := choice.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case 1:
		if err := choice.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err := ie.Setup.Encode(e); err != nil {
			return err
		}
	default:
		return errors.New("invalid choice value")
	}
	return nil
}

func (ie *Release_Setup_Choice[T]) Decode(d *per.Decoder) error {
	choice := d.NewChoiceDecoder(per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			{Name: "c1"},
			{Name: "messageClassExtension"},
		},
		ExtAlternatives: nil,
	})
	idx, _, _, err := choice.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(idx)
	switch idx {
	case 0:
		if err := d.DecodeNull(); err != nil {
			return err
		}
		ie.Release = &NULL{}
	case 1:
		var tmp T
		if err = tmp.Decode(d); err != nil {
			return err
		}
		ie.Setup = tmp
	default:
		return errors.New("invalid choice value")
	}
	return nil
}
