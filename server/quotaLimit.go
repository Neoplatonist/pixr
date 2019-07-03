package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type quotaLimitClient struct {
	quota    int
	lastSeen time.Time
}

type quotaLimiter struct {
	maxUploads int
	limitBan   int
	clients    map[string]*quotaLimitClient
	m          sync.Mutex
}

func (ql *quotaLimiter) addClient(ip string, quota int) int {
	ql.m.Lock()
	ql.clients[ip] = &quotaLimitClient{quota, time.Now()}
	ql.m.Unlock()

	return quota
}

func (ql *quotaLimiter) getClient(ip string, fileCount int) int {
	ql.m.Lock()
	quotaLimitClient, exists := ql.clients[ip]

	if !exists {
		ql.m.Unlock()
		return ql.addClient(ip, fileCount)
	}

	quotaLimitClient.quota = fileCount + quotaLimitClient.quota
	quotaLimitClient.lastSeen = time.Now()
	ql.m.Unlock()

	return quotaLimitClient.quota
}

func (ql *quotaLimiter) limit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return errors.Wrap(err, "returning multipart form")
		}

		files := form.File["file"]
		limiter := ql.getClient(c.RealIP(), len(files))

		if limiter > ql.maxUploads {
			return c.HTML(http.StatusTooManyRequests, http.StatusText(429))
		}

		return next(c)
	}
}

func (ql *quotaLimiter) cleanQuotaLimiter() {
	for {
		time.Sleep(time.Minute)

		ql.m.Lock()
		for ip, quotaLimitClient := range ql.clients {
			if time.Now().Sub(quotaLimitClient.lastSeen) > time.Duration(ql.limitBan)*time.Minute {
				delete(ql.clients, ip)
			}
		}
		ql.m.Unlock()
	}
}
