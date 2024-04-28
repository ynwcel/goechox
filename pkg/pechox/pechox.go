package pechox

import (
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	engine  *echo.Echo
	setting map[string]any
}

func New(cfg ...map[string]any) *Server {
	var engine = echo.New()

	engine.HideBanner = true
	engine.Pre(middleware.RemoveTrailingSlash(), GctxMiddleware())

	var setting = make(map[string]any)
	if len(cfg) > 0 && len(cfg[0]) > 0 {
		setting = cfg[0]
	}

	return &Server{
		engine:  engine,
		setting: setting,
	}
}

func (s *Server) Engine() *echo.Echo {
	return s.engine
}

func (s *Server) Run() error {
	setting := s.merge_setting()
	address := fmt.Sprintf("0.0.0.0:%d", setting.Listen)
	s.engine.Debug = setting.Debug
	if setting.Debug {
		return s.engine.Start(address)
	} else {
		svr := &http.Server{
			Addr:           address,
			Handler:        s.engine,
			ReadTimeout:    setting.ReadTimeout,
			WriteTimeout:   setting.WriteTimeout,
			IdleTimeout:    setting.IdleTimeout,
			MaxHeaderBytes: setting.MaxHeaderBytes,
		}
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	}
}

func (s *Server) merge_setting() *cfg {
	var (
		s_cfg   = default_cfg()
		setting = new(cfg)
	)
	if err := gconv.Scan(s.setting, &setting); err == nil {
		s_cfg.Debug = setting.Debug
		if setting.Listen > 0 {
			s_cfg.Listen = setting.Listen
		}
		if setting.ReadTimeout > 0 {
			s_cfg.ReadTimeout = setting.ReadTimeout
		}
		if setting.WriteTimeout > 0 {
			s_cfg.WriteTimeout = setting.WriteTimeout
		}
		if setting.IdleTimeout > 0 {
			s_cfg.IdleTimeout = setting.IdleTimeout
		}
		if setting.MaxHeaderBytes > 0 {
			s_cfg.MaxHeaderBytes = setting.MaxHeaderBytes
		}
	}
	return s_cfg
}
