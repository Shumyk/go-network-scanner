package nmap

import (
	log "log/slog"

	"github.com/Ullaakut/nmap/v3"
)

type Scanner struct {
	ip      string
	scanner *nmap.Scanner
}

// ex. 192.168.0.0/24
func (s *Scanner) Scan() *nmap.Run {
	log.Debug("Nmap starts scanning", log.String("ip", s.ip))

	results, warnings, err := s.scanner.Run()
	s.printWarnings(*warnings)
	s.checkError(err)
	return results
}

func (s *Scanner) printWarnings(warnings []string) {
	for _, warning := range warnings {
		log.Warn(
			"Warning while scanning",
			log.String("ip", s.ip),
			log.String("warning", warning),
		)
	}
}

func (s *Scanner) checkError(err error) {
	if err == nil {
		return
	}
	log.Error(
		"Error while scanning",
		log.String("ip", s.ip),
		log.Any("err", err),
	)
	panic(err)
}
