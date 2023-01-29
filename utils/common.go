package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Cfg struct {
	DBPort   int    `yaml:"dbPort,omitempty"`
	DBHost   string `yaml:"dbHost,omitempty"`
	DBName   string `yaml:"dbName,omitempty"`
	Password string `yaml:"password,omitempty"`
	Username string `yaml:"username,omitempty"`
}

type cfgParser struct{}

func ParserParseEnvCfg(path string) (Cfg, error) {

	b, err := ioutil.ReadFile(path) // b has type []byte
	if err != nil {
		log.Error(err)

	}

	result := Cfg{}

	err = yaml.Unmarshal(b, &result)
	if err != nil {
		log.Error(err)
		return result, err
	}

	return result, nil
}

func MD5(str string) string {
	b := []byte(str)
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5_SALT(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// LocalIP get the host machine local IP address
func LocalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if isPrivateIP(ip) {
				return ip, nil
			}
		}
	}

	return nil, errors.New("no IP")
}

func isPrivateIP(ip net.IP) bool {
	var privateIPBlocks []*net.IPNet
	for _, cidr := range []string{
		// don't check loopback ips
		//"127.0.0.0/8",    // IPv4 loopback
		//"::1/128",        // IPv6 loopback
		//"fe80::/10",      // IPv6 link-local
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, block)
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}

	return false
}
