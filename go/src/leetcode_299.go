func getHint(secret string, guess string) string {
    dict := make(map[byte]int)
    var bull, cow int
    for i:= 0; i < len(guess); i++ {
        if secret[i] == guess[i] {
            bull++
        } else {
            if dict[secret[i]] < 0 {
                cow++
            }
            dict[secret[i]]++
            if dict[guess[i]] > 0 {
                cow++
            }
            dict[guess[i]]--
        }
    }
    return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}
