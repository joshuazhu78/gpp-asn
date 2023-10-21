package nrrrcdefinitions

func (c *CsiResourcePeriodicityAndOffset) GetPeriodicityAndOffset() (int, int) {
	switch c.CsiResourcePeriodicityAndOffset.(type) {
	case *CsiResourcePeriodicityAndOffset_Slots4:
		return 4, int(c.GetSlots4())
	case *CsiResourcePeriodicityAndOffset_Slots5:
		return 5, int(c.GetSlots5())
	case *CsiResourcePeriodicityAndOffset_Slots8:
		return 8, int(c.GetSlots8())
	case *CsiResourcePeriodicityAndOffset_Slots10:
		return 10, int(c.GetSlots10())
	case *CsiResourcePeriodicityAndOffset_Slots16:
		return 16, int(c.GetSlots16())
	case *CsiResourcePeriodicityAndOffset_Slots20:
		return 20, int(c.GetSlots20())
	case *CsiResourcePeriodicityAndOffset_Slots32:
		return 32, int(c.GetSlots32())
	case *CsiResourcePeriodicityAndOffset_Slots40:
		return 40, int(c.GetSlots40())
	case *CsiResourcePeriodicityAndOffset_Slots64:
		return 64, int(c.GetSlots64())
	case *CsiResourcePeriodicityAndOffset_Slots80:
		return 80, int(c.GetSlots80())
	case *CsiResourcePeriodicityAndOffset_Slots160:
		return 160, int(c.GetSlots160())
	case *CsiResourcePeriodicityAndOffset_Slots320:
		return 320, int(c.GetSlots320())
	case *CsiResourcePeriodicityAndOffset_Slots640:
		return 640, int(c.GetSlots640())
	}
	return 0, 0
}

func (c *CsiResourcePeriodicityAndOffset) SetOffset(o int) {
	switch c.CsiResourcePeriodicityAndOffset.(type) {
	case *CsiResourcePeriodicityAndOffset_Slots4:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots4).Slots4 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots5:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots5).Slots5 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots8:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots8).Slots8 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots10:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots10).Slots10 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots16:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots16).Slots16 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots20:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots20).Slots20 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots32:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots32).Slots32 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots40:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots40).Slots40 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots64:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots64).Slots64 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots80:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots80).Slots80 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots160:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots160).Slots160 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots320:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots320).Slots320 = int32(o)
	case *CsiResourcePeriodicityAndOffset_Slots640:
		c.CsiResourcePeriodicityAndOffset.(*CsiResourcePeriodicityAndOffset_Slots640).Slots640 = int32(o)
	}
}
