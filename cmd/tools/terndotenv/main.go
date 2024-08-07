package main
import (
	"fmt"
    "github.com/joho/godotenv"
    "os/exec"
	"os"
)

func main () {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Imprimir variáveis de ambiente para verificação
	fmt.Println("WS_DATABASE_USER:", os.Getenv("WS_DATABASE_USER"))
	fmt.Println("WS_DATABASE_PASSWORD:", os.Getenv("WS_DATABASE_PASSWORD"))

	// Verificar se o tern está no PATH
    path, err := exec.LookPath("tern")
	if err != nil {
        panic(err)
    }
	fmt.Printf("Tern is available at %s\n", path)

	cmd := exec.Command("tern", "migrate", "--migrations", "./internal/store/pgstore/migrations", "--config", "./internal/store/pgstore/migrations/tern.conf")

	// Capturar saída de erro para diagnóstico
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		panic(err)
	}

}