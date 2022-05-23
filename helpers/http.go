package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HttpJsonResponse(w http.ResponseWriter, r *http.Request, httpStatusCode int, httpResponse interface{}) {
	message, err := json.Marshal(httpResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	HttpLog(w, r, httpStatusCode)
	w.WriteHeader(httpStatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

/*
# Utilizing NGINX Log Format
'$remote_addr - $remote_user [$time_local] "$request" '
'$status $body_bytes_sent "$http_referer" '
'"$http_user_agent" "$http_x_forwarded_for"';

# Sample logs
10.1.179.128 - - [09/Apr/2022:18:09:07 +0000] "GET /phpMyAdmin-4/index.php?lang=en HTTP/1.1" 200 206 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36" "46.38.30.151"
10.1.179.128 - - [09/Apr/2022:16:12:13 +0000] "GET /sql/phpmyadmin4/index.php?lang=en HTTP/1.1" 200 209 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36" "153.131.28.205"
192.168.86.6 - - [09/Apr/2022:15:35:50 +0000] "GET /.aws/credentials HTTP/1.1" 200 194 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36" "109.237.103.38"
10.1.179.128 - - [10/Apr/2022:13:52:27 +0000] "GET / HTTP/1.1" 200 - http://www.google.com.hk "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.90 Safari/537.36" "40.159.86.203"


# Parsed by Logstash Grok patterns
"%{IPORHOST:clientip} %{NGUSERNAME:ident} %{NGUSERNAME:auth} \[%{HTTPDATE:timestamp}\] \"%{WORD:verb} %{URIPATHPARAM:request} HTTP/%{NUMBER:httpversion}\" %{NUMBER:response} (?:%{NUMBER:bytes}|-) (?:\"(?:%{URI:referrer}|-)\"|%{QS:referrer}) %{QS:agent} \"%{IPORHOST:http_x_forwarded_for}\""
NGUSERNAME [a-zA-Z\.\@\-\+_%]+
*/
func HttpLog(w http.ResponseWriter, r *http.Request, httpStatusCode int) {
	// 	logger.Info.Printf("%v received %v call", r.URL.Path, r.Method)
	fmt.Printf("%s - - [%s] \"%s %s %s\" %d - %s %q %q\n",
		HttpHelperRemovePort(r.RemoteAddr),
		time.Now().Format("02/Jan/2006:15:04:05 -0700"),
		r.Method,
		r.URL.Path,
		r.Proto,
		httpStatusCode,
		HttpHelperReturnReferer(r),
		r.Header.Get("User-Agent"),
		HttpHelperReturnClientIp(r),
	)
}

// if r.RemoteAddr contains port, remove it like:
// "127.0.0.1:53956" => "127.0.0.1"
func HttpHelperRemovePort(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

func HttpHelperReturnClientIp(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return HttpHelperRemovePort(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		return parts[0]
	}
	return hdrRealIP
}

func HttpHelperReturnReferer(r *http.Request) string {
	headerReferer := r.Header.Get("Referer")
	if headerReferer == "" {
		// return "\"-\""
		return fmt.Sprintf("%q", "-")
	} else {
		return fmt.Sprintf("%q", headerReferer)
	}
}

func HttpRoute404(w http.ResponseWriter, r *http.Request) {
	message := map[string]string{"404": r.URL.Path}
	HttpJsonResponse(w, r, 404, message)
}

func HttpRouteDefault(w http.ResponseWriter, r *http.Request) {
	message := map[string]interface{}{
		"host_ip":        GetLocalIP(),
		"host_name":      GetLocalHostname(),
		"host_user":      GetLocalUsername(),
		"server_address": r.Host,
		"uri":            r.URL.Path,
		"uuid":           uuid.New(),
		"date":           TimeNow(),
		"remote_agent":   r.UserAgent(),
		"remote_ip":      HttpHelperRemovePort(r.RemoteAddr),
	}
	HttpJsonResponse(w, r, 200, message)
}

// not using jsonResponse, but writing byte style.  Doing this as an example
func HttpRoutePing(w http.ResponseWriter, r *http.Request) {
	HttpLog(w, r, 200)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ping":"pong"}`))
}
