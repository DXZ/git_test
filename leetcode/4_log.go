package main

import (
	"fmt"
)

// // find (m+n)/2 min number
// func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
// 	k := len(nums1) + len(nums2)
//     if (k%2==0) {
//         return float64(findkth(nums1,nums2,k/2) + findkth(nums1,nums2,k/2+1))/2.0
//     }else {
//         return float64(findkth(nums1,nums2,k/2+1))
//     }
// }

// func findkth3(nums1 []int,nums2 []int,k int) (int){
//     fmt.Println(nums1,len(nums1),nums2,k,"aaaaaaaaaaaaaa")
//     if len(nums1) > len(nums2) {
//         return findkth(nums2,nums1,k)
//     }
    
//     if len(nums1) == 0 {
//         fmt.Println(k,"aacc",nums2[k-1])
//         return nums2[k-1]
//     }

//     if k <= 1 {
//         if len(nums1) >= 1 && nums1[0] < nums2[0] {
//             return nums1[0]
//         }else {
//             return nums2[0]
//         }
//     }

//     var pa int
//     if k / 2 < len(nums1) {
//         pa = k/2
//     }else {
//         pa = len(nums1)
//     }

//     pb := k - pa
//     // fmt.Println(pa,pb,k,nums1,nums2)
//     if nums1[pa-1] < nums2[pb-1] {
//         return findkth(nums1[pa:],nums2,k-pa)
//     }else if nums1[pa-1] > nums2[pb-1] {
//         return findkth(nums1,nums2[pb:],k-pb)
//     }else{
//         return nums1[pa-1]
//     }

// }




// find (m+n)/2 min number
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     k := len(nums1) + len(nums2)
//     var flag bool
//     if k%2 == 1{
//         flag = false
//         k = k/2 +1
//     }else{
//         flag = true
//         k = k/2
//     }
//     return findkth(nums1,nums2,k,flag)
// }

// func findkth(nums1 []int,nums2 []int,k int,flag bool) (float64){
//     fmt.Println(nums1,len(nums1),nums2,k,"aaaaaaaaaaaaaa")
//     if len(nums1) > len(nums2) {
//         return findkth(nums2,nums1,k,flag)
//     }
    
//     if len(nums1) == 0 {
//         if flag{
//             return float64(nums2[k-1] + nums2[k])/2
//         }else{
//             return float64(nums2[k-1])
//         }
//     }

//     if k <= 1 {
//         if len(nums1) >= 1 && nums1[0] <= nums2[0] {
//             // return nums1[0]
//             if flag{
//                 return float64(nums1[k-1] + nums1[k-1])/2
//             }else{
//                 return float64(nums1[k-1])
//             }
//         }else {
//             // return nums2[0]
//             if flag{
//                 return float64(nums2[k-1] + nums2[k])/2
//             }else{
//                 return float64(nums2[k-1])
//             }
//         }
//     }

//     var pa int
//     if k / 2 < len(nums1) {
//         pa = k/2
//     }else {
//         pa = len(nums1)
//     }

//     pb := k - pa
//     // fmt.Println(pa,pb,k,nums1,nums2)
//     if nums1[pa-1] < nums2[pb-1] {
//         return findkth(nums1[pa:],nums2,k-pa,flag)
//     }else if nums1[pa-1] > nums2[pb-1] {
//         return findkth(nums1,nums2[pb:],k-pb,flag)
//     }else{
//         return float64(nums1[pa-1])
//     }

// }
func main() {
	nums1 := []int{4}
	nums2 := []int{2}
    // fmt.Println(nums1+nums2)
	fmt.Println("aaaa", findMedianSortedArrays(nums1, nums2))
}
