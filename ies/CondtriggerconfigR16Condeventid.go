package ies

// CondTriggerConfig-r16-condEventId ::= CHOICE
// Extensible
const (
	CondtriggerconfigR16CondeventidChoiceNothing = iota
	CondtriggerconfigR16CondeventidChoiceCondeventa3
	CondtriggerconfigR16CondeventidChoiceCondeventa5
	CondtriggerconfigR16CondeventidChoiceCondeventa4R17
	CondtriggerconfigR16CondeventidChoiceCondeventd1R17
	CondtriggerconfigR16CondeventidChoiceCondeventt1R17
)

type CondtriggerconfigR16Condeventid struct {
	Choice         uint64
	Condeventa3    *CondtriggerconfigR16CondeventidCondeventa3
	Condeventa5    *CondtriggerconfigR16CondeventidCondeventa5
	Condeventa4R17 *CondtriggerconfigR16CondeventidCondeventa4R17
	Condeventd1R17 *CondtriggerconfigR16CondeventidCondeventd1R17
	Condeventt1R17 *CondtriggerconfigR16CondeventidCondeventt1R17
}
