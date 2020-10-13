<?php

/**
 *
 * 面试题 01.02. 判定是否互为字符重排
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false
说明：

0 <= len(s1) <= 100
0 <= len(s2) <= 100
 *
 * https://leetcode-cn.com/problems/check-permutation-lcci/
 *
 * Class Solution
 */
class Solution {

    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    function CheckPermutation($s1, $s2) {
        $s1_len = strlen($s1);
        $s2_len = strlen($s2);

        if ($s1_len != $s2_len) {
            return false;
        }

        $s1_arr = [];
        $s2_arr = [];
        for ($i=0;$i<$s1_len;$i++) {
            if (empty($s1_arr[$s1[$i]])) {
                $s1_arr[$s1[$i]] = 1;
            }else{
                $s1_arr[$s1[$i]]++;
            }

            if (empty($s2_arr[$s2[$i]])) {
                $s2_arr[$s2[$i]] = 1;
            }else{
                $s2_arr[$s2[$i]]++;
            }
        }

        foreach ($s1_arr as $key => $v) {
            if (!empty($s2_arr[$key]) && $v == $s2_arr[$key]) {
                unset($s1_arr[$key]);
            }
        }

        if (count($s1_arr) > 0) {
            return false;
        }

        return true;
    }
}

$solution = new Solution();
$s1 = "abc";
$s2 = "bad";
var_dump($solution->CheckPermutation($s1, $s2));
die;