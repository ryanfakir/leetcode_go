func ipToCIDR(ip string, n int) []string {
    arr := strings.Split(ip, ".")
    var num int64
    for _, a := range arr {
        x, _ := strconv.Atoi(a)
        num += num * 256 + int64(x)
    }
    var res []string
    for n > 0 {
        step := num & (-num)
        for step > int64(n) {step /= 2}
        res = append(res, convert(num, step))
        num += step
        n -= int(step)
    }
    return res
}

func convert(num, step int64) string {
    return strconv.Itoa(int(num >> 24) & 255) + "." + strconv.Itoa(int(num >> 16) & 255) +
    "." + strconv.Itoa(int(num >> 8) & 255) + "." + strconv.Itoa(int(num) & 255) + "/" + strconv.Itoa(32 - int(math.Log2(float64(step))))
}
