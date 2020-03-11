<?php

/**
 *
 * 1013. 将数组分成和相等的三个部分
给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。



示例 1：

输出：[0,2,1,-6,6,-7,9,1,2,0,1]
输出：true
解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
示例 2：

输入：[0,2,1,-6,6,7,9,-1,2,0,1]
输出：false
示例 3：

输入：[3,3,6,5,-2,2,5,1,-9,4]
输出：true
解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4


提示：

3 <= A.length <= 50000
-10^4 <= A[i] <= 10^4
 *
 * https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
 * Class Solution
 */
class Solution {

    /**
     *
     * 思路：先计算所有值的总和，再取平均数，如果不能被3整除则为false
     * 依次相加，如果等于平均数则重新计算并且块数加1，如果小于3块或者最后没有清零，则为false
     * 如果平均数为0，则块数可以大于3
     *
     * @param Integer[] $A
     * @return Boolean
     */
    function canThreePartsEqualSum($A) {
        // 计算总和
        $sum = 0;
        foreach ($A as $a) {
            $sum += $a;
        }

        // 取余
        if ($sum%3 != 0) {
            return false;
        }

        // 计算平均值 并计算块数
        $avg = $sum/3;
        $chunk_num = 0;
        $chunk_count = 0;
        foreach ($A as $a) {
            $chunk_num += $a;
            if ($chunk_num == $avg) {
                $chunk_num = 0;
                $chunk_count++;
            }
        }

        // 如果清零并且块数等于3 或者 平均值等于0 块数大于3 则为正确
        return $chunk_num == 0 && ($chunk_count == 3 || $avg == 0 && $chunk_count >= 3);
    }
}

$solution = new Solution();
$A = [3,3,6,5,-2,2,5,1,-9,4];
var_dump($solution->canThreePartsEqualSum($A));
die;