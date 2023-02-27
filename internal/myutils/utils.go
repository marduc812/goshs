package myutils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"mime"
	"net"
	"strings"

	"github.com/patrickhener/goshs/internal/mylog"
)

// ByteCountDecimal generates human readable file sizes and returns a string
func ByteCountDecimal(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

// MimeByExtension returns the mimetype string depending on the filename and its extension
func MimeByExtension(n string) string {
	mylog.Debugf("The string handed to MimeByExtension is: %s\n", n)
	mylog.Debugf("Discovered Extension: %s\n", mime.TypeByExtension(ReturnExt(n)))
	return mime.TypeByExtension(ReturnExt(n))
}

// ReturnExt returns the extension without from a filename
func ReturnExt(n string) string {
	extSlice := strings.Split(n, ".")
	mylog.Debugf("The sliced extension is: %s\n", extSlice)
	return "." + extSlice[len(extSlice)-1]
}

// RandomNumber returns a random int64
func RandomNumber() (big.Int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		mylog.Errorf("when generating random number: %+v", err)
		return *big.NewInt(0), err
	}
	return *n, err
}

// CheckSpecialPath will check a slice of special paths against
// a folder on disk and return true if it matches
func CheckSpecialPath(check string) bool {
	specialPaths := []string{"425bda8487e36deccb30dd24be590b8744e3a28a8bb5a57d9b3fcd24ae09ad3c", "cf985bddf28fed5d5c53b069d6a6ebe601088ca6e20ec5a5a8438f8e1ffd9390", "14644be038ea0118a1aadfacca2a7d1517d7b209c4b9674ee893b1944d1c2d54"}

	for _, item := range specialPaths {
		if item == check {
			return true
		}
	}

	return false
}

// GetInterfaceIpv4Addr will return the ip address by name
func GetInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
	)
	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		return
	}
	if addrs, err = ief.Addrs(); err != nil { // get addresses
		return
	}
	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return "", fmt.Errorf("interface %s doesn't have an ipv4 address", interfaceName)
	}
	return ipv4Addr.String(), nil
}

// GetAllIPAdresses will return a map of interface and associated ipv4 addresses for displaying reasons
func GetAllIPAdresses() (map[string]string, error) {
	ifaceAddress := make(map[string]string)

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range ifaces {
		ip, err := GetInterfaceIpv4Addr(i.Name)
		if err != nil {
			continue
		}

		ifaceAddress[i.Name] = ip

	}
	return ifaceAddress, nil

}
