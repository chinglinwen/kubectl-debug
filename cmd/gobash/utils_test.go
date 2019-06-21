package main

import (
	"fmt"
	"testing"
)

func TestParseDF(t *testing.T) {
	a := `
	Filesystem           1K-blocks      Used Available Use% Mounted on
overlay               51474912  18410672  30426416  38% /
tmpfs                    65536         0     65536   0% /dev
tmpfs                 32804980         0  32804980   0% /sys/fs/cgroup
overlay               51474912  18410672  30426416  38% /container
/dev/sda2             51474912  18410672  30426416  38% /container/logs
/dev/sda2             51474912  18410672  30426416  38% /etc/hostname
shm                      65536         0     65536   0% /dev/shm
/dev/sda3            1676157212  23051112 1567938844   1% /container/dev/termination-log
/dev/sda3            1676157212  23051112 1567938844   1% /container/etc/hosts
172.31.83.26:/data/staticfile_yjr/file_data/openapi
                     5997913088 2783210496 3214702592  46% /container/apps/tangguo/storage/app/upload
tmpfs                 32804980        12  32804968   0% /run/secrets/kubernetes.io/serviceaccount
tmpfs                 32804980         0  32804980   0% /proc/acpi
tmpfs                    65536         0     65536   0% /proc/kcore
tmpfs                    65536         0     65536   0% /proc/keys
tmpfs                    65536         0     65536   0% /proc/timer_list
tmpfs                    65536         0     65536   0% /proc/sched_debug
tmpfs                 32804980         0  32804980   0% /proc/scsi
tmpfs                 32804980         0  32804980   0% /sys/firmware
overlay               51474912  18410672  30426416  38% /newroot
tmpfs                    65536         0     65536   0% /newroot/dev
overlay               51474912  18410672  30426416  38% /newroot
tmpfs                    65536         0     65536   0% /newroot/dev`
	s := parsedf(a)
	fmt.Println(s)
}
