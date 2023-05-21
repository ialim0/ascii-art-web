# Welcome to our PrintAscii README

### PrintAscii Fonction
The function takes as input a string input and an array of string arrays tab. The function parses each line of the input string, divides it into words separated by \n, and displays each character of each line as ASCII art using the tab array.

The count variable is used to count the number of non-empty lines in input. The variable final is used to store the resulting string containing the ASCII art.

The function uses two nested for loops: the first loop traverses each line of the input string and the second loop traverses each character of each line.

If the line is not empty, the function checks if the character is valid (i.e. if it is between ASCII codes 32 and 126 inclusive) and displays the corresponding ASCII art by retrieving the data from the tab array.

If the line is empty but there was at least one non-empty line before, the function adds an empty line.

The function returns the resulting string containing the ASCII art corresponding to the input string.

## Thank You :smile:

![](https://www.zone01dakar.sn/wp-content/uploads/2022/04/zoneBleu.png)