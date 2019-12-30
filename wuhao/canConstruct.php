<?php

/**
 * 给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。

(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)

注意：

你可以假设两个字符串均只含有小写字母。

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ransom-note
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 *
 * Class Solution
 */
class Solution {

    /**
     * @param String $ransomNote
     * @param String $magazine
     * @return Boolean
     */
    function canConstruct($ransomNote, $magazine) {
        $len1 = strlen($ransomNote);
        $len2 = strlen($magazine);

        if (empty($ransomNote) || empty($magazine)) {
            return true;
        }elseif(empty($ransomNote)) {
            return true;
        }

        if ($len1 > $len2) {
            return false;
        }

        $is_true = false;
        $arr = [];
        for ($i=0;$i<$len1;$i++) {
            $is_true = false;
            for ($j=0;$j<$len2;$j++) {
                if ($ransomNote[$i] == $magazine[$j]) {
                    // 查看字符是否已匹配过
                    $is_exist = false;
                    foreach ($arr as $value) {
                        if ($j == $value) {
                            $is_exist = true;
                            break;
                        }
                    }
                    if ($is_exist) continue;
                    $arr[] = $j;
                    $is_true = true;
                    break;
                }
            }
            if (!$is_true) {
                break;
            }
        }

        return $is_true;
    }
}

$solution = new Solution();
// "fffbfg"
// "effjfggbffjdgbjjhhdegh"
$a = 'fffbfg';
$b = 'effjfggbffjdgbjjhhdegh';
var_dump($solution->canConstruct($a,$b));die;