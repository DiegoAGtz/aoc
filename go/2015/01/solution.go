package solutions_2015 

func partOne(input string) int {
    floor := 0

    for i := 0; i < len(input); i++ {
        if input[i] == '(' {
            floor++
        } else {
            floor -= 1
        }
    }

    return floor
}

func partTwo(input string) int {
    floor, position := 0, -1
    
    for i := 0; i < len(input) && floor != -1; i++ {
        if input[i] == '(' {
            floor++
        } else {
            floor -= 1
        }
        if position == -1 && floor == -1 {
            position = i + 1
        }
    }

    return position
}
