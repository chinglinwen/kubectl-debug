package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Mount(rootfs, dst string) (out []byte, err error) {
	err = pathCheck(rootfs, dst)
	if err != nil {
		return nil, err
	}
	opts := fmt.Sprintf("rw,lowerdir=%v:%v", dst, rootfs) // we fake dst as lowerdir, to make cmd ok
	//fmt.Println( "mount", "-t", "overlay", "overlay", "-o", opts, dst)
	cmd := exec.Command("mount", "-t", "overlay", "overlay", "-o", opts, dst)
	return cmd.CombinedOutput()
}

func MountBind(src, dst string) (out []byte, err error) {
	err = pathCheck(src, dst)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("mount", "--bind", src, dst)
	return cmd.CombinedOutput()
}

// bind logs
func MountBindReadOnly(src, dst string) (out []byte, err error) {
	err = pathCheck(src, dst)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("mount", "--rbind", "-o", "ro", src, dst)
	return cmd.CombinedOutput()
}

func MountRBind(src, dst string) (out []byte, err error) {
	err = pathCheck(src, dst)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("mount", "-o", "rbind,ro", src, dst)
	return cmd.CombinedOutput()
}

func ReMount(src, dst string) (out []byte, err error) {
	err = pathCheck(src, dst)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("mount", "-o", "remount,ro", src, dst)
	return cmd.CombinedOutput()
}

// bind logs
func MountDev(dst string) (out []byte, err error) {
	// err = pathCheck(src, dst)
	// if err != nil {
	// 	return nil, err
	// }
	cmd := exec.Command("mount", "--bind", "/dev/", dst+"/dev")
	return cmd.CombinedOutput()
}

func MountProc(dst string) (out []byte, err error) {
	// err = pathCheck(src, dst)
	// if err != nil {
	// 	return nil, err
	// }
	cmd := exec.Command("mount", "--bind", "/proc", dst+"/proc")
	return cmd.CombinedOutput()
}

// bind logs
func UnMountProc(dst string) (out []byte, err error) {
	// err = pathCheck(src, dst)
	// if err != nil {
	// 	return nil, err
	// }
	cmd := exec.Command("umount", "-f", dst+"/proc")
	return cmd.CombinedOutput()
}

// bind logs
func UnMountDev(dst string) (out []byte, err error) {
	// err = pathCheck(src, dst)
	// if err != nil {
	// 	return nil, err
	// }
	cmd := exec.Command("umount", "-f", dst+"/dev")
	return cmd.CombinedOutput()
}

func pathCheck(src, dst string) error {
	_, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("src %v does not exist", src)
	}
	_ = os.MkdirAll(dst, 0755)
	// if err != nil {
	// 	fmt.Printf("create dst err: %v\n", err)
	// }
	_, err = os.Stat(dst)
	if err != nil {
		fmt.Printf("dst %v does not exist\n", dst)
		//return fmt.Errorf("dst %v does not exist", dst)
		return nil
	}
	return nil
}

func UnMount(dst string) (out []byte, err error) {
	cmd := exec.Command("umount", "-f", dst)
	out, err = cmd.CombinedOutput()
	if err != nil {
		cmd := exec.Command("umount", "-f", dst)
		return cmd.CombinedOutput()
	}
	err = os.RemoveAll(dst)
	return
}
