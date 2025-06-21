# [510. Inorder Successor in BST II]()
- https://hackmd.io/@Zihao0917/S193Uml9F
- https://zhuanlan.zhihu.com/p/462714826
- https://github.com/LLancelot/LeetCode/blob/master/%E6%AF%8F%E6%97%A5%E4%B8%80%E9%A2%98/510.%20In-order%20Successor%20in%20BST%20II.md


## Level
Medium

## Description
Given a binary search tree and a node in it, find the in-order successor of that node in the BST.
The successor of a node p is the node with the smallest key greater than p.val.
You will have direct access to the node but not to the root of the tree. Each node will have a reference to its parent node.


#### Example 1:
Input: tree = [2,1,3], node = 1
Output: 2
Explanation: 1's in-order successor node is 2. Note that both the node and the return value is of Node type.

#### Example 2:
Input: tree = [5,3,6,2,4,null,null,1], node = 6
Output: null
Explanation: There is no in-order successor of the current node, so the answer is null.

#### Example 3:
Input: tree = [15,6,18,3,7,17,20,2,4,null,13,null,null,null,null,null,null,null,null,9], node = 15
Output: 17


## Note:
- If the given node has no in-order successor in the tree, return null.
- It's guaranteed that the values of the tree are unique.
- Remember that we are using the Node type instead of TreeNode type so their string representation are different.
 

## Follow up:
Could you solve it without looking up any of the node's values?

## Similar
- 285. Inorder Successor In BST