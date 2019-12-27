<?php

class Solution {

    /**
     * @param String $a
     * @param String $b
     * @return String
     */
    function addBinary($a, $b) {
        // 计算两个字符串的长度
        $count1 = strlen($a);
        $count2 = strlen($b);

        // 如果a大于b，则循环a长度的次数，并且在b的最前面填充字符0，使其和a相同长度，反之亦然
        if ($count1 > $count2) {
            $count = $count1;
            $count3 = $count1-$count2;
            $str = str_repeat('0',$count3);
            $b = $str.$b;
        }else{
            $count = $count2;
            $count3 = $count2-$count1;
            $str = str_repeat('0',$count3);
            $a = $str.$a;
        }

        $d = 0;
        $sum = '';
        for ($i=1;$i<=$count;$i++) {
            // a和b最后一个字符相加，结果大于1则在字符末尾拼接字符0
            $c = ($a[$count-$i] + $b[$count-$i]);
            if ($c > 1) {
                $sum .= 0 + $d;
                $d = 1;
            }else{
                if (($c + $d) > 1) {
                    $d = 1;
                    $sum .= 0;
                }else{
                    $sum .= $c + $d;
                    $d = 0;
                }
            }

            if ($i == $count && $d > 0) {
                $sum .= $d;
            }
        }
        return strrev($sum);
    }
}

$solution = new Solution();
$a = '11';
$b = '11';
var_dump($solution->addBinary($a,$b));die;