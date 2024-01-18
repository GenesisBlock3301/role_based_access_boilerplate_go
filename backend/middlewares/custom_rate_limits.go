package middlewares

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type clientInfo struct {
	// Track the last request of this IP
	lastRequestTime time.Time
	// count this request
	requestCount int
}

type CustomRateLimiter struct {
	mu        sync.Mutex
	rateLimit int
	// maximum time
	duration time.Duration
	clients  map[string]*clientInfo
}

// NewCustomRateLimiter create a new instance of the CustomRateLimiter
func NewCustomRateLimiter(rateLimit int, duration time.Duration) *CustomRateLimiter {
	return &CustomRateLimiter{
		rateLimit: rateLimit,
		duration:  duration,
		clients:   make(map[string]*clientInfo),
	}
}

func (limiter *CustomRateLimiter) CustomRateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		// lock others go routines
		limiter.mu.Lock()
		// end of all operation unlock this go routine
		defer limiter.mu.Unlock()

		client, exists := limiter.clients[clientIP]

		now := time.Now()
		// if IP is new
		if !exists {
			client = &clientInfo{
				lastRequestTime: now,
			}
			limiter.clients[clientIP] = client
		}

		// Check the client exceed the time limit or not.
		if now.Sub(client.lastRequestTime) < limiter.duration {
			client.requestCount++
			if client.requestCount > limiter.rateLimit {
				c.AbortWithStatusJSON(429, gin.H{"error": "Rate limit exceeded."})
				return
			}
		} else {
			client.requestCount = 1
			client.lastRequestTime = now
		}

	}
}
