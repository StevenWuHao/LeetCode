<?php


/**
 *
 * 859. 亲密字符串
给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。



示例 1：

输入： A = "ab", B = "ba"
输出： true
示例 2：

输入： A = "ab", B = "ab"
输出： false
示例 3:

输入： A = "aa", B = "aa"
输出： true
示例 4：

输入： A = "aaaaaaabc", B = "aaaaaaacb"
输出： true
示例 5：

输入： A = "", B = "aa"
输出： false


提示：

0 <= A.length <= 20000
0 <= B.length <= 20000
A 和 B 仅由小写字母构成。
 *
 *链接：https://leetcode-cn.com/problems/buddy-strings/
 * Class Solution
 */
class Solution {

    /**
     * @param String $A
     * @param String $B
     * @return Boolean
     */
    function buddyStrings($A, $B) {
        $len_a = strlen($A);
        $len_b = strlen($B);

        if ($len_a != $len_b || $len_a < 2) {
            return false;
        }

        if ($len_a == 2) {
            if ($A[0] == $B[1] && $A[1] == $B[0]) {
                return true;
            }
        }

        // 交换次数
        $change_count = 0;
        // 需要交换字符的下标
        $tmp_index = -1;
        // 记录每个字符出现的次数
        $arrs = [];
        for ($i=0;$i<$len_a;$i++) {
            if ($A[$i] != $B[$i]) {
                // 如果已经交换过则返货false
                if ($change_count > 0) {
                    return false;
                }
                // 如果没有记录过下标则记录下标 否则用当前的字符和之前匹配失败的字符比较
                if ($tmp_index < 0) {
                    $tmp_index = $i;
                }else{
                    if ($A[$tmp_index] == $B[$i] && $A[$i] == $B[$tmp_index]) {
                        $change_count++;
                    }else{
                        return false;
                    }
                }
            }

            // 记录每个字符出现的次数
            if (empty($arrs[$A[$i]])) {
                $arrs[$A[$i]] = 1;
            }else{
                $arrs[$A[$i]]++;
            }
        }

        // 如果有相同的字符则返回TRUE
        foreach ($arrs as $arr) {
            if ($arr > 1) {
                return true;
            }
        }

        // 如果没有换过位置则返回FALSE
        if ($change_count < 1) {
            return false;
        }

        return true;
    }
}

$solution = new Solution();
$A = "abcaa";
$B = "abcbb";
var_dump($solution->buddyStrings($A,$B));
die;