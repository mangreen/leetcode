# [277. Find the Celebrity](https://leetcode.ca/2016-09-01-276-Paint-Fence/)

## Level
Medium

## Description
You are painting a fence of n posts with k different colors. You must paint the posts following these rules:

Every post must be painted exactly one color.
There cannot be three or more consecutive posts with the same color.
Given the two integers n and k, return the number of ways you can paint the fence.

### Ex1

```bash
Input: n = 3, k = 2

Output: 6

Explanation: All the possibilities are shown.
Note that painting all the posts red or all the posts green is invalid because there cannot be three posts in a row with the same color.
```

### Ex2

```bash
Input: n = 1, k = 1

Output: 1
```

### Ex3

```bash
Input: n = 7, k = 2
Output: 42
```

## Constraints

- 1 <= n <= 50
- 1 <= k <= 105
- The testcases are generated such that the answer is in the - range [0, 231 - 1] for the given n and k.
