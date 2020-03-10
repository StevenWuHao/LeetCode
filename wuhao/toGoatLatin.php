<?php

/**
 *
 * 824. 山羊拉丁文
给定一个由空格分割单词的句子 S。每个单词只包含大写或小写字母。

我们要将句子转换为 “Goat Latin”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。

山羊拉丁文的规则如下：

如果单词以元音开头（a, e, i, o, u），在单词后添加"ma"。
例如，单词"apple"变为"applema"。

如果单词以辅音字母开头（即非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
例如，单词"goat"变为"oatgma"。

根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从1开始。
例如，在第一个单词后添加"a"，在第二个单词后添加"aa"，以此类推。
返回将 S 转换为山羊拉丁文后的句子。

示例 1:

输入: "I speak Goat Latin"
输出: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
示例 2:

输入: "The quick brown fox jumped over the lazy dog"
输出: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
说明:

S 中仅包含大小写字母和空格。单词间有且仅有一个空格。
1 <= S.length <= 150。
 * Class Solution
 */
class Solution {

    /**
     * @param String $S
     * @return String
     */
    function toGoatLatin($S) {
        // 元音
        $vowel = ['a','e','i','o','u','A','E','I','O','U'];
        // 字符串长度
        $len = strlen($S);
        // 单词
        $word = '';
        // 转换后的字符串
        $str = '';
        // 单词索引
        $count = 1;
        for ($i=0;$i<$len;$i++) {
            if ($S[$i] != ' ') {
                $word .= $S[$i];
            }

            if ($S[$i] == ' ' || $i == ($len - 1)) {
                // 首字母元音直接在单词最后加ma
                if (in_array($word[0], $vowel)) {
                    $word .= 'ma';
                }else{
                    // 将第一个字符移到单词的最后并拼接ma
                    $word_len = strlen($word);
                    $first_str = $word[0];
                    for ($j=0;$j<$word_len-1;$j++) {
                        $word[$j] = $word[$j+1];
                    }

                    $word[$word_len-1] = $first_str;
                    $word .= 'ma';
                }

                // 根据单词的索引增加a的个数
                for ($k=0;$k<$count;$k++) {
                    $word .= 'a';
                }

                $str .= $word;
                if ($S[$i] == ' ') {
                    $str .= $S[$i];
                }
                $word = '';

                $count++;
            }
        }

        return $str;
    }
}

$solution = new Solution();
$S = "The quick brown fox jumped over the lazy dog";
var_dump($solution->toGoatLatin($S));
die;