package main

import (
    "bufio"
    "os"
    "strconv"
    "fmt"
    "log"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    games, err  := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Panic("err in games scanner")
    } 
        

    for i := 0; i <= games - 1; i++ {
        
        min := 0
        max := 1000
        num := 500
        found := false

        for found == false {
            fmt.Println(num)
            scanner.Scan()
            response := scanner.Text()

            if response == "lower" {
                max = num
                num = num - ((max-min)/2)


            } else if response == "higher" {
                min = num
                num = num + ((max-min)/2)

            } else {
                found = true
            }
        }
    }

}
