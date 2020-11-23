<?php

/**
 * 在一个「平衡字符串」中，'L' 和 'R' 字符的数量是相同的。

给出一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。

返回可以通过分割得到的平衡字符串的最大数量。

 

示例 1：

输入：s = "RLRRLLRLRL"
输出：4
解释：s 可以分割为 "RL", "RRLL", "RL", "RL", 每个子字符串中都包含相同数量的 'L' 和 'R'。
示例 2：

输入：s = "RLLLLRRRLR"
输出：3
解释：s 可以分割为 "RL", "LLLRRR", "LR", 每个子字符串中都包含相同数量的 'L' 和 'R'。
示例 3：

输入：s = "LLLLRRRR"
输出：1
解释：s 只能保持原样 "LLLLRRRR".
 

提示：

1 <= s.length <= 1000
s[i] = 'L' 或 'R'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-a-string-in-balanced-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function balancedStringSplit($s) {
        $slen = strlen($s);
        if ($slen < 2) {
            return 0;
        }

        // 待匹配字符串
        $left_str = '';
        // 连续相同的数量
        $left_count = 0;
        // 平衡数量
        $count = 0;
        $arr = str_split($s,1);
        foreach ($arr as $val) {
            // 第一个字符串
            if ($val == 'R' && empty($left_str)) {
                $left_str = 'R';
            }elseif ($val == 'L' && empty($left_str)) {
                $left_str = 'L';
            }

            // 如果和第一个字符相同，相同数量就加1，否则抵消1位
            if ($left_str == $val) {
                $left_count++;
            }else{
                $left_count--;
                // 如果相同数量为0，说明都抵消了，平衡数量就加1
                if ($left_count <= 0) {
                    $left_str = '';
                    $count++;
                }
            }
        }

        return $count;
    }
}

$solution = new Solution();
$str = "LLLLRRRR";
var_dump($solution->balancedStringSplit($str));die;