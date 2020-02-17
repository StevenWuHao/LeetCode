<?php

class Solution {

    /**
     * @param String[] $chars
     * @return Integer
     */
    function compress(&$chars) {
        $count = count($chars);

        for ($i=0;$i<$count;$i++) {
            $j = $i + 1;
            $k = 0;
            for (;$j<$count;$j++) {
                if ($chars[$i] == $chars[$j]) {
                    if ($k) {
                        $chars[$k]++;
                        unset($chars[$j]);
                    }else{
                        $k = $j;
                        $chars[$k] = 2;
                    }
                }else{
                    break;
                }
            }

            array_push($chars, $chars[$i]);
            unset($chars[$i]);

            if ($k) {
                $chars[$k] = (string) $chars[$k];
                $len = strlen($chars[$k]);
                for ($m=0;$m<$len;$m++) {
                    array_push($chars,$chars[$k][$m]);
                }
                unset($chars[$k]);
            }

            $i = $j - 1;
        }

        return count($chars);
    }
}

$solution = new Solution();
$chars = ["a","a","a","a","a","a","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","b","c","c","c","c","c","c","c","c","c","c","c","c","c","c"];
var_dump($solution->compress($chars));
var_dump($chars);die;