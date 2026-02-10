# [660. Remove 9](https://leetcode.ca/2017-09-20-660-Remove-9/)
- https://www.cnblogs.com/grandyang/p/8261714.html

## Level
Hard

## Description
Start from integer 1, remove any integer that contains 9 such as 9, 19, 29…

So now, you will have a new integer sequence: 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, …

Given a positive integer n, you need to return the n-th integer after removing. Note that 1 will be the first integer.

## Example 1:
Input: 9
Output: 10

## Hint: 
- n will not exceed 9 x 10^8.

## 說明
0，1，2，3，4，5，6，7，8 （移除了9）
10，11，12，13，14，15，16，17，18 （移除了19）
.....
80，81，82，83，84，85，86，87，88 （移除了89）
沒有9x（移除了 90 - 99 ）
100，101，102，103，104，105，106，107，108 （移除了109）