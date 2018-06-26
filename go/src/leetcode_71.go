func simplifyPath(path string) string {
    tmp := strings.Split(path, "/")
    var arr []string
    for _, v := range tmp {
        if v == "" || v == "." {continue}
        if v == ".."  {
            if len(arr) == 0 {continue} 
            arr = arr[:len(arr)-1]
        } else {
            arr = append(arr, v)
        }
    }
    return "/" + strings.Join(arr, "/")
}
