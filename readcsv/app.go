package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/go-pg/pg/v9"
    "github.com/go-pg/pg/v9/orm"
)

//Client you can use "-" to ignore a field e.g. NotUsed string `csv:"-"`
type Client struct {
	Name        string `csv:"Name"`
	Description string `csv:"Description"`
	Ring        int    `csv:"Ring"`
	Quadrant    string `csv:"Quadrant"`
	IsNew       bool   `csv:"is_new"`
}

func main2() {
	file, err := os.Open("csv/CDS.csv")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line + "...................----")
	}
	fmt.Println("main...!")
}
func main() {
	csvFile, err := os.OpenFile("csv/CDS.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	csvLines := []*Client{}
	if err := gocsv.UnmarshalFile(csvFile, &csvLines); err != nil { // Load  from file
		panic(err)
	}
	for _, line := range csvLines {
		fmt.Println("Hello", line.Name)
	}
}

// func main(){
// 	db := pg.Connect(&pg.Options{
// 		User: "postgres",
// 		Password: "test1234",
// 	})
// 	defer db.Close()

// 	fmt.Println(res.RowsAffected())

// 	err := createSchema(db)
//     if err != nil {
//         panic(err)
//     }
// }
func createSchema(db *pg.DB) error {
    for _, model := range []interface{}{(*User)(nil), (*Story)(nil)} {
        err := db.CreateTable(model, &orm.CreateTableOptions{
            Temp: false,
        })
        if err != nil {
            return err
        }
    }
    return nil
}
type User struct {
    Id     int64
    Name   string
    Emails []string
}

func (u User) String() string {
    return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
    Id       int64
    Title    string
    AuthorId int64
    Author   *User
}

func (s Story) String() string {
    return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}