package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"core-gin/constants"
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// Global store
// using in-memory store with goroutine which clears expired keys.
var store = memory.NewStore()

type RateLimitOption struct {
	period time.Duration
	limit  int64
}

type Option func(*RateLimitOption)

type RateLimitMiddleware struct {
	logger lib.Logger
	option RateLimitOption
	env    *lib.Env
	router infrastructure.Router
}

func NewRateLimitMiddleware(logger lib.Logger, env *lib.Env, router infrastructure.Router) RateLimitMiddleware {
	return RateLimitMiddleware{
		logger: logger,
		env:    env,
		router: router,
		option: RateLimitOption{
			period: time.Duration(env.RateLimitPeriod) * time.Minute,
			limit:  env.RateLimitRequests,
		},
	}
}

func (lm RateLimitMiddleware) Setup() {
	lm.logger.Info("Setting up rate limit middleware")

	lm.router.Use(lm.Handle())
}

func (lm RateLimitMiddleware) Handle(options ...Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() // Gets cient IP Address

		// Setting up rate limit
		// Limit -> # of API Calls
		// Period -> in a given time frame
		// setting default values
		opt := RateLimitOption{
			period: lm.option.period,
			limit:  lm.option.limit,
		}

		for _, o := range options {
			o(&opt)
		}

		rate := limiter.Rate{
			Limit:  opt.limit,
			Period: opt.period,
		}

		// Limiter instance
		instance := limiter.New(store, rate)

		// Returns the rate limit details for given identifier.
		// FullPath is appended with IP address. `/api/users&&127.0.0.1` as key
		context, err := instance.Get(c, c.FullPath()+"&&"+key)
		if err != nil {
			lm.logger.Error(err)
		}

		c.Set(constants.RateLimit, instance)

		// Setting custom headers
		c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

		// Limit exceeded
		if context.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate Limit Exceed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func WithOptions(period time.Duration, limit int64) Option {
	return func(o *RateLimitOption) {
		o.period = period
		o.limit = limit
	}
}
