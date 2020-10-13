<?php

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLastWord($s) {
        $len = 0;

        if (strpos($s,' ') !== false) {
            $s_arr = explode(' ',$s);
            $new_arr = [];
            foreach ($s_arr as $value) {
                if (!$value) continue;
                $new_arr[] = $value;
            }
            $str = array_pop($new_arr);
            $len = strlen($str);
        }elseif (!empty($s) && $s != " ") {
            $len = strlen($s);
        }

        return $len;
    }
}

$solution = new Solution();
$s = "          ";
var_dump($solution->lengthOfLastWord($s));die;