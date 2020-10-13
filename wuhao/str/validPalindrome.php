<?php

/**
 *
 * 680. 验证回文字符串 Ⅱ
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:

字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
 * Class Solution
 */
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function validPalindrome($s) {
        $len = strlen($s);
        // 是否回文字符串
        $is_palindrome = true;
        // 删除字符次数
        $delete_count = 0;
        for ($i=0,$j=$len-1;$i<$j;$i++,$j--) {
            // 如果两边不相等并且没有删除过字符
            if ($s[$i] != $s[$j]) {
                if ($delete_count < 1) {
                    // 如果左边移一位和右边当前索引相等并且后一位的匹配也相等则移左边，否则右边移动一位
                    // 如果右边移一位和左边当前索引不相等，则不是回文字符
                    if ($s[$i+1] == $s[$j] && $s[$i+2] == $s[$j-1]) {
                        $delete_count++;
                        $j++;
                        continue;
                    }elseif ($s[$i] == $s[$j-1]) {
                        $delete_count++;
                        $i--;
                        continue;
                    }else{
                        $is_palindrome = false;
                        break;
                    }
                }else{
                    $is_palindrome = false;
                    break;
                }
            }
        }

        return $is_palindrome;
    }
}

$solution = new Solution();
$s = "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga";
var_dump($solution->validPalindrome($s));
die;