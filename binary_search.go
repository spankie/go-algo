package main

// more optimal solution
func binarySearch(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    left := 0
    right := len(nums) - 1
    
    for left <= right {
        mid := (left + right) / 2
        
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        }
    }
    
    return -1
}

func search(nums []int, target int) int {
    seen := -1
    high := len(nums) - 1
    low := 0

    for low <= high {
        median := (high + low)/ 2
        if nums[median] == target {
            return median
        }
        if nums[median] < target {
            low = median+1
        }
        if nums[median] > target {
            high = median - 1
        }
    }
    return seen
    // [-1,0,3,5,9,12]
    
    
    // LINEAR SEARCH
    /*
    numMap := map[int]int{}
    for i, v := range nums {
        numMap[v] = i
    }
    if i, ok := numMap[target]; ok {
        return i
    }
    return -1
    */
}

/*

check the middle item in the slice
- if it is equal return l/2
- if it is greater set nums = nums[l/2:]
- if it is lesser set nums = nums[:l/2]
[-1,0,3,5,9,12]
l = 6
[5,9,12]
l = 3

*/


// question 2; first bad version
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    low, high := 0, n
    for low <= high {
        median := (high+low)/2
        if median == low {
            return high
        }
        if isBadVersion(median) {
            high = median
        } else {
            low = median
        }
    }
    return 1
}
