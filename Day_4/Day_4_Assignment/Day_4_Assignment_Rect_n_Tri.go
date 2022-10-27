package main

import "fmt"

func main() {
	
	//input variable declared
	fmt.Print("Masukkan angka integer: ")
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("Input adalah: %d\n", n)

	//looping variables declared
	var i int = 0
	var j int = 0
	var k int = 0
	
	//if n divided by 2 has residue of 0 (even) then create rectangle n row and n column
	if n % 2 == 0{
		//create a line of =
		for k=1; k<=23; k++{
			fmt.Print("=")
		}
		fmt.Println("\nIni yg Genap :")

		//create n number of row
		for i=1; i<=n; i++{
			//create n number of column
			for j=1;j<=n;j++{
				fmt.Print("*")
			}
			fmt.Print("\n")
		}
	}

	//if n divided by 2 has residue of 1 (odd) then create triangle of n row and i column
	if n % 2 == 1{
		//create a line of =
		for k=1; k<=23; k++{
			fmt.Print("=")
		}
		fmt.Println("\nIni yg Ganjil")

		// create n number of row
		for i:=1; i<=n; i++{
			//create i number of column
			for j:=1;j<=i;j++{
				fmt.Print("*")
			}
			fmt.Print("\n")
		}
	}
}

