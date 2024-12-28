package main


import (
	"fmt"
	"os"
	"encoding/json"
	"log"
)
type zadanie struct {
	ID		 int
	Nazwa	 string
	Wykonane bool
}
var zadania []zadanie
var lista_id = []int{0}
func DodajZadanie(Nazwa string) {
	var zad zadanie
	for _, v := range lista_id {
		if v+1 != zad.ID {
			zad.ID = v+1
			lista_id = append(lista_id, zad.ID)
		}
	}
	zad.Nazwa = Nazwa
	zad.Wykonane = false
	zadania = append(zadania, zad)
}
func Wyswietl() {
	for i:=0; i < len(zadania) ; i++{
		if zadania[i].Wykonane == true {
			fmt.Print(zadania[i].ID," ☑ ",zadania[i].Nazwa,"\n")
		} else {
			fmt.Print(zadania[i].ID," ☐ ",zadania[i].Nazwa,"\n")
		}
	}

}
func oznacz(ID int) {
	zadania[ID-1].Wykonane = true
}
func UsunZadanie(IDdoUsun int) {
	IDdoUsun--
	zadania = append(zadania[:IDdoUsun], zadania[IDdoUsun+1:]...)
	lista_id = append(lista_id[:IDdoUsun], lista_id[IDdoUsun+1:]...)
}
func odczytPlik(sciezka string,tablica []zadanie,tablica2 []int) {
	plik, err := os.ReadFile(sciezka)
	if err != nil {
		log.Fatal(err)
	}
	var dane []zadanie
	err = json.Unmarshal(plik, &dane)
	if err != nil {
		log.Fatal(err)
	}
	tablica = []zadanie{}
	tablica2 = []int{0}
	for _,z := range dane {
		tablica = append(tablica,z)
		tablica2 = append(tablica2,z.ID)
	}
}
func zapisPlik() {

}
func main() {
	var opcja int
	var UtworzID string
	var UsunID int
	var WykonajID int
	var Sciezka string
	for {
		fmt.Print("-----------------------------\n")
		fmt.Print("Wybierz opcję\n1.Wyświetl zadania.\n2.Dodaj zadanie.\n3.Usuń zadanie.\n4.Oznacz jako wykonane.\n5.Wczytaj Zadania z Pliku.\n6.Zapisz zadania do Pliku\n7.Opuść program\nNapisz numer opcji:  ")
		fmt.Scan(&opcja)
		switch opcja {
		case 1:
			fmt.Print("-----------------------------\n")
			if len(zadania) == 0{
				fmt.Print("Lista zadań jest pusta, dodaj zadania.\n")
			}else{
				Wyswietl()
			}
		case 2:
			fmt.Print("-----------------------------\n")
			fmt.Print("Podaj nazwę zadania: ")
			fmt.Scan(&UtworzID)
			DodajZadanie(UtworzID)
		case 3:
			fmt.Print("-----------------------------\n")
			fmt.Print("Napisz numer zadania które chcesz usunąć: ")
			fmt.Scan(&UsunID)
			UsunZadanie(UsunID)
		case 4:
			fmt.Print("-----------------------------\n")
			fmt.Print("Napisz numer zadania które wykonałeś: ")
			fmt.Scan(&WykonajID)
			oznacz(WykonajID)
		case 5:
			fmt.Print("-----------------------------\n")
			fmt.Print("Podaj ścieżke do pliku: ")
			fmt.Scan(&Sciezka)
			odczytPlik(Sciezka,zadania,lista_id)
		case 6:

		case 7:
			os.Exit(0)
		default:
			fmt.Print("-----------------------------\n")
			fmt.Print("Taka opcja nie istnieje\n")
		}
	}
}