<h1>Smallest alphabet string</h1>
Write a function smallestStr that, given a string S consisting of N characters, returns the alphabetically smallest string that can be obtained by removing exactly one letter from S.<br>
<br>
<b>Examples:</b>
<br>
<br>
1- Given S = "acb", by removing one letter, you can obtain "ac", "ab", or "cb". Your function should return "ab" (after removing 'c') since it's alphabetically smaller than "ac" and "bc".<br>
<br>
2- Given S = "hot", your function should return "ho", which is alphabetically smaller than "ht" and "ot".
<br>
<br>
3- Given S = "bbbbb", your function should return "bbbb". Any occurrence of 'b' can be removed.<br>
<br>
Write an <b>efficient</b> algorithm for the following assumptions:<br>
<br>
<ul>
	<li>N is an integer within the range [2 - 100000];</li>
	<li>string S consists only of lowercase letters (a-z).</li>
</ul> 