func pourWater(heights []int, V int, K int) []int {
    for i:= 0; i < V; i++ {
        
        l, r, lower, inc:= K-1, K+1, heights[K], K
        for l >=0 && heights[l] <= heights[l+1] {
            lower = heights[l]
            l--
        } 
        if lower == heights[K] {
            //right
            for r < len(heights) && heights[r] <= heights[r-1] {
                lower = heights[r]
                r++
            }
            if lower == heights[K] {
                inc = K
            } else {
                for inc < len(heights) && heights[inc] != lower {
                    inc++
                }
            }
        } else {
            for inc >= 0 && heights[inc] != lower {
                    inc--
                }
            
        }
        heights[inc]++
    }
    return heights
}
