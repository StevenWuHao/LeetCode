<?php

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function romanToInt($s) {
        // 以罗马数字为键，阿拉伯数字为值
        $romanArr = [
            'I' => 1,
            'V' => 5,
            'X' => 10,
            'L' => 50,
            'C' => 100,
            'D' => 500,
            'M' => 1000,
            'IV' => 4,
            'IX' => 9,
            'XL' => 40,
            'XC' => 90,
            'CD' => 400,
            'CM' => 900
        ];

        // 拆分字符串，每个字符都为一个数组的值
        $split = str_split($s,1);
        $splitCount = count($split);

        $num = 0;
        $i = 0;
        while ($i < $splitCount) {
            // 如果当前是最后一位则不去匹配特殊数字
            $joint = '';
            if($i != ($splitCount - 1)){
                $joint = $split[$i].$split[$i+1];
            }

            // 如果是特殊数字直接跳过下一个字符
            if (array_key_exists($joint, $romanArr)) {
                $num += $romanArr[$joint];
                $i += 2;
            } else {
                $num += $romanArr[$split[$i]];
                $i++;
            }
        }

        return $num;
    }
}

$solution = new Solution();
var_dump($solution->romanToInt('MCMXCIV'));
die;