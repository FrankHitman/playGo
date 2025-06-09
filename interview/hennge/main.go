package main


import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    N, _ := strconv.Atoi(strings.TrimSpace(input))

    // store the result of each test case
    results := make([]int, N)

    // use inner function as a replacement of for statement
    var processTestCase func(int)
    processTestCase = func(i int) {
        if i >= N {
            return
        }
        input, _ := reader.ReadString('\n')
        X, _ := strconv.Atoi(strings.TrimSpace(input))

        input, _ = reader.ReadString('\n')
        numbers := strings.Fields(input)
        if X != len(numbers){
            processTestCase(i + 1)
            return
        }

        sumOfSquares := 0
        var processNumbers func(int)
        processNumbers = func(j int) {
            if j >= len(numbers) {
                results[i] = sumOfSquares
                processTestCase(i + 1)
                return
            }
            num, _ := strconv.Atoi(numbers[j])
            if num >= 0 {
                sumOfSquares += num * num
            }
            processNumbers(j + 1)
        }
        processNumbers(0)
    }
    processTestCase(0)

    var printResults func(int)
    printResults = func(i int) {
        if i >= len(results) {
            return
        }
        fmt.Println(results[i])
        printResults(i + 1)
    }
    printResults(0)
}



