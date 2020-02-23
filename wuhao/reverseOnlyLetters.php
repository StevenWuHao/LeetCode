<?php

/**
 * 917. 仅仅反转字母
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。



示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"


提示：

S.length <= 100
33 <= S[i].ASCIIcode <= 122
S 中不包含 \ or "
 * Class Solution
 */
class Solution {

    /**
     * @param String $S
     * @return String
     */
    function reverseOnlyLetters($S) {
        $len = strlen($S);
        $j = $len - 1;
        for ($i=0;$i<=$j;$i++) {
            $ord_l = ord($S[$i]);
            $ord_r = ord($S[$j]);
            if ($ord_l < 65 || ($ord_l > 90 && $ord_l < 97) || $ord_l > 122) {
                if ($ord_r < 65 || ($ord_r > 90 && $ord_r < 97) || $ord_r > 122) {
                    $j--;
                }
                continue;
            }

            while (true) {
                if (($ord_r >= 65 && $ord_r <= 90) || ($ord_r >= 97 && $ord_r <= 122)) {
                    break;
                }
                $j--;
                $ord_r = ord($S[$j]);
            }

            $tmp = $S[$i];
            $S[$i] = $S[$j];
            $S[$j] = $tmp;

            $j--;
        }

        return $S;
    }
}

$solution = new Solution();
$S = "Test1ng-Leet=code-Q!";
var_dump($solution->reverseOnlyLetters($S));
die;