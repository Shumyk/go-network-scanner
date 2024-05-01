package pool

import (
	"go-network-scanner/api"
	"go-network-scanner/nmap"
)

type Worker struct {
	ipRange     string
	enableSocks bool
	nmap        *nmap.Scanner
}

func (w *Worker) Do() api.Scan { // TODO: implement
	w.nmap.Scan(w.ipRange)
	return api.Scan{}
}
