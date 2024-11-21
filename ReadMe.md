Diamond Pattern Generator

        *
       ***
      *****
     *******
    *********
     *******
      *****
       ***
        *

Description

The Diamond Pattern Generator is a simple Go program that takes two arguments — the height and width of a diamond — and prints an ASCII diamond pattern to the console.

This project demonstrates how to handle command-line arguments, validate input, and generate dynamic patterns in Go. It’s a great beginner project for practicing conditional logic, loops, and basic string manipulation.
Features

    Customizable Dimensions: Specify the height and width of the diamond.
    Input Validation: Ensures that the height and width are odd numbers.
    Error Handling: Provides clear error messages for invalid inputs.
    Elegant Design: Produces symmetric diamond patterns.

How to Use
Run the Program

    Make sure you have Go installed on your system.

    Clone this repository or download the file.

    Navigate to the project directory and run:

    go run main.go [Height] [Width]

Example

go run main.go 7 5

Output:

  *  
 *** 
*****
 *** 
  *  

Valid Inputs

    Both height and width must be odd numbers.
    Example: 5 5, 7 9, etc.

Invalid Inputs

    Even dimensions:
        Example: 4 6.
        Error: Height and width must be odd numbers to form a proper diamond.

    Missing or extra arguments:
        Example: go run main.go 7.
        Error: Invalid number of arguments.

    Non-numeric inputs:
        Example: go run main.go A B.
        Error: Height and width must be integers.

Sample Output
Diamond (Height = 7, Width = 7)

   *   
  ***  
 ***** 
******* 
 ***** 
  ***  
   *   

About the Code
Functions

    main:
        Handles argument validation and calls the diamond generator.
    Usage:
        Displays usage instructions for the program.
    checkArguments:
        Validates the number and type of input arguments.
    printDiamond:
        Generates and prints the diamond pattern.
    repeatChar:
        Repeats a character n times for pattern generation.
    abs:
        Returns the absolute value of a number.

ASCII Art Inspiration


    /\    
   /  \   
  /    \  
 /      \ 
/        \
\        /
 \      / 
  \    /  
   \  /   
    \/    

Author

Crafted with ❤️ by 'Anass Lazaar' a passionate Go developer.

    "Code, like diamonds, should be simple yet timeless."
