package bootstrap

import (
	"cms-server/internal/entity"
	pkglog "cms-server/pkg/logger"

	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env         *Env
	DB          *pg.DB
	Log         pkglog.Logger
	QueneClient QueueClient
}

func App() *Application {
	env := Env{}
	NewEnv(&env)

	logConfig := pkglog.NewConfig()
	log := pkglog.InitLogger(logConfig, zapcore.DebugLevel, env.IsProduction())

	qc := NewQueueClient(&env)

	entities := []interface{}{
		new(entity.User),
		new(entity.Media),
		new(entity.Session),
		new(entity.Role),
		new(entity.UserRole),
		new(entity.Module),
		new(entity.ModuleChild),
		new(entity.ModuleChildRole),
		new(entity.WebSetting),
		new(entity.Banner),
		new(entity.MethodPayment),
		new(entity.ActivityLog),
		new(entity.TypeMail),
		new(entity.MailProvider),
		new(entity.MailTemplate),
		new(entity.MailStatus),
		new(entity.MailHistory),
		new(entity.StatusHistory),
		new(entity.Category),
		new(entity.Post),
		new(entity.Tag),
		new(entity.PostTag),
		new(entity.Comment),
		new(entity.Like),
		new(entity.Coupon),
		new(entity.Product),
		new(entity.Product),
		new(entity.ProductVariant),
		new(entity.Attribute),
		new(entity.AttributeValue),
		new(entity.ProductAttribute),
		new(entity.VariantValue),
		new(entity.Cart),
		new(entity.CartItem),
		new(entity.Order),
		new(entity.StatusOrder),
		new(entity.OrderItem),
		new(entity.OrderStatusHistory),
		new(entity.Supplier),
		new(entity.Warehouse),
		new(entity.Menu),
	}

	db := NewPostgresDB(&env, entities, log)

	RegisterValidator()
	return &Application{
		Env:         &env,
		DB:          db,
		Log:         log,
		QueneClient: qc,
	}
}
