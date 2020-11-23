<?php

/**
 *
 *给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * Class Solution
 */
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isPalindrome($s) {
        if ($s == '') {
            return true;
        }

        $low_s = strtolower($s);
        $len = strlen($low_s);
        $left = 0;
        $right = $len - 1;
        while ($left < $right) {
            if (!$this->isVerification($low_s[$left])) {
                $left++;
                continue;
            }elseif (!$this->isVerification($low_s[$right])) {
                $right--;
                continue;
            }elseif ($low_s[$left] != $low_s[$right]) {
                return false;
            }

            $left++;
            $right--;
        }

        return true;
    }

    private function isVerification ($str) {
        $ascii_str = ord($str);
        if (($ascii_str >= 48 && $ascii_str <= 57) || ($ascii_str >= 65 && $ascii_str <= 90) || ($ascii_str >= 97 && $ascii_str <= 122)) {
            return true;
        }
        return false;
    }
}

$solution = new Solution();
$s = "race a car";
var_dump($solution->isPalindrome($s));
die;