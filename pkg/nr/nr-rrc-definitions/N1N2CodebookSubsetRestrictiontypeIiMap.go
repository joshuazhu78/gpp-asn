package nrrrcdefinitions

import "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

func (c *N1N2CodebookSubsetRestrictiontypeIi) GetN1N2() (int, int, *asn1.BitString) {
	switch c.N1N2CodebookSubsetRestrictiontypeIi.(type) {
	case *N1N2CodebookSubsetRestrictiontypeIi_TwoOne:
		return 2, 1, c.GetTwoOne()
	case *N1N2CodebookSubsetRestrictiontypeIi_TwoTwo:
		return 2, 2, c.GetTwoTwo()
	case *N1N2CodebookSubsetRestrictiontypeIi_FourOne:
		return 4, 1, c.GetFourOne()
	case *N1N2CodebookSubsetRestrictiontypeIi_ThreeTwo:
		return 3, 2, c.GetThreeTwo()
	case *N1N2CodebookSubsetRestrictiontypeIi_SixOne:
		return 6, 1, c.GetSixOne()
	case *N1N2CodebookSubsetRestrictiontypeIi_FourTwo:
		return 4, 2, c.GetFourTwo()
	case *N1N2CodebookSubsetRestrictiontypeIi_EightOne:
		return 8, 1, c.GetEightOne()
	case *N1N2CodebookSubsetRestrictiontypeIi_FourThree:
		return 4, 3, c.GetFourThree()
	case *N1N2CodebookSubsetRestrictiontypeIi_SixTwo:
		return 6, 2, c.GetSixTwo()
	case *N1N2CodebookSubsetRestrictiontypeIi_TwelveOne:
		return 12, 1, c.GetTwelveOne()
	case *N1N2CodebookSubsetRestrictiontypeIi_FourFour:
		return 4, 4, c.GetFourFour()
	case *N1N2CodebookSubsetRestrictiontypeIi_EightTwo:
		return 8, 2, c.GetEightTwo()
	case *N1N2CodebookSubsetRestrictiontypeIi_SixteenOne:
		return 16, 1, c.GetSixteenOne()
	}
	return 0, 0, nil
}
