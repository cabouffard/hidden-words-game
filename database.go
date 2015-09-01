package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type ORM struct {
	db       *sql.DB
	filename string
}

func InitDatabase(filename string) *ORM {
	var db *sql.DB
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		db = CreateDatabase(filename)
	} else {
		db, err = sql.Open("sqlite3", filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &ORM{db: db, filename: filename}
}

func (orm *ORM) Reset() *ORM {
	os.Remove(orm.filename)
	return InitDatabase(orm.filename)
}

func (orm *ORM) FindWord(query string) (string, error) {
	rows, err := orm.db.Query("select value from words where value like ?", query)
	if err != nil {
		log.Fatal(err)
	}
	var words []string
	var value string
	for rows.Next() {
		rows.Scan(&value)
		words = append(words, value)
	}
	if len(words) > 0 {
		return words[random(0, len(words))], nil
	} else {
		return "", fmt.Errorf("Unable to find a word corresponding to the query: %v", query)
	}
}

func CreateDatabase(filename string) *sql.DB {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	createStatement := `
    create table words (id integer not null primary key, value string);
    delete from words;
    `
	_, err = db.Exec(createStatement)

	if err != nil {
		log.Printf("%q: %s\n", err, createStatement)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into words(id, value) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	words := []string{"abaissa", "abaissable", "abaissables", "abaissai", "abaissaient", "abaissais", "abaissait", "abaissames", "abaissant", "abaissante", "abaissantes", "abaissants", "abaissas", "abaissasse", "abaissassent", "abaissasses", "abaissassiez", "abaissassions", "abaissat", "abaissates", "abaisse", "abaisse", "abaissee", "abaissees", "abaisse-langue", "abaissement", "abaissements", "abaissent", "abaisser", "abaissera", "abaisserai", "abaisseraient", "abaisserais", "abaisserait", "abaisseras", "abaisserent", "abaisserez", "abaisseriez", "abaisserions", "abaisserons", "abaisseront", "abaisses", "abaisses", "abaisseur", "abaisseurs", "abaissez", "abaissiez", "abaissions", "abaissons", "abajoue", "abajoues", "abandon", "abandonna", "abandonnai", "abandonnaient", "abandonnais", "abandonnait", "abandonnames", "abandonnant", "abandonnas", "abandonnasse", "abandonnassent", "abandonnasses", "abandonnassiez", "abandonnassions", "abandonnat", "abandonnataire", "abandonnataires", "abandonnates", "abandonne", "abandonne", "abandonnee", "abandonnees", "abandonnent", "abandonner", "abandonnera", "abandonnerai", "abandonneraient", "abandonnerais", "abandonnerait", "abandonneras", "abandonnerent", "abandonnerez", "abandonneriez", "abandonnerions", "abandonnerons", "abandonneront", "abandonnes", "abandonnes", "abandonnez", "abandonniez", "abandonnions", "abandonnons", "abandons", "abaque", "abaques", "abasourdi", "abasourdie", "abasourdies", "abasourdimes", "abasourdir", "abasourdira", "abasourdirai", "abasourdiraient", "abasourdirais", "abasourdirait", "abasourdiras", "abasourdirent", "abasourdirez", "abasourdiriez", "abasourdirions", "abasourdirons", "abasourdiront", "abasourdis", "abasourdissaient", "abasourdissais", "abasourdissait", "abasourdissant", "abasourdissante", "abasourdissantes", "abasourdissants", "abasourdisse", "abasourdissement", "abasourdissements", "abasourdissent", "abasourdisses", "abasourdissez", "abasourdissiez", "abasourdissions", "abasourdissons", "abasourdit", "abasourdit", "abasourdites", "abat", "abatage", "abatages", "abatardi", "abatardie", "abatardies", "abatardimes", "abatardir", "abatardira", "abatardirai", "abatardiraient", "abatardirais", "abatardirait", "abatardiras", "abatardirent", "abatardirez", "abatardiriez", "abatardirions", "abatardirons", "abatardiront", "abatardis", "abatardissaient", "abatardissais", "abatardissait", "abatardissant", "abatardisse", "abatardissement", "abatardissements", "abatardissent", "abatardisses", "abatardissez", "abatardissiez", "abatardissions", "abatardissons", "abatardit", "abatardit", "abatardites", "abatee", "abatees", "abat-jour", "abats", "abattage", "abattages", "abattaient", "abattais", "abattait", "abattant", "abattante", "abattantes", "abattants", "abatte", "abattee", "abattees", "abattement", "abattements", "abattent", "abattes", "abatteur", "abatteurs", "abattez", "abattiez", "abattimes", "abattions", "abattirent", "abattis", "abattisse", "abattissent", "abattisses", "abattissiez", "abattissions", "abattit", "abattit", "abattites", "abattoir", "abattoirs", "abattons", "abattra", "abattrai", "abattraient", "abattrais", "abattrait", "abattras", "abattre", "abattrez", "abattriez", "abattrions", "abattrons", "abattront", "abattu", "abattue", "abattues", "abattus", "abat-vent", "abat-voix", "abbatial", "abbatiale", "abbatiales", "abbatiaux", "abbaye", "abbayes", "abbe", "abbes", "abbesse", "abbesses", "abbevillien", "abbevillienne", "abbevilliennes", "abbevilliens", "abces", "abdicataire", "abdicataires", "abdication", "abdications", "abdiqua", "abdiquai", "abdiquaient", "abdiquais", "abdiquait", "abdiquames", "abdiquant", "abdiquas", "abdiquasse", "abdiquassent", "abdiquasses", "abdiquassiez", "abdiquassions", "abdiquat", "abdiquates", "abdique", "abdique", "abdiquee", "abdiquees", "abdiquent", "abdiquer", "abdiquera", "abdiquerai", "abdiqueraient", "abdiquerais", "abdiquerait", "abdiqueras", "abdiquerent", "abdiquerez", "abdiqueriez", "abdiquerions", "abdiquerons", "abdiqueront", "abdiques", "abdiques", "abdiquez", "abdiquiez", "abdiquions", "abdiquons", "abdomen", "abdomens", "abdominal", "abdominale", "abdominales", "abdominaux", "abducteur", "abducteurs", "abduction", "abductions", "abecedaire", "abecedaires", "abeille", "abeilles", "aberra", "aberrai", "aberraient", "aberrais", "aberrait", "aberrames", "aberrance", "aberrances", "aberrant", "aberrante", "aberrantes", "aberrants", "aberras", "aberrasse", "aberrassent", "aberrasses", "aberrassiez", "aberrassions", "aberrat", "aberrates", "aberration", "aberrations", "aberre", "aberre", "aberrent", "aberrer", "aberrera", "aberrerai", "aberreraient", "aberrerais", "aberrerait", "aberreras", "aberrerent", "aberrerez", "aberreriez", "aberrerions", "aberrerons", "aberreront", "aberres", "aberrez", "aberriez", "aberrions", "aberrons", "abeti", "abetie", "abeties", "abetimes", "abetir", "abetira", "abetirai", "abetiraient", "abetirais", "abetirait", "abetiras", "abetirent", "abetirez", "abetiriez", "abetirions", "abetirons", "abetiront", "abetis", "abetissaient", "abetissais", "abetissait", "abetissant", "abetissante", "abetissantes", "abetissants", "abetisse", "abetissement", "abetissements", "abetissent", "abetisses", "abetissez", "abetissiez", "abetissions", "abetissons", "abetit", "abetit", "abetites", "abhorra", "abhorrai", "abhorraient", "abhorrais", "abhorrait", "abhorrames", "abhorrant", "abhorras", "abhorrasse", "abhorrassent", "abhorrasses", "abhorrassiez", "abhorrassions", "abhorrat", "abhorrates", "abhorre", "abhorre", "abhorree", "abhorrees", "abhorrent", "abhorrer", "abhorrera", "abhorrerai", "abhorreraient", "abhorrerais", "abhorrerait", "abhorreras", "abhorrerent", "abhorrerez", "abhorreriez", "abhorrerions", "abhorrerons", "abhorreront", "abhorres", "abhorres", "abhorrez", "abhorriez", "abhorrions", "abhorrons", "abima", "abimai", "abimaient", "abimais", "abimait", "abimames", "abimant", "abimas", "abimasse", "abimassent", "abimasses", "abimassiez", "abimassions", "abimat", "abimates", "abime", "abime", "abimee", "abimees", "abiment", "abimer", "abimera", "abimerai", "abimeraient", "abimerais", "abimerait", "abimeras", "abimerent", "abimerez", "abimeriez", "abimerions", "abimerons", "abimeront", "abimes", "abimes", "abimez", "abimiez", "abimions", "abimons", "abiotique", "abiotiques", "abject", "abjecte", "abjectement", "abjectes", "abjection", "abjections", "abjects", "abjura", "abjurai", "abjuraient", "abjurais", "abjurait", "abjurames", "abjurant", "abjuras", "abjurasse", "abjurassent", "abjurasses", "abjurassiez", "abjurassions", "abjurat", "abjurates", "abjuration", "abjurations", "abjure", "abjure", "abjuree", "abjurees", "abjurent", "abjurer", "abjurera", "abjurerai", "abjureraient", "abjurerais", "abjurerait", "abjureras", "abjurerent", "abjurerez", "abjureriez", "abjurerions"}

	for i := 0; i < len(words); i++ {
		_, err = stmt.Exec(i, words[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	return db
}
