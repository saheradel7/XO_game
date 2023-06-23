package main

import "fmt"

func checkPlayerStatus(arr [3][3]string) bool {

	if arr[0][0] == arr[0][1] && arr[0][0] == arr[0][2] && arr[0][1] == arr[0][2] {
		return false
	} else if arr[1][0] == arr[1][1] && arr[1][0] == arr[1][2] && arr[1][1] == arr[1][2] {
		return false
	} else if arr[2][0] == arr[2][1] && arr[2][0] == arr[2][2] && arr[2][1] == arr[2][2] {
		return false
	} else if arr[0][0] == arr[1][0] && arr[0][0] == arr[2][0] && arr[1][0] == arr[2][0] {
		return false
	} else if arr[0][1] == arr[1][1] && arr[0][1] == arr[2][1] && arr[1][1] == arr[2][1] {
		return false
	} else if arr[0][2] == arr[1][2] && arr[0][2] == arr[2][2] && arr[1][2] == arr[2][2] {
		return false
	} else if arr[0][0] == arr[1][1] && arr[0][0] == arr[2][2] && arr[1][1] == arr[2][2] {
		return false
	} else if arr[0][2] == arr[1][1] && arr[0][2] == arr[2][0] && arr[1][1] == arr[2][0] {
		return false
	}
	return true
}

func playerOne(position string, XO [3][3]string) [3][3]string{

	switch position{
	case "1":
		XO[0][0] = "  x  "
	case "2":
		XO[0][1] = "  x  "
	case "3":
		XO[0][2] = "  x  "
	case "4":
		XO[1][0] = "x"
	case "5":
		XO[1][1] = "  x  "
	case "6":
		XO[1][2] = "  x  "
	case "7":
		XO[2][0] = "  x  "
	case "8":
		XO[2][1] = "  x  "
	case "9":
		XO[2][2] = "  x  "	
	default:
		fmt.Println("invalid position, please try again")

	}
	return XO
}

func playerTwo(position string, XO [3][3]string)[3][3]string{

	switch position{
	case "1":
		XO[0][0] = "  O  "
	case "2":
		XO[0][1] = "  O  "
	case "3":
		XO[0][2] = "  O  "
	case "4":
		XO[1][0] = "  O  "
	case "5":
		XO[1][1] = "  O  "
	case "6":
		XO[1][2] = "  O  "
	case "7":
		XO[2][0] = "  O  "
	case "8":
		XO[2][1] = "  O  "
	case "9":
		XO[2][2] = "  O  "	
	default:
		fmt.Println("invalid position, please try again")

	}
	return XO
}


func print(arr [3][3]string) {
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
}

func main() {
	var XO = [3][3]string{
		{"  1  ", "  2  ", "  3  "},
		{"  4  ", "  5  ", "  6  "},
		{"  7  ", "  8  ", "  9  "},
	}

	isFinished := true
	for isFinished{
		print(XO)
		fmt.Println("player 1 enter position")
		var p string
		fmt.Scan(&p)
		XO = playerOne(p,XO)
		print(XO)
		isFinished = checkPlayerStatus(XO)
		if !isFinished{
			fmt.Print("player one win the game")
			return
		}else{
			fmt.Println("player 2 enter position")
			fmt.Scan(&p)
			XO = playerTwo(p,XO)
			isFinished = checkPlayerStatus(XO)
			if !isFinished{
				fmt.Print("player tow win the game")
				return
			}
		}
	}
}
