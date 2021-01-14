package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func downloadFileToPath(downloadURL, path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	src, err := httpGet(downloadURL)
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = io.Copy(out, src)
	return err
}

func httpGet(addr string) (io.ReadCloser, error) {
	parsed, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(parsed.String())
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func httpGetAndRead(addr string) ([]byte, error) {
	bodySrc, err := httpGet(addr)
	if err != nil {
		return nil, err
	}
	defer bodySrc.Close()

	return ioutil.ReadAll(bodySrc)
}

func httpGetAndParseJSON(addr string, target interface{}) error {
	body, err := httpGetAndRead(addr)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}

type javaVersion struct {
	major int
	minor int
	patch int
}

func (v javaVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

func getSystemJavaVersion() javaVersion {
	version := javaVersion{}
	// if java isn't installed, w.e, something else'll surely break anyway
	path, err := exec.LookPath("java")
	if err != nil {
		return version
	}

	cmd := exec.Command(path, "-version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return version
	}

	// this part split off for sake of unit testing
	return parseJavaVersion(out)
}

func parseJavaVersion(in []byte) javaVersion {
	major, minor, patch := 0, 0, 0
	// this is tailor written expecting openjdk -- I ain't about to license
	// Oracle stuff to find a version string format :\
	re := regexp.MustCompile(`(openjdk|java) version "(?P<major>\d+).(?P<minor>\d+).(?P<patch>\d+)(_\d+)?"`)
	if re.Match(in) {
		matches := re.FindStringSubmatch(string(in))
		if len(matches) != re.NumSubexp()+1 {
			return javaVersion{}
		}

		major, _ = strconv.Atoi(matches[re.SubexpIndex("major")])
		minor, _ = strconv.Atoi(matches[re.SubexpIndex("minor")])
		patch, _ = strconv.Atoi(matches[re.SubexpIndex("patch")])
	} else {
		// alternative output format I've seen in 1.12+ opendjdk alpine builds
		re = regexp.MustCompile(`(openjdk|java) version "(?P<minor>\d+)-ea"`)
		if !re.Match(in) {
			return javaVersion{}
		}

		matches := re.FindStringSubmatch(string(in))
		if len(matches) != re.NumSubexp()+1 {
			return javaVersion{}
		}

		major = 1
		minor, _ = strconv.Atoi(matches[re.SubexpIndex("minor")])
		patch = 0
	}

	return javaVersion{major, minor, patch}
}

func javaArgs(jarPath string, ro RuntimeOpts, v javaVersion) []string {
	args := make([]string, 0)

	args = append(args, fmt.Sprintf("-Xms%s", ro.InitialMemory))
	args = append(args, fmt.Sprintf("-Xmx%s", ro.MaxMemory))

	// aikar's flags
	// https://aikar.co/2018/07/02/tuning-the-jvm-g1gc-garbage-collector-flags-for-minecraft/
	args = append(args, "-XX:+UseG1GC")
	args = append(args, "-XX:+ParallelRefProcEnabled")
	args = append(args, "-XX:MaxGCPauseMillis=200")
	args = append(args, "-XX:+UnlockExperimentalVMOptions")
	args = append(args, "-XX:+DisableExplicitGC")
	args = append(args, "-XX:+AlwaysPreTouch")

	if ro.MaxMemory > 12*Gigabyte {
		args = append(args, "-XX:G1NewSizePercent=40")
		args = append(args, "-XX:G1MaxNewSizePercent=50")
		args = append(args, "-XX:G1HeapRegionSize=16M")
		args = append(args, "-XX:G1ReservePercent=15")
		args = append(args, "-XX:InitiatingHeapOccupancyPercent=20")
	} else {
		args = append(args, "-XX:G1NewSizePercent=30")
		args = append(args, "-XX:G1MaxNewSizePercent=40")
		args = append(args, "-XX:G1HeapRegionSize=8M")
		args = append(args, "-XX:G1ReservePercent=20")
		args = append(args, "-XX:InitiatingHeapOccupancyPercent=15")
	}

	args = append(args, "-XX:G1HeapWastePercent=5")
	args = append(args, "-XX:G1MixedGCCountTarget=4")
	args = append(args, "-XX:G1MixedGCLiveThresholdPercent=90")
	args = append(args, "-XX:G1RSetUpdatingPauseTimePercent=5")
	args = append(args, "-XX:SurvivorRatio=32")
	args = append(args, "-XX:+PerfDisableSharedMem")
	args = append(args, "-XX:MaxTenuringThreshold=1")
	args = append(args, "-Dusing.aikars.flags=https://mcflags.emc.gs")
	args = append(args, "-Daikars.new.flags=true")

	if ro.DebugGC {
		if v.major == 1 && v.minor >= 8 && v.minor <= 10 { // java 1.8-1.10
			args = append(args, "-Xloggc:gc.log")
			args = append(args, "-verbose:gc")
			args = append(args, "-XX:+PrintGCDetails")
			args = append(args, "-XX:+PrintGCDateStamps")
			args = append(args, "-XX:+PrintGCTimeStamps")
			args = append(args, "-XX:+UseGCLogFileRotation")
			args = append(args, "-XX:NumberOfGCLogFiles=5")
			args = append(args, "-XX:GCLogFileSize=1M")
		} else if v.major == 1 && v.minor >= 11 || v.major > 1 { // java 1.11+
			args = append(args, "-Xlog:gc*:logs/gc.log:time,uptime:filecount=5,filesize=1M")
		} else {
			// I ain't got no clue, check aikar's blog
		}
	}

	args = append(args, "-jar", jarPath)
	args = append(args, "nogui")

	return args
}

func runJavaServer(binaryPath string, runOpts RuntimeOpts) error {
	if _, err := os.Stat(binaryPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("server not installed, refusing to run")
		}

		return err
	}

	path, err := exec.LookPath("java")
	if err != nil {
		return err
	}
	version := getSystemJavaVersion()
	args := javaArgs(binaryPath, runOpts, version)
	cmd := exec.Command(path, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	fmt.Printf("Starting '%s %s'\n", path, strings.Join(args, " "))
	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		io.Copy(in, os.Stdin)
	}()

	cmdErrChan := make(chan error, 1)
	go func() {
		cmdErrChan <- cmd.Wait()
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	timeoutChan := make(chan int, 1)

	for {
		select {
		case err := <-cmdErrChan:
			return err
		case <-sigChan:
			if _, err := in.Write([]byte("stop\n")); err != nil {
				// ru-roh
			}

			go time.AfterFunc(6*time.Second, func() { timeoutChan <- 1 })
		case <-timeoutChan:
			if err := cmd.Process.Kill(); err != nil {
				// woah, what a hardy process
			}

			return nil
		}
	}
}
