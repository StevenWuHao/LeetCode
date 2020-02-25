<?php

/**
 * 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * Class Solution
 */
class Solution {

    /**
     * @param String $s
     * @return String
     */
    function reverseVowels($s) {
        $vowel_arr = ['a','e','i','o','u','A','E','I','O','U'];

        $len = strlen($s);
        $j = $len - 1;
        for ($i=0;$i<$j;$i++) {
            if (!in_array($s[$i], $vowel_arr)) {
                continue;
            }elseif (!in_array($s[$j], $vowel_arr)) {
                while (true) {
                    $j--;
                    if (in_array($s[$j], $vowel_arr)) {
                        break;
                    }
                }
            }

            $tmp = $s[$i];
            $s[$i] = $s[$j];
            $s[$j] = $tmp;
            $j--;
        }

        return $s;
    }
}

$solution = new Solution();
$s = "aA";
var_dump($solution->reverseVowels($s));
die;