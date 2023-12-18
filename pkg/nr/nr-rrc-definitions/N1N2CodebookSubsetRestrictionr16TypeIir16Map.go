package nrrrcdefinitions

import "github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

func (c *N1N2CodebookSubsetRestrictionr16TypeIir16) GetN1N2() (int, int, *asn1.BitString) {
	switch c.N1N2CodebookSubsetRestrictionR16TypeIiR16.(type) {
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_TwoOne:
		return 2, 1, c.GetTwoOne()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_TwoTwo:
		return 2, 2, c.GetTwoTwo()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourOne:
		return 4, 1, c.GetFourOne()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_ThreeTwo:
		return 3, 2, c.GetThreeTwo()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_SixOne:
		return 6, 1, c.GetSixOne()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourTwo:
		return 4, 2, c.GetFourTwo()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_EightOne:
		return 8, 1, c.GetEightOne()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourThree:
		return 4, 3, c.GetFourThree()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_SixTwo:
		return 6, 2, c.GetSixTwo()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_TwelveOne:
		return 12, 1, c.GetTwelveOne()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_FourFour:
		return 4, 4, c.GetFourFour()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_EightTwo:
		return 8, 2, c.GetEightTwo()
	case *N1N2CodebookSubsetRestrictionr16TypeIir16_SixteenOne:
		return 16, 1, c.GetSixteenOne()
	}
	return 0, 0, nil
}
