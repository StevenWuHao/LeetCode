<?php

/**
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * Class Solution
 */
class Solution
{
    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s) {
        $slen = strlen($s);

        if ($slen <2 && $slen != 0) {
            return false;
        }

        $arr = [];
        $is_true = true;
        for ($i=0;$i<$slen;$i++) {
            if ($s[$i] == '{' || $s[$i] == '(' || $s[$i] == '[') {
                $arr[] = $s[$i];
            }

            $count = count($arr);
            if ($count < 1) {
                $is_true = false;
                break;
            }

            switch ($s[$i]) {
                case '}':
                    if ($arr[$count-1] != '{') {
                        $is_true = false;
                    }
                    array_pop($arr);
                    break;
                case ')':
                    if ($arr[$count-1] != '(') {
                        $is_true = false;
                    }
                    array_pop($arr);
                    break;
                case ']':
                    if ($arr[$count-1] != '[') {
                        $is_true = false;
                    }
                    array_pop($arr);
                    break;
            }

            if (!$is_true) {
                break;
            }
        }

        if (count($arr) > 0) {
            return false;
        }

        return $is_true;
    }
}