There is a sequence of words in CamelCase as a string of letters, `s` , having the following properties:

* It is a concatenation of one or more words consisting of English letters.
* All letters in the first word are lowercase.
* For each of the subsequent words, the first letter is uppercase and rest of the letters are lowercase.

Given `s` , determine the number of words in .

__Example__

`s = oneTwoThree`

There are `3` words in the string: 'one', 'Two', 'Three'.

__Function Description__

Complete the camelcase function in the editor below.

camelcase has the following parameter(s):
* string s: the string to analyze

__Returns__
* int: the number of words in `s`

__Input Format__

A single line containing string `s`.

__Constraints__
* `1 <= length of s <= 10^5`

__Sample Input__

`saveChangesInTheEditor`

__Sample Output__

`5`

__Explanation__

String  contains five words:
* save
* Changes
* In
* The
* Editor
