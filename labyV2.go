package main

import "fmt"
import "math/rand"
import "time"
import "strconv"
// import "os"

var (

	//taille du labyrinthe
	size int 
	//chaine de caractere permettant d'afficher le labyrinthe
	line string = ""

	placeX int 

	placeY int

	ended bool = false

	goThere int

	verifline string = ""

	tryExit int 

	verifX int 

	verifY int

	boolSolution int = 0
)

//fonction verifiant si un mur du labyrinthe doit etre afficher ou cacher
//elle prend en argument la forme de la case est le coté de la case

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

//fonction qui s'occupe d'afficher le labyrinthe 

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

//fonction renvoyant un chiffre aléatoire compris entre min et max en argument 

func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

//fonction permettant de changer la forme d'une case d'un labyrinthe 
//elle renvoie la nouvelle forme que doit avoir la case

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

// fonction trancant le chemin de l'entree du labyrinthe jusqu'a sa sortie
func findExit(verifBoard [][]int, board [][]int){

	placeY = (size/2);
	placeX = 0;
	verifBoard[placeY][placeX] = 1;
	board[placeY][placeX] = 8;


	for(!ended){

	tryExit++;

	if(tryExit > 500){
		tryExit = 0;
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
			if(placeY != 0){ // si nous ne somme pas sur un bord du labyrinthe
				if(verifBoard[placeY-1][placeX] != 1){ //si nous ne sommes pas déja passé sur cette case
					verifBoard[placeY-1][placeX] = 1; // ajout check d'utilisation
					board[placeY][placeX] = changeForm(board[placeY][placeX], 1); // changement de la forme de la case
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
					verifBoard[placeY][placeX + 1] = 1;
					board[placeY][placeX] = changeForm(board[placeY][placeX], 2);
					placeX += 1;
					tryExit = 0;
				}
			} else { // si nous sommes sur le cote droit du labyrinthe
				board[placeY][placeX] = changeForm(board[placeY][placeX], 2); // créer ouverture de sortie
				ended = true // fin de la fonction et du tracage du chemin de sortie
			}

			case 3 : 
			//fmt.Println("3");
			if(placeY != size-1){
				if(verifBoard[placeY + 1][placeX ] != 1){
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

//fonction permettant d'afficher les case déja traité par la fonction findExit 
//elle permet de voir le chemin que la fonction find exit a tracé

func verifSolution(verifBoard [][]int ){

		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {

				verifline += strconv.Itoa(verifBoard[j][i]);

				}

				fmt.Println(verifline);
				verifline = "";
			}
}

//fonction permettant de modifié l'ensemble des case non traité par la fonction finEcit afin de "brouillé" 
//le chemin et rendre plus difficile sa sortie 

func shakeLaby(verifBoard [][]int, board [][]int, size int){

	for j := 0; j < size-1; j++ {

			for i := 1; i < size; i++ {
					
						if(verifBoard[j][i] != 1){

							board[j][i] = randomInt(8,10)
						}
				
			}
	}

}

func main() {

//	reader := bufio.NewReader(os.Stdin)
	fmt.Println("entrée la taille du labyrinthe : ")
	fmt.Scanf("%d", &size)

//initialisation du tableau 2 dimensions de la taille du labyrinthe permettant de tracer le labyrinthe 

	board := make([][]int, size)
	for i := range board {
	    board[i] = make([]int, size)
	}

	boardSolution := make([][]int, size)
	for i := range boardSolution {
	    boardSolution[i] = make([]int, size)
	}

//initialisation du tableau 2 dimensions permettant de savoir si une case a déja était traité lors du tracage du chemin de sortie
	verifBoard := make([][]int, size)
	for i := range verifBoard {
	    verifBoard[i] = make([]int, size)
	}

//tracage de la sortie du labyrinthe 

	findExit(verifBoard, board);


//duplication du chemin dans un autre tableau afin de garder la solution de sortie

		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {

				boardSolution[j][i] = board[j][i];

				}

			}
	

//affichage du chemin déssiné par la fonction find exit
//	verifSolution(verifBoard);

//modification de l'ensemble des cases non traité par la fonction findExit afin de "brouillé" le chemin et de rendre difficile sa sortie
//mettre en commentaire cette fonction permettra d'afficher uniquement le chemin de sortie créer par la fonction findExit

	shakeLaby(verifBoard , board , size );



	
//affichage du labyrinthe 
	printLaby(board);

	for (boolSolution != 1){

		fmt.Println("Entre 1 si vous voulez la solution  : ")
		fmt.Scanf("%d", &boolSolution)

	}

	printLaby(boardSolution);

	


	

	

}