func toLowerCase(s string) string {
    // diff := 32 // 97 - 65 // A a
    var lower string
    for _, val := range(s) {
        if 90 >= val && val >= 65 {
            lower += string(val + 32)
            continue
        }
        lower += string(val)
    }
    return lower
}
