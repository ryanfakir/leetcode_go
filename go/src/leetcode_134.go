func canCompleteCircuit(gas []int, cost []int) int {
    var sum, total, start int
    for i := 0; i < len(gas); i++ {
        sum += gas[i] - cost[i]
        total += gas[i] - cost[i]
        if sum < 0 {
            sum = 0
            start = i +1
        }
    }
    if total >= 0 {
        return start
    } else {
        return -1
    }
}
