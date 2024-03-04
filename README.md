<div align = "center">

# TextScanner
</div>

## Description
Golang-Based Program for counting the number of words within a string, all based on user input. 

<div align = "center">

![img](https://i.imgur.com/PLIKx99.gif)
</div>

## Motivation
This program has been created for testing a different approach towards learning Golang: following the '[Question Driven Development](https://www.delenamalan.co.za/2021/2021-04-14-question-driven-development.html)' that aims towards learning thigs by doing.

## Concept         
Create an infinite shell until user has written "exit" on console:               
        1. Let user enter a string or paragraph        
        2. Let user search for a particular word that they want to find word count of     
        3. Print the number of times a word was mentioned in the entered string    

## How to do it? 
Creating Shell 
        - Can be done with a for loop
Scanning for input
        - Can be done with fmt.Scanln
Scanning for word-to-search
        - Can be done with fmt.Scanln
Finding word-count
        - For-Loop to range over the user input 

1. Let User Enter & Search String
Faced problem with scanner:
```
	fmt.Println("\nPlease enter the string!")
	fmt.Scanln("%s", &input)
```
- Program would end without taking user-input
- [Input not taken by fmt.Scanf or fmt.Scanln](https://stackoverflow.com/questions/69785464/input-not-taken-by-fmt-scanf-or-fmt-scanln)
- Solved: fmt Scan family scan space-separated tokens. Instead of %s (string), use %q (quoted string) )and enter string within "quotes" 

Instead we could have used bufio to provide a temporary storage:
```
        scanner := bufio.NewScanner(os.Stdin)
	input = scanner.Text()
	scanner.Scan()
	input = scanner.Text()
	fmt.Printf("The enter input is: %s\n", input)
```

- [Buffered vs unbuffered IO](https://stackoverflow.com/a/1450563/21819272)
- [Table Difference](https://gosamples.dev/read-user-input/#:~:text=kl%20mn%20op-,Use%20the%20fmt.,the%20end%20of%20each%20line.)
- Stop being distracted by theory all the time - Try to finish using whatever method you already know

2. Range over input and scan how many times specific word mentioned
    - Referred to [this guide](https://www.educative.io/answers/how-to-check-if-an-element-is-inside-a-slice-in-golang)
    - Basic blueprint:
```
	wordCount := 0
	for _, singleInput := range splitInput {
		if word == singleInput {
			wordCount++
		}
	}
	fmt.Printf("The KeyWord %s is found %d times\n", word, wordCount)
```
