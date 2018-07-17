func summaryRanges(nums []int) []string {
    var res []string
    for i, v := range nums {
        if i == 0 {
            res = append(res, strconv.Itoa(v))
            continue
        }
        if v - nums[i-1] == 1 {
            tmp := res[len(res)-1]
            if strings.Contains(tmp, "->") {
                arr := strings.Split(tmp, "->")
                arr[len(arr)-1] = strconv.Itoa(v)
                tmp = strings.Join(arr, "->")
            } else {
                tmp += "->" + strconv.Itoa(v)
            }
            res[len(res)-1] = tmp
        } else {
            res = append(res, strconv.Itoa(v))
        }
    }
    return res
}
