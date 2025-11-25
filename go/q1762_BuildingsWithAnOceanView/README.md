# [1762 - Buildings With an Ocean View]()
- https://leetcode.ca/2020-09-26-1762-Buildings-With-an-Ocean-View/
- https://hackmd.io/@Zihao0917/SJMvZ6EHY
- https://allenhsieh1992.com/leetcode/medium/1762_buildings_with_an_ocean_view/

## Level
Medium

## Description
There are n buildings in a line. You are given an integer array heights of size n that represents the heights of the buildings in the line.

The ocean is to the right of the buildings. A building has an ocean view if the building can see the ocean without obstructions. Formally, a building has an ocean view if all the buildings to its right have a smaller height.

Return a list of indices (0-indexed) of buildings that have an ocean view, sorted in increasing order.


#### Example 1:
Input: heights = [4,2,3,1]
Output: [0,2,3]
Explanation: Building 1 (0-indexed) does not have an ocean view because building 2 is taller.
 

#### Example 2:
Input: heights = [4,2,3,1]
Output: [0,2,3]
Explanation: Building 1 (0-indexed) does not have an ocean view because building 2 is taller.


#### Example 3:
Input: heights = [1,3,2,4]
Output: [3]
Explanation: Only building 3 has an ocean view.


## Constraints
- 1 <= heights.length <= 105
- 1 <= heights[i] <= 109