package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Palabra oculta que el jugador debe adivinar
	palabraSecreta := "golang"
	intentosMaximos := 6
	intentos := 0
	adivinadas := make([]bool, len(palabraSecreta))
	letrasAdivinadas := []string{}

	fmt.Println("Bienvenido al juego de ahorcado en Go!")
	fmt.Printf("La palabra tiene %d letras. Adivina la palabra.\n", len(palabraSecreta))

	for {
		fmt.Print("Palabra: ")
		mostrarPalabraSecreta(palabraSecreta, adivinadas)
		fmt.Printf("Letras adivinadas: %s\n", strings.Join(letrasAdivinadas, ", "))
		fmt.Printf("Intentos restantes: %d\n", intentosMaximos-intentos)

		if intentos == intentosMaximos {
			fmt.Println("¡Perdiste! La palabra era:", palabraSecreta)
			break
		}

		fmt.Print("Ingresa una letra: ")
		letra, _ := obtenerEntrada()
		letra = strings.ToLower(letra)

		if len(letra) != 1 || !esLetra(letra) {
			fmt.Println("Ingresa una sola letra válida.")
			continue
		}

		if yaAdivinada(letra, letrasAdivinadas) {
			fmt.Println("Ya adivinaste esa letra.")
			continue
		}

		if contieneLetra(palabraSecreta, letra) {
			fmt.Println("¡Correcto! La letra está en la palabra.")
			actualizarAdivinadas(palabraSecreta, letra, adivinadas)
			letrasAdivinadas = append(letrasAdivinadas, letra)
		} else {
			fmt.Println("¡Incorrecto! La letra no está en la palabra.")
			intentos++
		}

		if adivinadaCompletamente(adivinadas) {
			fmt.Println("¡Felicidades! Adivinaste la palabra:", palabraSecreta)
			break
		}
	}
}

func mostrarPalabraSecreta(palabraSecreta string, adivinadas []bool) {
	for i, letra := range palabraSecreta {
		if adivinadas[i] {
			fmt.Printf("%c ", letra)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func obtenerEntrada() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	entrada, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	entrada = strings.TrimSpace(entrada)
	return entrada, nil
}

func esLetra(s string) bool {
	return len(s) == 1 && s >= "a" && s <= "z"
}

func yaAdivinada(letra string, letrasAdivinadas []string) bool {
	for _, adivinada := range letrasAdivinadas {
		if letra == adivinada {
			return true
		}
	}
	return false
}

func contieneLetra(palabra string, letra string) bool {
	return strings.Contains(palabra, letra)
}

func actualizarAdivinadas(palabra string, letra string, adivinadas []bool) {
	for i, l := range palabra {
		if string(l) == letra {
			adivinadas[i] = true
		}
	}
}

func adivinadaCompletamente(adivinadas []bool) bool {
	for _, adivinada := range adivinadas {
		if !adivinada {
			return false
		}
	}
	return true
}
