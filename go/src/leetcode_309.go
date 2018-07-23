func maxProfit(prices []int) int {
    presell, prebuy, sell, buy := 0, 0, 0, -1 << 7
    for _ ,v := range prices {
        prebuy = buy
        buy = int(math.Max(float64(presell - v), float64(prebuy)))
        presell = sell
        sell = int(math.Max(float64(prebuy + v), float64(presell)))
    }
    return sell
}
