<?php

class Solution {

    /**
     * @param Integer $n
     * @return String
     */
    function countAndSay($n) {
        // 初始值设置为一
        $str = '1';

        for ($i=0;$i<$n-1;$i++) {
            $arr = [];
            $len = strlen($str);
            $count = 1;
            if ($len<2) {
                $arr[] = $count;
                $arr[] = $str;
            }else{
                for ($j=1;$j<$len;$j++) {
                    if ($str[$j] == $str[$j-1]) {
                        $count++;
                    }else{
                        $arr[] = $count;
                        $arr[] = $str[$j-1];
                        $count = 1;
                    }
                    if ($j == ($len-1)) {
                        $arr[] = $count;
                        $arr[] = $str[$j];
                    }
                }
            }
            $str = '';
            foreach ($arr as $v) {
                $str .= $v;
            }
        }

        return $str;
    }
}

$solution = new Solution();
var_dump($solution->countAndSay(1));die;