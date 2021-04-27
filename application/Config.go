package application

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/smarest/smarest-api/infrastructure/persistence"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/service"
	"github.com/smarest/smarest-common/util"
	"gopkg.in/gorp.v3"
)

type Bean struct {
	DbMap             *gorp.DbMap
	RestaurantService *RestaurantService
	IngredientService *IngredientService
	ProductService    *ProductService
	AreaService       *AreaService
	TableService      *TableService
	CategoryService   *CategoryService
	CommentService    *CommentService
	OrderService      *OrderService
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
	dbName := util.GetEnvDefault("DB_NAME", "anit_pos_server_new")
	//	dsn := fmt.Sprintf("%s:%s@unix(%s:%s)/%s?parseTime=true", user, password, host, port,dbName)
	dsn := fmt.Sprintf("%s:%s@unix(/Applications/XAMPP/xamppfiles/var/mysql/mysql.sock)/%s?parseTime=true", user, password, dbName)
	fmt.Printf("dns: %s", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	restaurantRepository := persistence.NewRestaurantRepository(dbMap)
	ingredientRepository := persistence.NewIngredientRepository(dbMap)
	productRepository := persistence.NewProductRepository(dbMap)
	areaRepository := persistence.NewAreaRepository(dbMap)
	tableRepository := persistence.NewTableRepository(dbMap)
	orderRepository := persistence.NewOrderRepository(dbMap)
	categoryRepository := persistence.NewCategoryRepository(dbMap)
	commentRepository := persistence.NewCommentRepository(dbMap)

	restaurantFactory := entity.NewRestaurantFactory()
	ingredientFactory := entity.NewIngredientFactory()
	productFactory := entity.NewProductFactory()
	areaFactory := entity.NewAreaFactory()
	tableFactory := entity.NewTableFactory()
	commentFactory := entity.NewCommentFactory()
	orderFactory := entity.NewOrderFactory()

	restaurantService := NewRestaurantService(
		restaurantRepository,
		productRepository,
		areaRepository,
		restaurantFactory,
		productFactory,
		areaFactory)
	areaService := NewAreaService(
		areaRepository,
		areaFactory,
		tableRepository,
		tableFactory,
		orderRepository)

	tableService := NewTableService(
		tableRepository,
		tableFactory)

	productService := NewProductService(
		productRepository,
		productFactory)

	ingredientService := NewIngredientService(
		ingredientRepository,
		ingredientFactory)
	categoryService := NewCategoryService(categoryRepository)
	commentService := NewCommentService(
		commentRepository,
		commentFactory)

	domainOrderService := service.NewOrderService(
		orderRepository,
		orderFactory,
		restaurantRepository,
		tableRepository)
	orderService := NewOrderService(
		domainOrderService,
		orderRepository)
	// Will log all SQL statements + args as they are run
	// The first arg is a string prefix to prepend to all log messages
	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	return &Bean{DbMap: dbMap,
		RestaurantService: restaurantService,
		IngredientService: ingredientService,
		ProductService:    productService,
		AreaService:       areaService,
		TableService:      tableService,
		CategoryService:   categoryService,
		CommentService:    commentService,
		OrderService:      orderService}, nil
}
