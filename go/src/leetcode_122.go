package src

func maxProfit(prices []int) int {
    res := 0
    for i, _ := range prices {
        if i == len(prices)-1 {
            break;
        }
        if prices[i+1] - prices[i] >= 0 {
            res += prices[i+1] - prices[i]
        }
    }
    return res
}
