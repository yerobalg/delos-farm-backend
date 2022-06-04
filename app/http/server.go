package main

import (
	// "delos-farm-backend/alamat"
	// "delos-farm-backend/config"
	// "delos-farm-backend/daerah"
	// "delos-farm-backend/handlers"
	// "delos-farm-backend/kategori"
	// "delos-farm-backend/keranjang"
	"delos-farm-backend/middlewares"
	// "delos-farm-backend/pesanan"
	// "delos-farm-backend/produk"
	// "delos-farm-backend/role"
	// "delos-farm-backend/user"
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"delos-farm-backend/bootstrap"
)

func main() {

	//init env variable
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("failed to load env from local file")
	}

	//init database
	_, err := bootstrap.InitDB()
	if err != nil {
		panic(err)
	}
	
	//init router
	engine := gin.Default()
	engine.Use(middlewares.CorsMiddleware())
	// router.Gin.Static("/public/images/users", "./public/images/users")
	// router.Gin.Static("/public/images/products", "./public/images/products")
	
	//init middleware
	// middleware := middlewares.Middleware{}
	// roleMiddleware := middlewares.RoleMiddleware{}

	//init role
	// roleRepo := role.NewRoleRepository(db)
	// roleService := role.NewRoleService(roleRepo)

	//init auth
	// userRepo := user.NewUserRepository(db)
	// userService := user.NewUserService(userRepo)
	// authHandler := handlers.NewAuthHandler(router, userService, roleService)

	//init daerah
	// daerahRepo := daerah.NewDaerahRepository(db)
	// daerahService := daerah.NewDaerahService(daerahRepo)
	// daerahHandler := handlers.NewDaerahHandler(
	// 	router,
	// 	daerahService,
	// 	middleware,
	// )

	//init alamat
	// alamatRepo := alamat.NewAlamatRepository(db)
	// alamatService := alamat.NewAlamatService(alamatRepo)
	// alamatHandler := handlers.NewAlamatHandler(
	// 	router,
	// 	alamatService,
	// 	roleService,
	// 	middleware,
	// )

	//init kategori
	// kategoriRepo := kategori.NewKategoriRepository(db)
	// kategoriService := kategori.NewKategoriService(kategoriRepo)
	// kategoriHandler := handlers.NewKategoriHandler(
	// 	router,
	// 	kategoriService,
	// 	middleware,
	// )

	//init produk
	// produkRepo := produk.NewProdukRepository(db)
	// produkService := produk.NewProdukService(produkRepo)
	// produkHandler := handlers.NewProdukHandler(
	// 	router,
	// 	produkService,
	// 	kategoriService,
	// 	middleware,
	// )

	//init keranjang
	// keranjangRepo := keranjang.NewKeranjangRepository(db)
	// keranjangService := keranjang.NewKeranjangService(keranjangRepo)
	// keranjangHandler := handlers.NewKeranjangHandler(
	// 	router,
	// 	keranjangService,
	// 	userService,
	// 	alamatService,
	// 	produkService,
	// 	middleware,
	// )

	//init pesanan
	// pesananRepo := pesanan.NewPesananRepository(db)
	// pesananService := pesanan.NewPesananService(pesananRepo)
	// pesananHandler := handlers.NewPesananHandler(
	// 	router,
	// 	pesananService,
	// 	keranjangService,
	// 	alamatService,
	// 	produkService,
	// 	middleware,
	// )

	//setup router
	// handlers.NewHandlers(
	// 	authHandler,
	// 	daerahHandler,
	// 	alamatHandler,
	// 	kategoriHandler,
	// 	produkHandler,
	// 	keranjangHandler,
	// 	pesananHandler,
	// ).Setup()

	//run server
	engine.Run(os.Getenv("PORT"))
}
