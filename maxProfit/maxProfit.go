package maxProfit

func maxProfit(prices []int) int {

	var profit, buy, sales int

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			buy = prices[i]
		} else if prices[i]-buy > sales {
			sales = prices[i] - buy
		} else if prices[i] < buy && sales == 0 {
			buy = prices[i]
		} else {
			profit += sales
			buy = prices[i]
			sales = 0
		}
	}
	return profit + sales
}
