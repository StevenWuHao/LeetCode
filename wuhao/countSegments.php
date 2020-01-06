<?php

/**
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。

示例:

输入: "Hello, my name is John"
输出: 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-segments-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 *
 */
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function countSegments($s) {
        $len = strlen($s);

        $count = 0;
        for ($i=0;$i<$len;$i++) {
            if ($s[$i] != ' ' && ($i == $len - 1 || $s[$i+1] == ' ')) {
                $count++;
            }
        }

        return $count;
    }
}

$solution = new Solution();
$s = ", , , ,        a, eaefa";
var_dump($solution->countSegments($s));
die;