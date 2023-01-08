/*
 * @lc app=leetcode id=4 lang=c
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
int max(int x, int y) {
    return x > y ? x : y;
}

int min(int x, int y) {
    return x < y ? x : y;
}



double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    
    if (nums1Size > nums2Size) {
        int temp_size;
        int* temp_ptr;
        
        temp_size = nums1Size;
        nums1Size = nums2Size;
        nums2Size = temp_size;
        
        temp_ptr = nums1;
        nums1 = nums2;
        nums2 = temp_ptr;
        
    }
    

    int bound_low = 0;
    int bound_high = nums1Size;
    

    
    while (bound_low <= bound_high) {
        
        int part_1;
        int part_2;
    
        int part_1_left_max;
        int part_1_right_min;
        int part_2_left_max;
        int part_2_right_min;
        
        part_1 = (bound_low + bound_high)/2;
        part_2 = (nums1Size + nums2Size + 1)/2 - part_1;
        
        
        part_1_left_max = part_1 > 0 ? nums1[part_1 - 1] : INT_MIN;
        part_1_right_min = part_1 == nums1Size ? INT_MAX : nums1[part_1];
        
        
        part_2_left_max = part_2 > 0 ? nums2[part_2 - 1] : INT_MIN;
        part_2_right_min = part_2 == nums2Size ? INT_MAX : nums2[part_2];
        
        
        if (part_1_left_max <= part_2_right_min && part_2_left_max <= part_1_right_min) {
            if ((nums1Size + nums2Size)%2 == 0) {
                return (double)(max(part_1_left_max, part_2_left_max) + min(part_1_right_min, part_2_right_min))/2;
            } 
            
            else 
                return (double)(max(part_1_left_max, part_2_left_max));
        
        }
        if (part_1_left_max > part_2_right_min)
            bound_high = part_1 - 1;
        
        else
            bound_low = part_1 + 1;
           
    }
    
    return 0;
    

}
// @lc code=end

