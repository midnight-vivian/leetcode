package maximum_binary_tree

import (
	"github.com/midnight-vivian/go-data-structures/src/utils"
)

/*
Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
Note:
The size of the given array will be in the range [1,1000].
*/

func constructMaximumBinaryTree(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := 0
	for i := range nums {
		if nums[i] > nums[max] {
			max = i
		}
	}
	n := &utils.TreeNode{
		Val:   nums[max],
		Left:  constructMaximumBinaryTree(nums[:max]),
		Right: constructMaximumBinaryTree(nums[max+1:]),
	}
	return n
}