package transfer

import "time"

func GetFakeTransfers() []Transfer {
	return []Transfer{
		{
			OriginID:      1,
			DestinationID: 2,
			Amount:        500,
			EffectiveDate: time.Date(2020, 10, 20, 0,0,0, 0, time.UTC),
		},
		{
			OriginID:      1,
			DestinationID: 2,
			Amount:        500,
			EffectiveDate: time.Date(2020, 9, 20, 0,0,0, 0, time.UTC),
		},
		{
			OriginID:      2,
			DestinationID: 1,
			Amount:        500,
			EffectiveDate: time.Date(2020, 8, 20, 0,0,0, 0, time.UTC),
		},
		{
			OriginID:      2,
			DestinationID: 1,
			Amount:        500,
			EffectiveDate: time.Date(2020, 7, 20, 0,0,0, 0, time.UTC),
		},
		{
			OriginID:      1,
			DestinationID: 2,
			Amount:        500,
			EffectiveDate: time.Date(2020, 6, 20, 0,0,0, 0, time.UTC),
		},
	}
}

func GetFakeTransfer() Transfer{
	return Transfer{
		OriginID:      1,
		DestinationID: 2,
		Amount:        500,
		EffectiveDate: time.Date(2020, 5, 20, 0,0,0, 0, time.UTC),
	}
}