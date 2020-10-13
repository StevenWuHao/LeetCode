<?php

/**
 *
 * 686. 重复叠加字符串匹配
给定两个字符串 A 和 B, 寻找重复叠加字符串A的最小次数，使得字符串B成为叠加后的字符串A的子串，如果不存在则返回 -1。

举个例子，A = "abcd"，B = "cdabcdab"。

答案为 3， 因为 A 重复叠加三遍后为 “abcdabcdabcd”，此时 B 是其子串；A 重复叠加两遍后为"abcdabcd"，B 并不是其子串。

注意:

A 与 B 字符串的长度在1和10000区间范围内。
 * Class Solution
 */
class Solution {

    /**
     * @param String $A
     * @param String $B
     * @return Integer
     */
    function repeatedStringMatch($A, $B) {
        $len_a = strlen($A);
        $len_b = strlen($B);
        for ($i=0;$i<$len_a;$i++) {
            // A字符串匹配B字符串第一个字符，如果相同则继续匹配，否则跳过此次循环
            if ($A[$i] == $B[0]) {
                $k = $i;
                // 重叠次数
                $count = 1;
                $j = 0;
                while ($A[$k] == $B[$j]) {
                    $k++;
                    $j++;
                    // 如果B字符串的索引大于等于B字符串的长度，则说明匹配成功
                    if ($j >= $len_b) {
                        return $count;
                    }
                    // A字符串循环匹配B字符串，如abc和abca匹配，
                    // 前三位匹配完后用B字符串的第四位匹配A字符串的第一位
                    if ($k >= $len_a) {
                        $k = 0;
                        $count++;
                    }
                }
            }
        }
        return -1;
    }
}

$solution = new Solution();
$A = "abc";
$B = "cabcabca";
//$A = "abcd";
//$B = "cdabcdvvabcdab";
var_dump($solution->repeatedStringMatch($A,$B));
die;