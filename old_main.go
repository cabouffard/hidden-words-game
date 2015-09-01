// // Package main provides ...
package main

//
// import (
// 	"database/sql"
// 	"flag"
// 	"fmt"
// 	_ "github.com/mattn/go-sqlite3"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"time"
// )
//
// type Orientation int
//
// const (
// 	S Orientation = 1 + iota
// 	E
// 	N
// 	// W
// 	// NE
// 	// NW
// 	// SE
// 	// SW
// )
//
// var orientations = [...]string{
// 	"South",
// 	"East",
// 	"North",
// 	// "West",
// 	// "Northeast",
// 	// "Northwest",
// 	// "Southeast",
// 	// "Southwest",
// }
//
// const NilChar rune = '_'
//
// func (o Orientation) String() string { return orientations[o-1] }
//
// func createDatabase() *sql.DB {
// 	os.Remove("./words.db")
//
// 	db, err := sql.Open("sqlite3", "./words.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	sqlStmt := `
// 	create table words (id integer not null primary key, value string);
// 	delete from words;
// 	`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		log.Printf("%q: %s\n", err, sqlStmt)
// 	}
//
// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stmt, err := tx.Prepare("insert into words(id, value) values(?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
// 	// words := []string{"huitres", "moules", "seiche", "crabe", "tarte", "master", "puppet", "josie", "hello", "world", "super", "sayien", "nono", "nul", "annul", "son", "sup"}
//
// 	words := []string{"abaissa", "abaissable", "abaissables", "abaissai", "abaissaient", "abaissais", "abaissait", "abaissames", "abaissant", "abaissante", "abaissantes", "abaissants", "abaissas", "abaissasse", "abaissassent", "abaissasses", "abaissassiez", "abaissassions", "abaissat", "abaissates", "abaisse", "abaisse", "abaissee", "abaissees", "abaisse-langue", "abaissement", "abaissements", "abaissent", "abaisser", "abaissera", "abaisserai", "abaisseraient", "abaisserais", "abaisserait", "abaisseras", "abaisserent", "abaisserez", "abaisseriez", "abaisserions", "abaisserons", "abaisseront", "abaisses", "abaisses", "abaisseur", "abaisseurs", "abaissez", "abaissiez", "abaissions", "abaissons", "abajoue", "abajoues", "abandon", "abandonna", "abandonnai", "abandonnaient", "abandonnais", "abandonnait", "abandonnames", "abandonnant", "abandonnas", "abandonnasse", "abandonnassent", "abandonnasses", "abandonnassiez", "abandonnassions", "abandonnat", "abandonnataire", "abandonnataires", "abandonnates", "abandonne", "abandonne", "abandonnee", "abandonnees", "abandonnent", "abandonner", "abandonnera", "abandonnerai", "abandonneraient", "abandonnerais", "abandonnerait", "abandonneras", "abandonnerent", "abandonnerez", "abandonneriez", "abandonnerions", "abandonnerons", "abandonneront", "abandonnes", "abandonnes", "abandonnez", "abandonniez", "abandonnions", "abandonnons", "abandons", "abaque", "abaques", "abasourdi", "abasourdie", "abasourdies", "abasourdimes", "abasourdir", "abasourdira", "abasourdirai", "abasourdiraient", "abasourdirais", "abasourdirait", "abasourdiras", "abasourdirent", "abasourdirez", "abasourdiriez", "abasourdirions", "abasourdirons", "abasourdiront", "abasourdis", "abasourdissaient", "abasourdissais", "abasourdissait", "abasourdissant", "abasourdissante", "abasourdissantes", "abasourdissants", "abasourdisse", "abasourdissement", "abasourdissements", "abasourdissent", "abasourdisses", "abasourdissez", "abasourdissiez", "abasourdissions", "abasourdissons", "abasourdit", "abasourdit", "abasourdites", "abat", "abatage", "abatages", "abatardi", "abatardie", "abatardies", "abatardimes", "abatardir", "abatardira", "abatardirai", "abatardiraient", "abatardirais", "abatardirait", "abatardiras", "abatardirent", "abatardirez", "abatardiriez", "abatardirions", "abatardirons", "abatardiront", "abatardis", "abatardissaient", "abatardissais", "abatardissait", "abatardissant", "abatardisse", "abatardissement", "abatardissements", "abatardissent", "abatardisses", "abatardissez", "abatardissiez", "abatardissions", "abatardissons", "abatardit", "abatardit", "abatardites", "abatee", "abatees", "abat-jour", "abats", "abattage", "abattages", "abattaient", "abattais", "abattait", "abattant", "abattante", "abattantes", "abattants", "abatte", "abattee", "abattees", "abattement", "abattements", "abattent", "abattes", "abatteur", "abatteurs", "abattez", "abattiez", "abattimes", "abattions", "abattirent", "abattis", "abattisse", "abattissent", "abattisses", "abattissiez", "abattissions", "abattit", "abattit", "abattites", "abattoir", "abattoirs", "abattons", "abattra", "abattrai", "abattraient", "abattrais", "abattrait", "abattras", "abattre", "abattrez", "abattriez", "abattrions", "abattrons", "abattront", "abattu", "abattue", "abattues", "abattus", "abat-vent", "abat-voix", "abbatial", "abbatiale", "abbatiales", "abbatiaux", "abbaye", "abbayes", "abbe", "abbes", "abbesse", "abbesses", "abbevillien", "abbevillienne", "abbevilliennes", "abbevilliens", "abces", "abdicataire", "abdicataires", "abdication", "abdications", "abdiqua", "abdiquai", "abdiquaient", "abdiquais", "abdiquait", "abdiquames", "abdiquant", "abdiquas", "abdiquasse", "abdiquassent", "abdiquasses", "abdiquassiez", "abdiquassions", "abdiquat", "abdiquates", "abdique", "abdique", "abdiquee", "abdiquees", "abdiquent", "abdiquer", "abdiquera", "abdiquerai", "abdiqueraient", "abdiquerais", "abdiquerait", "abdiqueras", "abdiquerent", "abdiquerez", "abdiqueriez", "abdiquerions", "abdiquerons", "abdiqueront", "abdiques", "abdiques", "abdiquez", "abdiquiez", "abdiquions", "abdiquons", "abdomen", "abdomens", "abdominal", "abdominale", "abdominales", "abdominaux", "abducteur", "abducteurs", "abduction", "abductions", "abecedaire", "abecedaires", "abeille", "abeilles", "aberra", "aberrai", "aberraient", "aberrais", "aberrait", "aberrames", "aberrance", "aberrances", "aberrant", "aberrante", "aberrantes", "aberrants", "aberras", "aberrasse", "aberrassent", "aberrasses", "aberrassiez", "aberrassions", "aberrat", "aberrates", "aberration", "aberrations", "aberre", "aberre", "aberrent", "aberrer", "aberrera", "aberrerai", "aberreraient", "aberrerais", "aberrerait", "aberreras", "aberrerent", "aberrerez", "aberreriez", "aberrerions", "aberrerons", "aberreront", "aberres", "aberrez", "aberriez", "aberrions", "aberrons", "abeti", "abetie", "abeties", "abetimes", "abetir", "abetira", "abetirai", "abetiraient", "abetirais", "abetirait", "abetiras", "abetirent", "abetirez", "abetiriez", "abetirions", "abetirons", "abetiront", "abetis", "abetissaient", "abetissais", "abetissait", "abetissant", "abetissante", "abetissantes", "abetissants", "abetisse", "abetissement", "abetissements", "abetissent", "abetisses", "abetissez", "abetissiez", "abetissions", "abetissons", "abetit", "abetit", "abetites", "abhorra", "abhorrai", "abhorraient", "abhorrais", "abhorrait", "abhorrames", "abhorrant", "abhorras", "abhorrasse", "abhorrassent", "abhorrasses", "abhorrassiez", "abhorrassions", "abhorrat", "abhorrates", "abhorre", "abhorre", "abhorree", "abhorrees", "abhorrent", "abhorrer", "abhorrera", "abhorrerai", "abhorreraient", "abhorrerais", "abhorrerait", "abhorreras", "abhorrerent", "abhorrerez", "abhorreriez", "abhorrerions", "abhorrerons", "abhorreront", "abhorres", "abhorres", "abhorrez", "abhorriez", "abhorrions", "abhorrons", "abima", "abimai", "abimaient", "abimais", "abimait", "abimames", "abimant", "abimas", "abimasse", "abimassent", "abimasses", "abimassiez", "abimassions", "abimat", "abimates", "abime", "abime", "abimee", "abimees", "abiment", "abimer", "abimera", "abimerai", "abimeraient", "abimerais", "abimerait", "abimeras", "abimerent", "abimerez", "abimeriez", "abimerions", "abimerons", "abimeront", "abimes", "abimes", "abimez", "abimiez", "abimions", "abimons", "abiotique", "abiotiques", "abject", "abjecte", "abjectement", "abjectes", "abjection", "abjections", "abjects", "abjura", "abjurai", "abjuraient", "abjurais", "abjurait", "abjurames", "abjurant", "abjuras", "abjurasse", "abjurassent", "abjurasses", "abjurassiez", "abjurassions", "abjurat", "abjurates", "abjuration", "abjurations", "abjure", "abjure", "abjuree", "abjurees", "abjurent", "abjurer", "abjurera", "abjurerai", "abjureraient", "abjurerais", "abjurerait", "abjureras", "abjurerent", "abjurerez", "abjureriez", "abjurerions"}
// 	for i := 0; i < len(words); i++ {
// 		_, err = stmt.Exec(i, words[i])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	tx.Commit()
// 	return db
//
// }
//
// func findWord(word string, db *sql.DB) string {
// 	// return "word"
// 	rows, err := db.Query("select value from words where value like ?", word)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var words []string
// 	var name string
// 	for rows.Next() {
// 		rows.Scan(&name)
// 		words = append(words, name)
// 	}
// 	if len(words) > 0 {
// 		return words[random(0, len(words))]
// 	} else {
// 		return "-1"
// 	}
// }
//
// func testGrid(grid [][]rune) {
// 	grid[0][0] = 'a'
// 	grid[0][1] = 'l'
// 	grid[0][2] = 'l'
// 	grid[0][3] = 'o'
// }
//
// func main() {
// 	// NOTE(cab): printing formats
// 	// https://golang.org/pkg/fmt/
// 	db := createDatabase()
// 	sizePtr := flag.Int("size", 5, "an int")
// 	flag.Parse()
//
// 	// Init array
// 	grid := make([][]rune, *sizePtr)
// 	for i := range grid {
// 		grid[i] = make([]rune, *sizePtr)
// 	}
//
// 	// Set array
// 	for x, row := range grid {
// 		for y, _ := range row {
// 			grid[x][y] = NilChar
// 		}
// 	}
//
// 	testGrid(grid)
//
// 	fmt.Printf("STARTS \n")
// 	// fmt.Printf("Grid size: %v \n", *sizePtr)
// 	// printGrid(grid)
//
// 	var usedWords []string
// 	// for freeSpaceInGrid(grid) > 40 {
// 	for i := 0; i < 30; i++ {
// 		posX := random(0, *sizePtr)
// 		posY := random(0, *sizePtr)
// 		// orientation := Orientation(random(1, 4))
// 		orientation := N
//
// 		if grid[posX][posX] == NilChar {
// 			if orientation == E {
// 				// Word must be bigger than 3 chars
// 				for *sizePtr-posX < 3 {
// 					posX = posX - 1
// 				}
// 				// fmt.Printf("Word starts at pos: x: %v, y: %v \n", posX, posY)
//
// 				wordLength := random(3, *sizePtr)
// 				for posX+wordLength > *sizePtr {
// 					wordLength = random(3, *sizePtr)
// 				}
// 				// fmt.Printf("Found word: %v \n", word)
// 				var query string
// 				for i := posX; i < posX+wordLength; i++ {
// 					char := grid[posY][i]
// 					if char == NilChar {
// 						query = query + "_"
// 					} else {
// 						s := fmt.Sprintf("%c", char)
// 						query = query + s
// 					}
// 				}
// 				// fmt.Printf("Query: %v \n", query)
// 				word := findWord(query, db)
// 				if word == "-1" {
// 					continue
// 				}
//
// 				testCount := 0
// 				for stringInSlice(word, usedWords) {
// 					word = findWord(query, db)
// 					if word == "-1" {
// 						word = "-1"
// 					}
// 					if testCount == 3 {
// 						word = "-1"
// 					}
// 					testCount = testCount + 1
// 				}
// 				if word == "-1" {
// 					continue
// 				}
// 				usedWords = append(usedWords, word)
// 				for i := posX; i < posX+wordLength; i++ {
// 					grid[posY][i] = rune([]rune(word)[i-posX])
// 					// printGrid(grid)
// 					// println()
// 				}
//
// 			}
// 			if orientation == N {
// 				// Word must be bigger than 3 chars
// 				for *sizePtr-posY > *sizePtr-3 {
// 					posY = posY + 1
// 				}
//
// 				wordLength := random(3, *sizePtr)
// 				for posY-wordLength < 0 {
// 					wordLength = random(3, *sizePtr)
// 				}
// 				var query string
// 				posX = 0
// 				posY = 4
// 				for i := posY - wordLength; i < posY; i++ {
// 					char := grid[i][posX]
// 					if char == NilChar {
// 						query = query + "_"
// 					} else {
// 						s := fmt.Sprintf("%c", char)
// 						query = query + s
// 					}
// 					// query = reverse(query)
// 				}
// 				word := findWord(query, db)
// 				if word == "-1" {
// 					continue
// 				}
//
// 				testCount := 0
// 				for stringInSlice(word, usedWords) {
// 					word = findWord(query, db)
// 					if word == "-1" {
// 						word = "-1"
// 					}
// 					if testCount == 3 {
// 						word = "-1"
// 					}
// 					testCount = testCount + 1
// 				}
// 				if word == "-1" {
// 					continue
// 				}
// 				usedWords = append(usedWords, word)
// 				fmt.Printf("Query: %v \n", query)
// 				// fmt.Printf("Found word: %v \n", word)
// 				// fmt.Printf("Word starts at pos: x: %v, y: %v \n", posX, posY)
// 				// fmt.Printf("wordLength: %v \n", wordLength)
// 				for i := posY - wordLength; i < posY; i++ {
// 					grid[i][posX] = rune([]rune(word)[posY-i-1])
// 					// printGrid(grid)
// 					// println()
// 				}
//
// 			}
// 			if orientation == S {
// 				// Word must be bigger than 3 chars
// 				for *sizePtr-posY < 3 {
// 					posY = posY - 1
// 				}
//
// 				wordLength := random(3, *sizePtr)
// 				for posY+wordLength > *sizePtr {
// 					wordLength = random(3, *sizePtr)
// 				}
// 				var query string
// 				for i := posY; i < posY+wordLength; i++ {
// 					char := grid[i][posX]
// 					if char == NilChar {
// 						query = query + "_"
// 					} else {
// 						s := fmt.Sprintf("%c", char)
// 						query = query + s
// 					}
// 				}
// 				word := findWord(query, db)
// 				if word == "-1" {
// 					continue
// 				}
//
// 				testCount := 0
// 				for stringInSlice(word, usedWords) {
// 					word = findWord(query, db)
// 					if word == "-1" {
// 						word = "-1"
// 					}
// 					if testCount == 3 {
// 						word = "-1"
// 					}
// 					testCount = testCount + 1
// 				}
// 				if word == "-1" {
// 					continue
// 				}
// 				usedWords = append(usedWords, word)
// 				// fmt.Printf("Query: %v \n", query)
// 				// fmt.Printf("Found word: %v \n", word)
// 				// fmt.Printf("Word starts at pos: x: %v, y: %v \n", posX, posY)
// 				for i := posY; i < posY+wordLength; i++ {
// 					grid[i][posX] = rune([]rune(word)[i-posY])
// 					// printGrid(grid)
// 					// println()
// 				}
//
// 			}
//
// 		}
//
// 		// Show array
// 		// printGrid(grid)
// 	}
// 	printGrid(grid)
// 	printWords(usedWords)
// 	fmt.Printf("Free space left: %d \n", freeSpaceInGrid(grid))
// }
//
// func printWords(words []string) {
// 	for _, word := range words {
// 		println(word)
// 	}
// }
//
// func freeSpaceInGrid(grid [][]rune) int {
// 	// For each rows
// 	count := 0
// 	for _, row := range grid {
// 		// For each columns
// 		for _, value := range row {
// 			if value == NilChar {
// 				count += 1
// 			}
// 		}
// 	}
// 	return count
// }
//
// func printGrid(grid [][]rune) {
// 	// For each rows
// 	for _, row := range grid {
// 		// For each columns
// 		for _, value := range row {
// 			fmt.Printf("[%+q]", value)
// 		}
// 		println()
// 	}
// }
//
// func random(min, max int) int {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	value := rand.Intn(max-min) + min
// 	return value
// }
//
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
//
// func stringInSlice(a string, list []string) bool {
// 	for _, b := range list {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// // func reverse(a string) string {
// // 	for _, v := range a {
// // 		fmt.Printf("(%+q)", v)
// // 	}
// // 	return a
// // }
