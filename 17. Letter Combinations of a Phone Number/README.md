# 17. Letter Combinations of a Phone Number

|Category|Difficulty|Tags|
|:-:|:-:|:-:|
|algorithms|Medium|string, backtracking|

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

![Letter Combinations of a Phone Number]("https://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png")

**Example 1:**

``` text
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**Example 2:**

``` text
Input: digits = ""
Output: []
```

**Example 3:**

``` text
Input: digits = "2"
Output: ["a","b","c"]
```

**Constraints:**

+ $0 <= digits.length <= 4$
+ `digits[i]` is a digit in the range `['2', '9']`.
