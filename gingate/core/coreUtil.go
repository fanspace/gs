package core

import (
	"bytes"
	"errors"
	"fmt"
	"gingate/commons"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

var (
	commonBaseSearchPaths = []string{
		".",
		"..",
		"../..",
		"../../..",
	}
	LOG_FILENAME = "critic.log"
)

func GetIpStr() string {
	ips, err := get_internal()
	log.Println("当前ipv4 ", ips)
	if ips == "" || err != nil {
		log.Println("返回ipv4字符串为随机码 ")
		return RandomStr(8)
	} else {
		arr := strings.Split(ips, ".")
		if len(arr) != 4 {
			return RandomStr(8)
		}
		res := ""
		for _, v := range arr {
			i64, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return RandomStr(8)
			}
			res += dec2Hex(i64)
		}
		log.Println("返回ipv4字符串为16进制码 ", res)
		if len(res) != 8 {
			return RandomStr(8)
		}
		return res
	}
}
func get_internal() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println(ipnet.IP.To4())
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}
func randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}
func RandomStr(l int) string {
	return randomString(l)
}
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func dec2Hex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "00"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}

func FindPath(path string, baseSearchPaths []string, filter func(os.FileInfo) bool) string {
	//判斷是否是絕對路徑
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err == nil {
			return path
		}

		return ""
	}

	searchPaths := []string{}
	for _, baseSearchPath := range baseSearchPaths {
		searchPaths = append(searchPaths, baseSearchPath)
	}

	var binaryDir string
	//返回启动当前进程的可执行文件的路径名称。
	if exe, err := os.Executable(); err == nil {
		if exe, err = filepath.EvalSymlinks(exe); err == nil {
			if exe, err = filepath.Abs(exe); err == nil {
				binaryDir = filepath.Dir(exe)
			}
		}
	}
	if binaryDir != "" {
		for _, baseSearchPath := range baseSearchPaths {
			searchPaths = append(
				searchPaths,
				filepath.Join(binaryDir, baseSearchPath),
			)
		}
	}

	for _, parent := range searchPaths {
		found, err := filepath.Abs(filepath.Join(parent, path))
		if err != nil {
			continue
		} else if fileInfo, err := os.Stat(found); err == nil {
			if filter != nil {
				if filter(fileInfo) {
					return found
				}
			} else {
				return found
			}
		}
	}

	return ""
}

func GetLogFileLocation(fileLocation string) string {
	if fileLocation == "" {
		fileLocation, _ = FindDir("logs")
	}

	return filepath.Join(fileLocation, LOG_FILENAME)
}

func FindDir(dir string) (string, bool) {
	found := FindPath(dir, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
		return fileInfo.IsDir()
	})
	if found == "" {
		return "./", false
	}

	return found, true
}

func FindFile(path string) string {
	return FindPath(path, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir()
	})
}

func checkFileExit(filename string) bool {
	var exit = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exit = false
	}
	return exit
}

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func ConvertExt2ContentType(ext string) (string, error) {

	upext := strings.ToUpper(ext)
	switch upext {
	case "JPG", "JPEG":
		return "image/jpeg", nil
	case "PNG", "GIF":
		return fmt.Sprintf("image/%s", strings.ToLower(ext)), nil
	// application
	case "ZIP", "CVS", "PDF":
		return fmt.Sprintf("application/%s", strings.ToLower(ext)), nil
		// text
	case "HTM", "HTML", "CSS", "XML":
		return fmt.Sprintf("text/%s", strings.ToLower(ext)), nil
	case "TXT":
		return "text/plain", nil
		// video
	case "AVI", "MP4":
		return fmt.Sprintf("video/%s", strings.ToLower(ext)), nil
		// audio
	case "MP3":
		return fmt.Sprintf("audio/%s", strings.ToLower(ext)), nil
		// office
	case "DOC":
		return "application/msword", nil
	case "DOCX":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document", nil
	case "XLS":
		return "application/x-xls", nil
	case "XLSX":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", nil
	case "PPT", "PPS":
		return "application/vnd.ms-powerpoint", nil
	case "SVG":
		return "text/xml", nil
	default:
		return "", errors.New(commons.CUS_ERR_4021)
	}
}
