func containsDuplicate(nums []int) bool {
    // res := 0 // the element might appear more than two times, not works
    // for _, num := range(nums) {
    //     res ^= num
    // }
    // if res == 1 {
    //     return true
    // }
    // return false
    
    hMap := make(map[int]int)
    for _, num := range(nums) {
        _, ok := hMap[num]; if ok {
            // map[num] = val ++
            return true
        } else {
            hMap[num] = 1
        }
    }
    return false
}
