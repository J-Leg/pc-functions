package f

import (
	"context"
	"fmt"
	"github.com/J-Leg/tracula"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

var cfg *tracula.Config

func init() {
	if err := godotenv.Load(); err == nil {
		log.Println("Found .env file. Function must have been triggered locally.")
	}

	var ctx context.Context = context.Background()
	cfg = tracula.InitConfig(ctx, initDb(ctx))
	cfg.LocalEnabled = isLocal()

	log.Printf("Local: %t", cfg.LocalEnabled)
}

// ProcessDaily - Daily process receptor
func ProcessDaily(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	fmt.Println("~~~~~~~ Execute Daily Update ~~~~~~~")
	tracula.Execute(cfg)
	fmt.Println("\n\n~~~~~~~ Daily Update Complete ~~~~~~~")

	end := time.Now()

	executionElapsed := end.Sub(start)
	fmt.Printf("Total elapsed (Daily) execution time: %s\n\n", executionElapsed.String())
}

// ProcessMonthly - Monthly process receptor
func ProcessMonthly(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	fmt.Println("~~~~~~~ Execute Monthly Update ~~~~~~~")
	tracula.ExecuteMonthly(cfg)
	fmt.Println("\n\n~~~~~~~ Monthly Update Complete ~~~~~~~")

	end := time.Now()

	executionElapsed := end.Sub(start)
	fmt.Printf("Total elapsed (Monthly) execution time: %s\n\n", executionElapsed.String())
}

// Recover - Recovery process receptor
func Recover(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	fmt.Println("~~~~~~~ Execute Recovery ~~~~~~~")
	tracula.ExecuteRecovery(cfg)
	fmt.Println("~~~~~~~ Daily Update Complete ~~~~~~~")

	end := time.Now()

	executionElapsed := end.Sub(start)
	fmt.Printf("Total elapsed recovery execution time: %s\n\n", executionElapsed.String())
}
