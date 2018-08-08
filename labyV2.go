package main

import "fmt"
import "math/rand"
import "time"
import "strconv"
// import "os"

var (
	size int 
	line string = ""
	placeX int 
	placeY int
	ended bool = false
	goThere int
	verifline string = ""
	tryExit int 
	verifX int 
	verifY int
)

func isOk(forme int, side int) bool {
	switch forme {
		case 0 : 
		return true

		case 1 : 
		if (side == 2 || side == 4) {
			return true
		}
		return false

		case 2 : 
		if (side == 1 || side == 3) {
			return true
		}
		return false

		case 3 : 
		if (side == 4 || side == 3) {
			return true
		}
		return false

		case 4 : 
		if (side == 3 || side == 2) {
			return true
		}
		return false

		case 5 : 
		if (side == 1 || side == 2) {
			return true
		}
		return false

		case 6 : 
		if (side == 4 || side == 1) {
			return true
		}
		return false

		case 7 : 
		if (side == 4 || side == 3 || side == 2) {
			return true
		}
		return false

		case 8 : 
		if (side == 1 || side == 2 || side == 3) {
			return true
		}
		return false

		case 9 : 
		if (side == 1 || side == 2 || side == 4) {
			return true
		}
		return false

		case 10 : 
		if (side == 1 || side == 4 || side == 3) {
			return true
		}
		return false

		case 11 : 
		if (side == 1) {
			return true
		}
		return false

		case 12 : 
		if (side == 2) {
			return true
		}
		return false

		case 13 : 
		if (side == 3) {
			return true
		}
		return false

		case 14 : 
		if (side == 4) {
			return true
		}
		return false


		default : 
		return false
	}
}

func printLaby(board [][]int){

	for i := 0; i < size; i++ {
			if(isOk(board[0][i], 1)){
				line += " _";
			} else{
				line += "  ";
			}
		}
		fmt.Println(line);
		line = "";

		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {

				if(i>0){
					if(isOk(board[j][i-1], 2) && isOk(board[j][i], 4)){
						line += "|";
					} else {
						line += " ";
					}
				} else if(isOk(board[j][i], 4)){
					line += "|";
					} else{
						line += " ";
					}


				if(j<size-1){
					if(isOk(board[j][i], 3) && isOk(board[j+1][i], 1)){
						line += "_";
					} else {
						line += " ";
					}
				} else if(isOk(board[j][i], 3)){
					line += "_";
					} else{
						line += " ";
					}
			}
			if(isOk(board[j][size-1], 2)){
				line += "|";
				fmt.Println(line);
				line = "";
			} else {
				line += " ";
				fmt.Println(line);
				line = "";
			}
		}

}

func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func changeForm(actualForm int, side int) int {
	switch actualForm {
		case 0 : 
		switch side {
			case 1 : 
			return 7
			case 2 : 
			return 10
			case 3 : 
			return 9
			case 4 : 
			return 8
		}

		case 10 : 
		switch side {
			case 1 : 
			return 3
			case 3 : 
			return 6
			case 4 : 
			return 1
		}

		case 9 : 
		switch side {
			case 1 : 
			return 1
			case 2 : 
			return 6
			case 4 : 
			return 5
		}

		case 8 : 
		switch side {
			case 1 : 
			return 4
			case 2 : 
			return 2
			case 3 : 
			return 5
		}

		case 7 : 
		switch side {
			case 2 : 
			return 3
			case 3 :
			return 1
			case 4 : 
			return 4
		}

		case 6 : 
		switch side {
		case 1 :
		return 14
		case 4 : 
		return 11 
		}

		case 5 : 
		switch side {
		case 1 :
		return 12
		case 2 :
		return 11 
		}

		case 4 : 
		switch side {
		case 2 : 
		return 13
		case 3 : 
		return 12 
		}

		case 3 : 
		switch side {
		case 3 :
		return 14
		case 4 : 
		return 13
		}

		case 2 : 
		switch side {
			case 1 : 
			return 13
			case 3 : 
			return 11
		}

		case 1 : 
		switch side {
			case 2 : 
			return 14
			case 4 : 
			return 12
		}
	}

	return 1 
}

