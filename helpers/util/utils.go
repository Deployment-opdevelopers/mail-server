package util

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/mail"
	"net/url"
	"strconv"
	"strings"

	"opdevelopers.live/user/config"
	"opdevelopers.live/user/constants"
	"opdevelopers.live/user/helpers/request"
)

func StructToJSON(val interface{}) interface{} {
	jsonEncoded, _ := json.Marshal(val)
	var respJSON interface{}
	json.Unmarshal([]byte(jsonEncoded), &respJSON)
	return respJSON
}
func IsAdmin(mailData *request.VerificationData) bool {
	ok := true
	if mailData.AdminId != config.Get().MailAdminId {
		ok = false
	}
	if mailData.Password != config.Get().MailPassword {
		ok = false
	}
	return ok
}
func GetHexToken(length int) (string, error) {
	var token string
	codeAlphabet := constants.CODE_ALPHABET_SHORT
	max := len(codeAlphabet) - 1
	for i := 0; i < length; i++ {
		j, err := CryptoRandSecure(0, max)
		if err != nil {
			return "", err
		}
		token += string(codeAlphabet[j])
	}
	return token, nil
}

func CryptoRandSecure(min int, max int) (int, error) {
	rangee := max - min
	if rangee < 1 {
		return min, nil
	}

	log := math.Ceil(math.Log2((float64)(rangee)))

	bytes := (int)(log/8) + 1    // length in bytes
	bits := (int)(log) + 1       // length in bits
	filter := (int)(1<<bits) - 1 // set all lower bits to 1

	hex := RandomPseudoBytes(bytes)

	rndTemp, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	rnd := (int)(rndTemp) & filter // discard irrelevant bits
	for rnd >= rangee {
		hex = RandomPseudoBytes(bytes)
		rndTemp, err = strconv.ParseInt(hex, 16, 64)
		if err != nil {
			return 0, err
		}
		rnd = (int)(rndTemp) & filter // discard irrelevant bits
	}

	return min + rnd, nil
}

func RandomPseudoBytes(length int) string {
	buff := make([]byte, length)
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str
}

func GetTokenUtility(length int) (string, error) {
	var token string
	codeAlphabet := constants.CODE_ALPHABET_LONG
	max := len(codeAlphabet) - 1
	for i := 0; i < length; i++ {
		j, err := CryptoRandSecure(0, max)
		if err != nil {
			return "", err
		}
		token += string(codeAlphabet[j])
	}
	return token, nil
}

func Contains(str []string, key string) bool {
	for _, v := range str {
		if v == key {
			return true
		}
	}
	return false
}

func IsInteger(val float64) bool {
	return math.Floor(val) == math.Ceil(val)
}

func ValidEmail(email string) (string, bool) {
	mailId, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}
	return mailId.Address, true
}

// GenerateRandomString generate a string of random characters of given length
func GenerateRandomString(n int) string {
	const letterBytes = "1234567890"
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		idx := rand.Int63() % int64(len(letterBytes))
		sb.WriteByte(letterBytes[idx])
	}
	return sb.String()
}

func CallAPI(URL string, method string, ContentType string, payLoad map[string]interface{}, header map[string]interface{}) (string, error) {
	response := ""
	if method == http.MethodPost {
		body, err := json.Marshal(payLoad)
		if err != nil {
			return "", err
		}
		client := &http.Client{}
		var req *http.Request

		if ContentType == constants.CONTENT_TYPE_XML_ENCODED {
			form := url.Values{}
			for k, v := range payLoad {
				form.Add(k, v.(string))
			}
			req, err = http.NewRequest(http.MethodPost, URL, strings.NewReader(form.Encode()))
		} else {
			req, err = http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(body))
		}
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", ContentType)

		for key, value := range header {
			req.Header.Set(key, value.(string))
		}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}
		response = string(responseBody)

	} else if method == http.MethodGet {
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodGet, URL, nil)
		if err != nil {
			return "", err
		}

		// appending to existing query args
		q := req.URL.Query()
		for key, value := range payLoad {
			q.Add(key, fmt.Sprintf("%v", value))
		}

		// assign encoded query string to http request
		req.URL.RawQuery = q.Encode()

		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		response = string(responseBody)
	}

	return response, nil
}
