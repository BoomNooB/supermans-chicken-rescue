
# Problem 2: Superman's Chicken Rescue

## TechStack
All are written in **Go**,
Please use **Go version** at least **1.22**

## How to run

To get started, follow these steps:

1. **Clone the Repository:**
   ```
   git clone https://github.com/BoomNooB/supermans-chicken-rescue.git
   cd supermans-chicken-rescue
   ```
   

4. **Running a makefile:**\
   Simply just type this to your terminal after `cd` from step one
   `make run`
   then it will do a couple of thing here
   - `go mod tidy` for download and install dependencies  
   - `go run ./cmd/app/main.go` for running the program
  
	After that it will running with some testcase in `constant.go` file 
	the 
	````
	TCN =  map[string][]int{
	FirstLine: {n, k},
	SecondLine: {1, 2, 3, 4, 5, 6, 7, 8, 35, 120}, // position of each chicken
	}
	````
	and then add `constant.TCN` into `main.go` file in `testCases` variable
   I've assume that the length of `SecondLine` are equal to `n` as description in assessment sheet
   And you can run `make test` for running unit test file with verbose

---
## What can be improve
1. I still think that might be the way that I cannot figured out (yet) to not using nested for loop if I can do that, that would be decrease time complexity 