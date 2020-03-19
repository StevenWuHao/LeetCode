<?php

/**
 * 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 *
 * https://leetcode-cn.com/problems/longest-palindrome/
 *
 * Class Solution
 */
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function longestPalindrome($s) {
        $len = strlen($s);

        if ($len <= 1) {
            return $len;
        }

        $arr = [];
        for ($i=0;$i<$len;$i++) {
            if (empty($arr[$s[$i]])) {
                $arr[$s[$i]] = 1;
            }else{
                $arr[$s[$i]]++;
            }
        }

        // 总长度
        $count = 0;
        // 字符数量是否是奇数
        $jishu_num = 0;
        // 字符种类
        $arr_count = count($arr);
        foreach ($arr as $k => $v) {
            if ($v%2 != 0 && $arr_count > 1) {
                if ($jishu_num > 0) {
                    if ($v - 1 <= 0) {
                        unset($arr[$k]);
                    }else{
                        $count += $v - 1;
                    }
                }else{
                    $count += $v;
                    $jishu_num++;
                }
            }else{
                $count += $v;
            }
        }

        return $count;
    }
}

$solution = new Solution();
$s = "abccccdd";
var_dump($solution->longestPalindrome($s));
die;