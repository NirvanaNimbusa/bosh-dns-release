package main

import (
	"flag"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/cloudfoundry/dns-release/src/dns/nameserverconfig/handler"
	"github.com/cloudfoundry/dns-release/src/dns/nameserverconfig/monitor"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var bindAddress string
	flag.StringVar(&bindAddress, "bindAddress", "", "address that our dns server is binding to")
	flag.Parse()

	if net.ParseIP(bindAddress) == nil {
		log.Fatalf("invalid ip: %s", bindAddress)
	}

	logger := boshlog.NewAsyncWriterLogger(boshlog.LevelDebug, os.Stdout, os.Stderr)
	defer logger.FlushTimeout(5 * time.Second)

	cmdRunner := boshsys.NewExecCmdRunner(logger)

	shutdown := make(chan struct{})
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)

	monitor := monitor.NewMonitor(
		handler.NewResolvConfHandler(bindAddress, boshsys.NewOsFileSystem(logger), cmdRunner),
		logger,
		3*time.Second,
	)
	go monitor.Run(shutdown)

	func() {
		<-sigterm
		close(shutdown)
	}()
}
