package main

import (
    "fmt"
    "encoding/json"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
  //  _ "github.com/go-sql-driver/mysql"
    "log"
  //  "server/handler"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/static"
    "net/http"
    "math/rand"
    "time"
    "os"
 //   "strings"
)

const (
    // データベース
    Dialect = "mysql"

    // ユーザー名
    DBUser = "root"

    // パスワード
    DBPass = "password"

    // プロトコル
    DBProtocol = "tcp(mysql:3306)"
    //DBProtocol = "tcp(127.0.0.1:3306)"

    // DB名
    DBName = "go_sample"

    // オプション???
    DBOption = "?parseTime=true"
)

func connectGorm() *gorm.DB {
    connectTemplate := "%s:%s@%s/%s%s"
    connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName, DBOption)
    db, err := gorm.Open(Dialect, connect)

    if err != nil {
        log.Println(err.Error())
    }

    return db
}

/*
type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
 //   UpdatedAt time.Time
 //   DeletedAt *time.Time `sql:"index"`
}*/

type Book struct {
    //最初は大文字にしないとだめ
    gorm.Model
    Uid string `json:"uid"`
    Title  string `json:"title"`
    Author  string `json:"author"`
    Price string `json:"price"`
    Publisher string `json:"publisher"`
    Published string `json:"published"`
    Imagetype string `json:"imagetype"`
    Image string `json:"image"`
    Url string `json:"url"`
    Memo string `json:"memo"`
}

type Data struct {
 Num int `json:"num"`
 Items []Book `json:"items" `
}


func (b Book) String() string {
 j, _ := json.Marshal(b)
 return string(j)
}


func insert(books []Book, db *gorm.DB) {
    for _, book := range books {
        db.NewRecord(book)
        db.Create(&book)
    }
}

func insertBook( book Book, db *gorm.DB){
    if existByID( book.Uid , db ){//あるときはmemoの更新のみ
     var book1 Book
     db.Model(&book1).Where("uid=?",book.Uid).Update("memo",book.Memo)
    } else {//ないとき
     db.NewRecord(book)
     db.Create(&book)
    }
}

func Count(db *gorm.DB ) int {
 var book Book
 var count = 0
 db.Model(&book).Count(&count)
 return count
}

func findAll(db *gorm.DB) []Book {
    var allBooks []Book
    db.Find(&allBooks)
    return allBooks
}

func existByID(id string , db *gorm.DB) bool {
 var count = 0
 var book Book
 db.Model(&book).Where("uid=?", id ).Count(&count)
 if count == 0{
  return false
 }else{
  return true
 }
}

func deleteByID(id string, db *gorm.DB) {
    var book Book
    db.Where("uid = ?", id).Delete(book)
}

func UpdateByID(id string , memo string , db *gorm.DB ){
    var book Book
    db.Model(&book).Where("uid=?",id).Update("memo",memo)
}

func findByID(id string  , db *gorm.DB ) Book{
 var book Book
 var res Book
 db.Model(&book).Where("uid=?",id).First(&res )
 return res
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
func FilenameByID( id string , db *gorm.DB ) string {
 return findByID(id, db).Image
}

func main() {
    //乱数シード
    rand.Seed(time.Now().UnixNano())

    db := connectGorm()
    defer db.Close()
    db.AutoMigrate(&Book{})//テーブル作成
  // db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Book{})

  /*
    book1 := Book{Id: "abcd" , Title: "okok" , Author: "obata" , Price: "0" , Publisher: "today", Published: "yesterday", Imagetype: "pdf" , Image: "koko" , Url: "localhost"}
    book2 := Book{Id: "abcdeg" , Title: "okok2" , Author: "obata" , Price: "0" , Publisher: "today",Published: "yesterday", Imagetype: "pdf" , Image: "koko" , Url: "localhost"}
    insertBooks := []Book{book1, book2}
    insert(insertBooks, db)
*/



   // routing 
   r := gin.Default()
   r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:8080"},
        AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"*"},
    }))


    r.Use(static.Serve("/", static.LocalFile("./images", true)))


    r.GET("/list", func(c *gin.Context) {
     books := findAll(db)
     data := Data{ Num: Count(db) , Items: books }
     c.JSON(http.StatusOK , data )
    })

    r.GET("/id", func( c *gin.Context ) {
      id := c.Query("id")
      c.JSON(http.StatusOK , findByID(id,db) )
    })

    r.DELETE("/delete", func(c *gin.Context ){
     id := c.Query("id")
     filename := FilenameByID(string(id),db)
     deleteByID(string(id),db)
     //if strings.HasSuffix( filename , "http://localhost:8888/" ) {
     if filename[:22] == "http://localhost:8888/"{
      os.Remove("images/"+filename[22:])
     }
    })

     r.POST("/create" , func( c *gin.Context ){
    /* id := c.Query("id")
     title := c.Query("title")
     author := c.Query("author")
     price := c.Query("price")
     publisher := c.Query("publisher")
     published := c.Query("published")
     imagetype:= c.Query("imagetype")
     image := c.Query("image")
     url := c.Query("url")
     memo := c.Query("memo")
     */
     var book Book
     c.BindJSON(&book)
     //book := Book{Uid: id, Title: title , Author: author , Price: price,Publisher: publisher, Published: published, Imagetype: imagetype, Image: image, Url: url , Memo: memo}
    //insertBooks := []Book{book}
     insertBook(book , db)
     //c.JSON(http.StatusOK , book )
    })

     r.POST("/create_noid" , func( c *gin.Context ){
     id := RandomString(10)
     var book Book
     c.BindJSON(&book)
     book.Uid = id
     //book := Book{Uid: id, Title: title , Author: author , Price: price,Publisher: publisher, Published: published, Imagetype: imagetype, Image: image, Url: url , Memo: memo}
    //insertBooks := []Book{book}
     insertBook(book , db)
     //c.JSON(http.StatusOK , book )
    })


    r.POST("/update" , func( c *gin.Context ){
     id := c.Query("id")
     memo := c.Query("memo")
     //book := Book{Uid: id, Memo: memo}
     UpdateByID(id,memo,db)
     //c.JSON(http.StatusOK , book )
    })

    r.POST("/images", func( c *gin.Context ){
     form, _ := c.MultipartForm()
     files := form.File["file"]
     filename := c.Query("filename")
     //file := form.File["file"]
     // uuid を所得
     //uuid := c.PostForm("uuid")
     for _, file := range files {
        err := c.SaveUploadedFile(file, "images/"+filename )
        //err := c.SaveUploadedFile(file, "images/"+file.Filename)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        }
    }

    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
    //c.JSON(http.StatusOK, gin.H{"message": "success!!"})
    })

    r.Run(":8888")

}


