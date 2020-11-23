<?php

/**
 * 28. 实现 strStr()
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 * Class Solution
 */
class Solution {

    /**
     * @param String $haystack
     * @param String $needle
     * @return Integer
     */
    function strStr($haystack, $needle) {
        $haystack_len = strlen($haystack);
        $needle_len = strlen($needle);

        if ($haystack_len < $needle_len) {
            return -1;
        }elseif ($haystack == $needle || $needle == '') {
            return 0;
        }

        $j = 0;
        $key = 0;
        $first_key_exist = false;
        for ($i=0;$i<$haystack_len;$i++) {
            if ($haystack[$i] != $needle[$j]) {
                if ($first_key_exist) {
                    $i = $key;
                }
                $j = 0;
                $key = 0;
                $first_key_exist = false;
                continue;
            }

            if ($haystack[$i] == $needle[0] && !$first_key_exist) {
                $key = $i;
                $first_key_exist = true;
                $j++;
            }elseif ($haystack[$i] == $needle[$j] && $first_key_exist) {
                $j++;
            }

            if ($j == $needle_len && $first_key_exist) {
                return $key;
            }
        }

        return -1;
    }
}

$solution = new Solution();
$haystack = "mississippi";
$needle = "pi";
var_dump($solution->strStr($haystack, $needle));
die;