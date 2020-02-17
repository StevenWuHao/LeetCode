<?php

/**
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
 

注意事项：您可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * Class Solution
 */
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function firstUniqChar($s) {
        // 计算长度
        $count = strlen($s);
        // 循环字符串
        for ($i=0;$i<$count;$i++) {
            // 默认没有重复
            $is_repetition = false;
            // 依次进行比较
            for ($j=0;$j<$count;$j++) {
                // 如果是当前字符则跳过
                if ($j == $i) {
                    continue;
                }
                // 有重复则跳过之后所有的比较
                if ($s[$i] == $s[$j]) {
                    $is_repetition = true;
                    break;
                }
            }
            // 没有重复则跳过之后所有的比较并且返回字符串索引
            if (! $is_repetition) {
                return $i;
            }
        }

        return -1;
    }
}

$solution = new Solution();
$s = "leetcode";
var_dump($solution->firstUniqChar($s));
die;