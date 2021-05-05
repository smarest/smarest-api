package application

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/smarest/smarest-api/infrastructure/persistence"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-api/domain/service"
	"github.com/smarest/smarest-common/util"
	"gopkg.in/gorp.v3"
)

type Bean struct {
	DbMap                       *gorp.DbMap
	PortalService               *PortalService
	OrderService                *service.OrderService
	RestaurantService           *service.RestaurantService
	AreaService                 *service.AreaService
	TableService                *service.TableService
	CategoryService             *service.CategoryService
	CommentService              *service.CommentService
	RestaurantRepository        repository.RestaurantRepository
	RestaurantProductRepository repository.RestaurantProductRepository
	AreaRepository              repository.AreaRepository
	TableRepository             repository.TableRepository
	CategoryRepository          repository.CategoryRepository
	ProductRepository           repository.ProductRepository
	IngredientRepository        repository.IngredientRepository
	OrderRepository             repository.OrderRepository
	CommentRepository           repository.CommentRepository
}

func (bean *Bean) DestroyBean() {

	// Turn off tracing
	bean.DbMap.TraceOff()
	bean.DbMap.Db.Close()
}

func InitBean() (*Bean, error) {
	user := util.GetEnvDefault("DB_USER", "root")
	password := util.GetEnvDefault("DB_PASSWORD", "")
	//	host := util.GetEnvDefault("DB_HOST", "127.0.0.1")
	//	port := util.GetEnvDefault("DB_PORT", "3306")
	//dbName := util.GetEnvDefault("DB_NAME", "anit_pos_server_new")
	dbName := util.GetEnvDefault("DB_NAME", "smarest")
	//	dsn := fmt.Sprintf("%s:%s@unix(%s:%s)/%s?parseTime=true", user, password, host, port,dbName)
	dsn := fmt.Sprintf("%s:%s@unix(/Applications/XAMPP/xamppfiles/var/mysql/mysql.sock)/%s?parseTime=true", user, password, dbName)
	fmt.Printf("dns: %s", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	bean := &Bean{}
	bean.DbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	bean.RestaurantRepository = persistence.NewRestaurantRepository(bean.DbMap)
	bean.RestaurantProductRepository = persistence.NewRestaurantProductRepository(bean.DbMap)
	bean.IngredientRepository = persistence.NewIngredientRepository(bean.DbMap)
	bean.ProductRepository = persistence.NewProductRepository(bean.DbMap)
	bean.AreaRepository = persistence.NewAreaRepository(bean.DbMap)
	bean.TableRepository = persistence.NewTableRepository(bean.DbMap)
	bean.OrderRepository = persistence.NewOrderRepository(bean.DbMap)
	bean.CategoryRepository = persistence.NewCategoryRepository(bean.DbMap)
	bean.CommentRepository = persistence.NewCommentRepository(bean.DbMap)

	bean.OrderService = service.NewOrderService(
		bean.AreaRepository,
		bean.OrderRepository,
		bean.ProductRepository,
		bean.RestaurantProductRepository,
		bean.TableRepository,
		entity.NewOrderFactory())
	bean.RestaurantService = service.NewRestaurantService(bean.RestaurantRepository, bean.RestaurantProductRepository)
	bean.AreaService = service.NewAreaService(bean.AreaRepository)
	bean.TableService = service.NewTableService(bean.AreaRepository, bean.TableRepository)
	bean.CategoryService = service.NewCategoryService(bean.CategoryRepository)
	bean.CommentService = service.NewCommentService(bean.ProductRepository, bean.CommentRepository)
	bean.PortalService = NewPortalService(bean)
	// Will log all SQL statements + args as they are run
	// The first arg is a string prefix to prepend to all log messages
	bean.DbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	return bean, nil
}
