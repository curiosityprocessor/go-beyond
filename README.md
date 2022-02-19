# ready to GO
Toy project for studying Go lang

## Teck Stack
GO (1.17.6)

## Table of Contents
### 1. GO-101
Go basics with **Tucker의 Go 언어 프로그래밍**  

|package|description|note|
|-------|-----------|---|
|ch00|write, build, run a .go file| |
|ch04|variables| |
|ch05|fmt|standard I/O|
|ch06|operators| |
|ch07|functions|multi-return types, named return types |
|ch08|constants|iota, typless constants |
|ch09|if|init statement|
|ch10|switch|init statement, break, fallthrough |
|ch11|for loop|optional init, conditional, post statements |
|ch12|array| |
|ch13|struct|nameless nested structs, memory alignment |
|ch14|pointer|pass by value of pointer i.e. pass by reference|
|ch15|string| |
|ch16|package| |
|ch18|slice|slice vs array, append, make |

### 2. guess_my_number
Guess a random number  
using  
* math/rand, time package
* standard I/O
* for loop


## adopted conventions
#### packages
* snake_case
* no starting numbers, special charaters

#### functions
* exporting functions: UpperCamelCase (PascalCase)
* non-exporting functions: camelCase
* compound words for clarity

#### variables
* camelCase
* exporting constants in UpperCamelCase
* has, is, can prefix for bool types
* preferably short, single words

#### structs
* exporting structs: UpperCamelCase
* non-exporting stucts: camelCase