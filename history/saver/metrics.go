package saver

import (
	"time"

	"golang.org/x/sys/unix"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	queueLength = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "saver_trades_queue_length",
		Help: "The current number of trades in queue waiting to be writed",
	})

	diskAll = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "disk_all_size",
		Help: "Total bytes of disk size.",
	})
	diskAvail = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "disk_avail_size",
		Help: "The available bytes of disk space",
	})
	diskFree = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "disk_free_size",
		Help: "The free bytes of disk space",
	})
	diskUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "disk_used_size",
		Help: "The used bytes of disk space",
	})
)

func collectDiskMetrics(path string, frequency time.Duration) {
	fs := unix.Statfs_t{}
	err := unix.Statfs(path, &fs)
	if err != nil {
		log.WithError(err).Error("cant start disk metrics collecting")
		return
	}
	ticker := time.NewTicker(frequency)
	//s.Add(1)  TODO add label and start/stop node functionality
	go func() {
		//defer s.Done()
		for {
			diskAll.Set(float64(fs.Blocks * uint64(fs.Bsize)))
			diskAvail.Set(float64(fs.Bavail * uint64(fs.Bsize)))
			diskFree.Set(float64(fs.Bfree * uint64(fs.Bsize)))
			diskUsed.Set(float64((fs.Blocks - fs.Bfree) * uint64(fs.Bsize)))
			<-ticker.C
		}
	}()
}
