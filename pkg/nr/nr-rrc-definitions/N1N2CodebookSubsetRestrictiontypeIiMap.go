package nrrrcdefinitions

func (c *N1N2CodebookSubsetRestrictiontypeIi) GetN1N2() (int, int) {
	switch c.N1N2CodebookSubsetRestrictiontypeIi.(type) {
	case *N1N2CodebookSubsetRestrictiontypeIi_TwoOne:
		return 2, 1
	case *N1N2CodebookSubsetRestrictiontypeIi_TwoTwo:
		return 2, 2
	case *N1N2CodebookSubsetRestrictiontypeIi_FourOne:
		return 4, 1
	case *N1N2CodebookSubsetRestrictiontypeIi_ThreeTwo:
		return 3, 2
	case *N1N2CodebookSubsetRestrictiontypeIi_SixOne:
		return 6, 1
	case *N1N2CodebookSubsetRestrictiontypeIi_FourTwo:
		return 4, 2
	case *N1N2CodebookSubsetRestrictiontypeIi_EightOne:
		return 8, 1
	case *N1N2CodebookSubsetRestrictiontypeIi_FourThree:
		return 4, 3
	case *N1N2CodebookSubsetRestrictiontypeIi_SixTwo:
		return 6, 2
	case *N1N2CodebookSubsetRestrictiontypeIi_TwelveOne:
		return 12, 1
	case *N1N2CodebookSubsetRestrictiontypeIi_FourFour:
		return 4, 4
	case *N1N2CodebookSubsetRestrictiontypeIi_EightTwo:
		return 8, 2
	case *N1N2CodebookSubsetRestrictiontypeIi_SixteenOne:
		return 16, 1
	}
	return 0, 0
}
