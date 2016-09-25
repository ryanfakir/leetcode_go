package src

func reverseVowels(s string) string {
    arr := []rune(s)
    dict := make(map[rune]bool)
    dict['a'] = true
    dict['e'] = true
    dict['i'] = true
    dict['o'] = true
    dict['u'] = true
    dict['A'] = true
    dict['E'] = true
    dict['I'] = true
    dict['O'] = true
    dict['U'] = true
    i, j := 0, len(s)-1
    for i < j {
        if dict[arr[i]] && dict[arr[j]] {
            temp := arr[i]
            arr[i] = arr[j]
            arr[j] = temp
            i++
            j--
        } else if _, ok := dict[arr[i]]; ok {
            j--;
        } else if _, ok := dict[arr[j]]; ok {
            i++;
        } else {
            i++;
            j--;
        }
    }
    return string(arr)
}