<?php

class Solution {

    /**
     * @param String $paragraph
     * @param String[] $banned
     * @return String
     */
    function mostCommonWord($paragraph, $banned) {
        $len = strlen($paragraph);
        $paragraph = strtolower($paragraph);

        $arr = [];
        $str = '';
        $word = '';
        $count = 0;
        for ($i=0;$i<$len;$i++) {
            $ord = ord($paragraph[$i]);
            if ($ord >= 97 && $ord <= 122) {
                $str .= $paragraph[$i];
            }else{
                if (!empty($str) && !in_array($str,$banned)) {
                    if (empty($arr[$str])) {
                        $arr[$str] = 1;
                    }else{
                        $arr[$str]++;
                    }
                    if ($arr[$str] > $count) {
                        $count = $arr[$str];
                        $word = $str;
                    }
                }
                $str = '';
                continue;
            }

            if ($i == ($len - 1) && !in_array($str,$banned)) {
                if (empty($arr[$str])) {
                    $arr[$str] = 1;
                }else{
                    $arr[$str]++;
                }

                if ($arr[$str] > $count) {
                    $count = $arr[$str];
                    $word = $str;
                }
            }
        }

        return $word;
    }
}