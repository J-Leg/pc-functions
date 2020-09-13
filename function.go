package f

import (
  "context"
  "fmt"
  "github.com/j-leg/tracula"
  "github.com/j-leg/tracula/config"
  "github.com/joho/godotenv"
  "log"
  "net/http"
  "time"
)

var cfg *config.Config

func init() {
	if err := godotenv.Load(); err == nil {
		log.Println("Found .env file. Function must have been triggered locally.")
	}

	var ctx context.Context = context.Background()
	cfg = config.InitConfig(ctx, initDb(ctx))
	cfg.LocalEnabled = isLocal()

	log.Printf("Local: %t", cfg.LocalEnabled)
}

// ProcessDaily - Daily process receptor
func ProcessDaily(w http.ResponseWriter, r *http.Request) {
  start(cfg, "Daily", tracula.ExecuteDaily)
}

// ProcessMonthly - Monthly process receptor
func ProcessMonthly(w http.ResponseWriter, r *http.Request) {
  start(cfg, "Monthly", tracula.ExecuteMonthly)
}

// Recover - Recovery process receptor
func Recover(w http.ResponseWriter, r *http.Request) {
  start(cfg, "Recovery", tracula.ExecuteRecovery)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
  start(cfg, "Refresh", tracula.ExecuteRefresh)
}

func Track(w http.ResponseWriter, r *http.Request) {
  start(cfg, "Tracker", tracula.ExecuteTracker)
}

type exe func(cfg *config.Config)
func start(cfg *config.Config, jobType string, leg exe) {

	fmt.Printf("~~~~~~~ Execute %s Update ~~~~~~~\n", jobType)
	start := time.Now()
  leg(cfg)
  end := time.Now()
	fmt.Printf("~~~~~~~ %s Complete ~~~~~~~\n", jobType)

	executionElapsed := end.Sub(start)
	fmt.Printf("Total elapsed (Daily) execution time: %s\n\n", executionElapsed.String())
}
