package src

func merge(nums1 []int, m int, nums2 []int, n int)  {
    index, i, j := m+n-1, m-1, n-1;
    for i>=0 && j >=0 {
        if nums1[i] > nums2[j] {
            nums1[index] = nums1[i]
            index--
            i--
        } else {
            nums1[index] = nums2[j]
            index--
            j--
        }
    }

    for i>=0 {
        nums1[index] = nums1[i]
        index--
        i--
    }
    for j>=0 {
        nums1[index] = nums2[j]
        index--
        j--
    }
}