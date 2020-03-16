<?php

/**
 *
 * 面试题 01.06. 字符串压缩
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

输入："aabcccccaaa"
输出："a2b1c5a3"
示例2:

输入："abbccd"
输出："abbccd"
解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：

字符串长度在[0, 50000]范围内。
 *
 * https://leetcode-cn.com/problems/compress-string-lcci/
 *
 * Class Solution
 */
class Solution {

    /**
     * @param String $S
     * @return String
     */
    function compressString($S) {
        $len = strlen($S);

        if ($len <= 1) {
            return $S;
        }

        // 新字符串
        $str = '';
        // 相同字母
        $tmp_s = '';
        // 相同字母数量
        $count = 1;
        for ($i=0;$i<$len;$i++) {
            if (! $tmp_s) {
                $tmp_s = $S[$i];
            }

            if ($i > 0) {
                if ($S[$i] != $S[$i-1]) {
                    $str .= $tmp_s.$count;
                    $tmp_s = $S[$i];
                    $count = 1;
                    if ($i == $len - 1) {
                        $str .= $tmp_s.$count;
                    }
                }elseif ($i == $len - 1) {
                    $count++;
                    $str .= $tmp_s.$count;
                }else{
                    $count++;
                }
            }
        }

        $len_str = strlen($str);

        return $len_str >= $len ? $S : $str;
    }
}

$solution = new Solution();
$S = "abbccd";
var_dump($solution->compressString($S));
die;