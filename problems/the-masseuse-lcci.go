func massage(nums []int) int {
    
    i := 0
    if len(nums) == 0 {
        return 0
    }
    for i = range nums {
        if i == 0{

        }else if i == 1{
            if nums[0] > nums[1]{
                nums[1] = nums[0] 
            }
        }else{
            if nums[i-1] > nums[i-2]+nums[i]{
                nums[i] = nums[i-1]
            }else{
                nums[i] = nums[i-2]+nums[i]
            }
        }
        /*else if b > a + nums[i] {
            a = b       
        }else{
            //c = a +v
            //a = b
            //b = c
            a,b = b,a +nums[i]
        }*/

    }
    return nums[i]
}
