package nrrrcdefinitions

func (c *N1N2CodebookSubsetRestrictionr16TypeIir16) GetN1N2() (int, int) {
	switch c.N1N2CodebookSubsetRestrictionR16TypeIiR16.(type) {
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_TwoOne:
		return 2, 1
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_TwoTwo:
		return 2, 2
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourOne:
		return 4, 1
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_ThreeTwo:
		return 3, 2
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_SixOne:
		return 6, 1
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourTwo:
		return 4, 2
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_EightOne:
		return 8, 1
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourThree:
		return 4, 3
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_SixTwo:
		return 6, 2
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_TwelveOne:
		return 12, 1
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourFour:
		return 4, 4
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_EightTwo:
		return 8, 2
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_SixteenOne:
		return 16, 1
	}
	return 0, 0
}
