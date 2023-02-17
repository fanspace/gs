package commons

import (
	"regexp"
	"strings"
	"unicode"
)

// 密码是否符合要求
func IsPwdFormatValid(text string) bool {
	// golang 的正则太弱智， 无法编译负向零宽断言
	r, _ := regexp.Compile(`^[A-Za-z0-9]{8,16}$`)
	lengthok := r.MatchString(text)
	r2, _ := regexp.Compile(`[A-Za-z]{1,15}`)
	alphaok := r2.MatchString(text)
	r3, _ := regexp.Compile(`[0-9]{1,15}`)
	numok := r3.MatchString(text)
	return lengthok && alphaok && numok
}

// 上传的文件是否为多重后缀
func IsMultiExt(text string) bool {
	extarr := strings.Split(text, ".")
	if len(extarr) != 2 {
		return true
	}
	return false
}

// 是否符合条件的图片
func IsImage(text string) bool {
	name := strings.ToLower(text)
	r, _ := regexp.Compile(`.*\.(jpg|jpeg|png|gif)$`)
	IsImage := r.MatchString(name)
	return IsImage

}

//  是否符合条件的上传
func IsValidFile(text string) bool {
	if text == "" {
		return false
	}
	name := strings.ToLower(text)
	r, _ := regexp.Compile(`.*\.(jpg|jpeg|png|gif|doc|xls|zip|pdf|docx|xlsx|ppt|pptx)$`)
	IsValidFile := r.MatchString(name)
	return IsValidFile

}

//邮箱
func IsEmailAddress(text string) bool {
	r, _ := regexp.Compile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`)
	ismail := r.MatchString(text)
	return ismail
}

// 手机
func IsMobile(text string) bool {
	r, _ := regexp.Compile(`^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`)
	ismobile := r.MatchString(text)
	return ismobile
}

//判断中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// 判断全角
func IsFullWidth(str string) bool {
	r, _ := regexp.Compile(`[^\x00-\xff]`)
	return r.MatchString(str)
}

// 身份证，正则方式，简化
func IsIdCardSimple(text string) bool {
	r, _ := regexp.Compile(`^\d{6}(18|19|20)?\d{2}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])\d{3}(\d|[xX])$`)
	isid := r.MatchString(text)
	return isid

}

// 身份证，按规则，严格
func IsIdCardStrict(text string) bool {
	if text == "" {
		return false
	}
	if len(text) != 18 {
		return false
	}
	if isok := IsIdCardSimple(text); !isok {
		return false
	}
	tal := 0
	offset := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	a := []rune(text[0:17])
	last := text[17:]
	for i := 0; i < len(a); i++ {
		tal = tal + int(a[i]-48)*offset[i]
	}
	basenum := tal % 11
	switch basenum {
	case 0:
		if last == "1" {
			return true
		} else {
			return false
		}
	case 1:
		if last == "0" {
			return true
		} else {
			return false
		}
	case 2:
		if last == "x" || last == "X" {
			return true
		} else {
			return false
		}
	case 3:
		if last == "9" {
			return true
		} else {
			return false
		}
	case 4:
		if last == "8" {
			return true
		} else {
			return false
		}
	case 5:
		if last == "7" {
			return true
		} else {
			return false
		}
	case 6:
		if last == "6" {
			return true
		} else {
			return false
		}
	case 7:
		if last == "5" {
			return true
		} else {
			return false
		}
	case 8:
		if last == "4" {
			return true
		} else {
			return false
		}
	case 9:
		if last == "3" {
			return true
		} else {
			return false
		}
	case 10:
		if last == "2" {
			return true
		} else {
			return false
		}
	default:
		return false

	}

}

// passport
func IsPassport(text string) bool {
	r, _ := regexp.Compile(`^([a-zA-z]|[0-9]){5,17}$`)
	isid := r.MatchString(text)
	return isid

}

// Tw Card
func IsTwCard(text string) bool {
	r, _ := regexp.Compile(`^\d{8}|^[a-zA-Z0-9]{10}|^\d{18}$`)
	isid := r.MatchString(text)
	return isid

}

// officail card
func IsOfficailCard(text string) bool {
	r, _ := regexp.Compile(`^[a-zA-Z0-9]{7,21}$`)
	isid := r.MatchString(text)
	return isid

}

// 判断idstr 是否合理， 原型  1,2,3
func IsIdStr(text string) bool {
	st := strings.Index(text, ",")
	if st == 0 || st == (len(text)-1) {
		return false
	}
	r, _ := regexp.Compile(`^\d+$`)
	text2 := strings.Replace(text, ",", "", -1)
	IsIdStr := r.MatchString(text2)
	return IsIdStr
}

// 验证ip 地址
func IsIpv4Address(text string) bool {
	if text == "" {
		return false
	}
	r, _ := regexp.Compile(`^([1-9]\d?|1\d{2}|2[0-4]\d|25[0-5])(\.([1-9]?\d|1\d{2}|2[0-4]\d|25[0-5])){3}$`)
	return r.MatchString(text)
}

// 是否初始密码 8位英文
func IsDefaultPwd(text string) bool {
	r, _ := regexp.Compile(`^[A-Z]{8}$`)
	return r.MatchString(text)
}
