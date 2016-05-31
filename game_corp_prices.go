package acquire

import "github.com/svera/acquire/interfaces"

//Fills the prices chart array with the amounts corresponding to the corporation
//category
func (g *Game) setPricesChart(corpNumber int) map[int]interfaces.Prices {
	initialValues := new([3]interfaces.Prices)
	initialValues[0] = interfaces.Prices{Price: 200, MajorityBonus: 2000, MinorityBonus: 1000}
	initialValues[1] = interfaces.Prices{Price: 300, MajorityBonus: 3000, MinorityBonus: 1500}
	initialValues[2] = interfaces.Prices{Price: 400, MajorityBonus: 4000, MinorityBonus: 2000}
	pricesChart := make(map[int]interfaces.Prices)
	category := g.corpNumberToCategory(corpNumber)

	pricesChart[2] = interfaces.Prices{Price: initialValues[category].Price, MajorityBonus: initialValues[category].MajorityBonus, MinorityBonus: initialValues[category].MinorityBonus}
	pricesChart[3] = interfaces.Prices{Price: initialValues[category].Price + 100, MajorityBonus: initialValues[category].MajorityBonus + 1000, MinorityBonus: initialValues[category].MinorityBonus + 500}
	pricesChart[4] = interfaces.Prices{Price: initialValues[category].Price + 200, MajorityBonus: initialValues[category].MajorityBonus + 2000, MinorityBonus: initialValues[category].MinorityBonus + 1000}
	pricesChart[5] = interfaces.Prices{Price: initialValues[category].Price + 300, MajorityBonus: initialValues[category].MajorityBonus + 3000, MinorityBonus: initialValues[category].MinorityBonus + 1500}
	var i int
	for i = 6; i < 11; i++ {
		pricesChart[i] = interfaces.Prices{Price: initialValues[category].Price + 400, MajorityBonus: initialValues[category].MajorityBonus + 4000, MinorityBonus: initialValues[category].MinorityBonus + 2000}
	}
	for i = 11; i < 21; i++ {
		pricesChart[i] = interfaces.Prices{Price: initialValues[category].Price + 500, MajorityBonus: initialValues[category].MajorityBonus + 5000, MinorityBonus: initialValues[category].MinorityBonus + 2500}
	}
	for i = 21; i < 31; i++ {
		pricesChart[i] = interfaces.Prices{Price: initialValues[category].Price + 600, MajorityBonus: initialValues[category].MajorityBonus + 6000, MinorityBonus: initialValues[category].MinorityBonus + 3000}
	}
	for i = 31; i < 41; i++ {
		pricesChart[i] = interfaces.Prices{Price: initialValues[category].Price + 700, MajorityBonus: initialValues[category].MajorityBonus + 7000, MinorityBonus: initialValues[category].MinorityBonus + 3500}
	}
	pricesChart[41] = interfaces.Prices{Price: initialValues[category].Price + 800, MajorityBonus: initialValues[category].MajorityBonus + 8000, MinorityBonus: initialValues[category].MinorityBonus + 4000}

	return pricesChart
}

// Corporation position in the corporations array do matter.
// First two corporations belong to the first tier, next three to the second one
// and the last two to the third and most expensive tier.
func (g *Game) corpNumberToCategory(corpNumber int) int {
	switch {
	case corpNumber <= 1:
		return 0
	case corpNumber <= 4:
		return 1
	}
	return 2
}
