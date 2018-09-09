package module

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"

	"os/exec"

	"github.com/doovemax/ban_ip/conf"
)

func IpAddFilter() (err error) {
	// logrus.Println(IPcount)
	iPcountRWMutex.RLock()
	defer iPcountRWMutex.RUnlock()
	for urliptime, count := range IPcount {
		if count > conf.DefaultConfig.IPToBlackCount {
			// fmt.Println(urliptime, count)
			urlSplit := strings.Split(urliptime, ":")
			err := IPaddFirewoall(urlSplit[1])
			if err != nil {
				if n:= strings.Index(err.Error(),"IP in firewall");n != -1{
					delete(IPcount,urliptime)
				}else{
					logrus.Warn(err)
				}
				// logrus.Warn(urlSplit," add black list fail")
				continue
			}else {
				logrus.Info(urlSplit[1], " add black list success")
			}

		}
	}
	return
}

func IPaddFirewoall(IP string) (err error) {
	err = firewallCheck(IP)
	if err != nil {
		return
	}
	// logrus.Fatal(err)
	cmdString := fmt.Sprintf("--add-entry=%s", IP)
	cmd := exec.Command("/usr/bin/firewall-cmd","--permanent","--zone=public","--ipset=blacklist",cmdString)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {

		return err
	}

	outerr, _ := ioutil.ReadAll(stderr)

	if len(outerr) != 0 {
		s := string(outerr)
		return fmt.Errorf(s)
	}

	err = firewallReload()
	if err != nil {
		return err
	}

	return nil
}
func firewallReload() (err error) {
	cmd := exec.Command("/usr/bin/firewall-cmd","--reload")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {

		return err
	}

	outerr, _ := ioutil.ReadAll(stderr)

	if len(outerr) != 0 {
		s := string(outerr)
		return fmt.Errorf(s)
	}

	return nil
}

func firewallCheck(IP string)(err error){
	cmdString := fmt.Sprintf("--query-entry=%s", IP)
	cmd := exec.Command("/usr/bin/firewall-cmd","--permanent","--zone=public","--ipset=blacklist",cmdString)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	output, _ := ioutil.ReadAll(stdout)

	if len(output) != 0 {
		s := string(output)
		switch s {
		case "yes\n":
		return 	fmt.Errorf("IP in firewall",cmd.Args)
		case "no\n":
		return nil

		}
	}
	return nil
}
