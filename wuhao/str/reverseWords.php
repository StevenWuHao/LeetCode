<?php

/**
 * 557. 反转字符串中的单词 III
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 * Class Solution
 */
class Solution {

    /**
     * @param String $s
     * @return String
     */
    function reverseWords($s) {
        $len = strlen($s);

        $str = '';
        $word = '';
        for ($i=0;$i<$len;$i++) {
            if ($s[$i] != ' ') {
                $word .= $s[$i];
            }

            if ($s[$i] == ' ' || $i == $len - 1) {
                $word_len = strlen($word);
                $l = 0;
                $r = $word_len - 1;
                while ($l <= $r) {
                    $tmp = $word[$l];
                    $word[$l] = $word[$r];
                    $word[$r] = $tmp;
                    $l++;
                    $r--;
                }
                $str .= $word;
                $word = '';
                if ($s[$i] == ' ') {
                    $str .= $s[$i];
                }
            }
        }

        return $str;
    }
}

$solution = new Solution();
$s = "Let's take LeetCode contest";
var_dump($solution->reverseWords($s));
die;