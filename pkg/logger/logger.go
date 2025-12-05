package logger

import (
    "os"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func InitLogger(env string) {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    zerolog.SetGlobalLevel(zerolog.InfoLevel)
    if env == "development" {
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
        log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    }
}

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        statusCode := c.Writer.Status()
        log.Info().
            Int("status", statusCode).
            Str("method", c.Request.Method).
            Str("path", c.Request.URL.Path).
            Str("ip", c.ClientIP()).
            Msg(strconv.Itoa(statusCode) + " " + c.Request.Method + " " + c.Request.URL.Path)
    }
}
