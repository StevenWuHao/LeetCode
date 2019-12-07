<?php

class Solution {

    /**
     * @param String[] $strs
     * @return String
     */
    function longestCommonPrefix($strs) {
        // 如果数组不存在返回空字符串
        if (!$strs) {
            return '';
        }

        $strs_count = count($strs);
        // 如果数组中只有一个元素，则返回该元素
        if ($strs_count == 1) {
            return $strs[0];
        }

        $first_str_len = strlen($strs[0]); // 数组第一个元素的字符个数
        $common_str = '';
        $is_finish = false;
        // 循环数组第一个元素的所有字符
        for ($i = 0; $i < $first_str_len; $i++) {
            $is_common = false;
            // 用数组的第一个元素的每个字符和其他数组元素的每个字符匹配
            for ($j = 1; $j < $strs_count; $j++) {

                if (empty($strs[$j][$i])) {
                    $is_common = false;
                    $is_finish = true;
                    break;
                }

                if ($strs[0][$i] == $strs[$j][$i]) {
                    $is_common = true;
                } else {
                    $is_common = false;
                    break;
                }
            }
            if ($is_common) {
                $common_str .= $strs[0][$i];
            }
            if ($is_finish) {
                break;
            }
        }

        return $common_str;
    }
}

$solution = new Solution();
$strs = ["aca","cba"];
var_dump($solution->longestCommonPrefix($strs));die;