package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	// "github.com/docker/docker/pkg/mount"
)

var root = "/container"

func main() {
	var err error

	mounts, err := getdf()
	if err != nil {
		log.Fatal("get mounts info err")
	}

	// infos, err := mount.GetMounts()
	// if err != nil {
	// 	log.Fatal("get mounts info err")
	// }
	// for _, v := range infos {
	// 	if strings.HasPrefix("/container", v.Mountpoint) {
	// 		fmt.Printf("%#v\n", v)
	// 	}
	// }

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Println("get workingdir err", err)
	// }
	// log.Println("workingdir: ", dir)

	// mount container bind as read only
	dst := "/newroot"
	// dst := "/container"
	_, err = MountBindReadOnly("/container", dst)
	// err = mount.Mount("/container", dst, "ext4", "rbind,ro,relatime")
	if err != nil {
		log.Fatal("mount err", err)
	}

	_, err = MountDev(dst)
	if err != nil {
		log.Fatal("mount dev err", err)
	}

	_, err = MountProc(dst)
	if err != nil {
		log.Fatal("mount proc err", err)
	}

	for _, v := range mounts {
		t := strings.Replace(v, root, dst, -1)
		// fmt.Println("start mount ", v, t)
		_, err = MountBindReadOnly(v, t)
		if err != nil {
			log.Fatalf("mount %v to %v err: %v\n", v, t, err)
		}
	}
	// }

	// _, err = ReMount("/newroot", dst)
	// if err != nil {
	// 	log.Fatal("mount err", err)
	// }

	defer func() {
		UnMountDev(dst)
		UnMountProc(dst)
		UnMount(dst)

		for _, v := range mounts {
			t := strings.Replace(v, root, dst, -1)
			UnMount(t)
		}
	}()
	//execute shell

	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// cmd.Dir = c.WorkDir

	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		// Unshareflags: syscall.CLONE_NEWNS,
		Credential: &syscall.Credential{Uid: 7373, Gid: 7373},
	}

	//chroot
	err = syscall.Chroot(dst)
	if err != nil {
		err = fmt.Errorf("chroot %v error, err: %v", dst, err)
		log.Fatal(err)
	}
	err = os.Chdir("/")
	if err != nil {
		err = fmt.Errorf("change dir to / error, err: %v", err)
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal("exec err", err)
	}
	log.Fatal(cmd.Wait())
}

func getdf() (mounts []string, err error) {
	cmd := exec.Command("sh", "-c", "df")
	// cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return
	}
	mounts = parsedf(string(out))
	return
}

func parsedf(df string) (a []string) {
	s := strings.Split(df, "\n")
	for _, v := range s {
		// this will mount fail
		if strings.Contains(v, "termination-log") {
			continue
		}
		if strings.Contains(v, root+"/") {
			s1 := strings.Fields(v)
			a = append(a, s1[len(s1)-1])
		}
	}
	return
}