func findExit(verifBoard [][]int, board [][]int){

	placeY = (size/2);
	placeX = 0;
	verifBoard[placeY][placeX] = 1;
	board[placeY][placeX] = 8;


	for(!ended){

	tryExit++;

	if(tryExit > 50){
		placeX++;
	}


	// if(tryExit > 50){
	// 	fmt.Println("reboot");
	// 	tryExit = 0;

	// 	board := make([][]int, size)
	// 	for i := range board {
	// 	    board[i] = make([]int, size)
	// 	}

	// 	verifBoard := make([][]int, size)
	// 	for i := range verifBoard {
	// 	    verifBoard[i] = make([]int, size)
	// 	}



	// }

	rand.Seed(time.Now().UnixNano())


		goThere = randomInt(1,5);
		switch goThere {
			case 1 :
				//fmt.Println("1");
			if(placeY != 0){
				if(verifBoard[placeY-1][placeX] != 1){
					tryExit = 0;
					verifBoard[placeY-1][placeX] = 1; // ajout check d'utilisation
					board[placeY][placeX] = changeForm(board[placeY][placeX], 1);
					placeY -= 1; // deplacement dans le laby 
					tryExit = 0;
					//changement de la forme
				}
			} else{
				//board[placeY][placeX] = changeForm(board[placeY][placeX], 1);
				ended = false
			}

			case 2 : 
			//fmt.Println("2");
			if(placeX != size-1){
				if(verifBoard[placeY][placeX + 1] != 1){
					tryExit = 0;
					verifBoard[placeY][placeX + 1] = 1;
					board[placeY][placeX] = changeForm(board[placeY][placeX], 2);
					placeX += 1;
					tryExit = 0;
				}
			} else {
				board[placeY][placeX] = changeForm(board[placeY][placeX], 2);
				ended = true
			}

			case 3 : 
			//fmt.Println("3");
			if(placeY != size-1){
				if(verifBoard[placeY + 1][placeX ] != 1){
					tryExit = 0;
					verifBoard[placeY + 1][placeX ] = 1;
					board[placeY][placeX] = changeForm(board[placeY][placeX], 3);
					placeY += 1;
					tryExit = 0;
				}
			} else {
				//board[placeY][placeX] = changeForm(board[placeY][placeX], 3);
				ended = false
			}

			case 4 : 
			//fmt.Println("4");
			if(placeX != 0){
				if(verifBoard[placeY][placeX  - 1] != 1){
					tryExit = 0;
					verifBoard[placeY][placeX - 1] = 1;
					board[placeY][placeX] = changeForm(board[placeY][placeX], 4);
					placeX -= 1;
					tryExit = 0;
				}
			} else {
				//board[placeY][placeX] = changeForm(board[placeY][placeX], 4);
				ended = false
			}

			default : 
			fmt.Println("default");
			fmt.Println("case defaut");

		}


	}

}

func verifSolution(verifBoard [][]int ){

			for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {

				verifline += strconv.Itoa(verifBoard[j][i]);

				}

				fmt.Println(verifline);
				verifline = "";
			}
}


func main() {

//	reader := bufio.NewReader(os.Stdin)
	fmt.Println("entrée la taille du labyrinthe : ")
	fmt.Scanf("%d", &size)



	board := make([][]int, size)
	for i := range board {
	    board[i] = make([]int, size)
	}


	verifBoard := make([][]int, size)
	for i := range verifBoard {
	    verifBoard[i] = make([]int, size)
	}

	////////// find exit function ///////////////

	findExit(verifBoard, board);

	verifSolution(verifBoard);


	for j := 0; j < size-1; j++ {

			for i := 1; i < size; i++ {
					
						if(verifBoard[j][i] != 1){

							board[j][i] = randomInt(8,10)
						}
				
			}
	}



	



	/////////////////////////////////////////////

	printLaby(board);


	

	

}