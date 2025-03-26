package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"unicode"

	"github.com/gocolly/colly/v2"
)

type Data struct {
	Texts []string
	Ja    string
	En    string
	Page  int
}

type Bible struct {
	ja   string
	en   string
	q    string
	page int
}

var dataArray = []Bible{
	Bible{ja: "創世記", en: "Genesis", q: "GEN", page: 50},
	Bible{ja: "出エジプト記", en: "Exodus", q: "EXO", page: 40},
	Bible{ja: "レビ記", en: "Leviticus", q: "LEV", page: 27},
	Bible{ja: "民数記", en: "Numbers", q: "NUM", page: 36},
	Bible{ja: "申命記", en: "Deuteronomy", q: "DEU", page: 34},
	Bible{ja: "ヨシュア記", en: "Joshua", q: "JOS", page: 24},
	Bible{ja: "士師記", en: "Judges", q: "JDG", page: 21},
	Bible{ja: "ルツ記", en: "Ruth", q: "RUT", page: 4},
	Bible{ja: "サムエル記(上)", en: "Samuel(1)", q: "1SA", page: 31},
	Bible{ja: "サムエル記(下)", en: "Samuel(2)", q: "2SA", page: 24},
	Bible{ja: "列王記(上)", en: "Kings(1)", q: "1KI", page: 22},
	Bible{ja: "列王記(下)", en: "Kings(2)", q: "2KI", page: 25},
	Bible{ja: "歴代誌(上)", en: "Chronicles(1)", q: "1CH", page: 29},
	Bible{ja: "歴代誌(下)", en: "Chronicles(2)", q: "2CH", page: 36},
	Bible{ja: "エズラ記", en: "Ezra", q: "EZR", page: 10},
	Bible{ja: "ネヘミヤ記", en: "Nehemiah", q: "NEH", page: 13},
	Bible{ja: "エステル記", en: "Esther", q: "EST", page: 10},
	Bible{ja: "ヨブ記", en: "Job", q: "JOB", page: 42},
	Bible{ja: "詩編", en: "Psalms", q: "PSA", page: 150},
	Bible{ja: "箴言", en: "Proverbs", q: "PRO", page: 31},
	Bible{ja: "コヘレトの言葉", en: "Ecclesiastes", q: "ECC", page: 12},
	Bible{ja: "雅歌", en: "Song-of-Songs", q: "SNG", page: 8},
	Bible{ja: "イザヤ書", en: "Isaiah", q: "ISA", page: 66},
	Bible{ja: "エレミヤ書", en: "Jeremiah", q: "JER", page: 52},
	Bible{ja: "哀歌", en: "Lamentation", q: "LAM", page: 5},
	Bible{ja: "エゼキエル書", en: "Ezekiel", q: "EZK", page: 48},
	Bible{ja: "ダニエル書", en: "Daniel", q: "DAN", page: 12},
	Bible{ja: "ホセア書", en: "Hosea", q: "HOS", page: 14},
	Bible{ja: "ヨエル書", en: "Joel", q: "JOL", page: 4},
	Bible{ja: "アモス書", en: "Amos", q: "AMO", page: 9},
	Bible{ja: "オバデヤ書", en: "Obadiah", q: "OBA", page: 1},
	Bible{ja: "ヨナ書", en: "Jonah", q: "JON", page: 4},
	Bible{ja: "ミカ書", en: "Micah", q: "MIC", page: 7},
	Bible{ja: "ナホム書", en: "Nahum", q: "NAM", page: 3},
	Bible{ja: "ハバクク書", en: "Habakkuk", q: "HAB", page: 3},
	Bible{ja: "ゼファニヤ書", en: "Zephaniah", q: "ZEP", page: 3},
	Bible{ja: "ハガイ書", en: "Haggai", q: "HAG", page: 2},
	Bible{ja: "ゼカリヤ書", en: "Zechariah", q: "ZEC", page: 14},
	Bible{ja: "マラキ書", en: "Malachi", q: "MAL", page: 3},
	Bible{ja: "マタイによる福音書", en: "Matthew", q: "MAT", page: 28},
	Bible{ja: "マルコによる福音書", en: "Mark", q: "MRK", page: 16},
	Bible{ja: "ルカによる福音書", en: "Luke", q: "LUK", page: 24},
	Bible{ja: "ヨハネによる福音書", en: "John", q: "JHN", page: 21},
	Bible{ja: "使徒言行録", en: "Acts", q: "ACT", page: 28},
	Bible{ja: "ローマの信徒への手紙", en: "Romans", q: "ROM", page: 16},
	Bible{ja: "コリントの信徒への手紙(1)", en: "Corinthians(1)", q: "1CO", page: 16},
	Bible{ja: "コリントの信徒への手紙(2)", en: "Corinthians(2)", q: "2CO", page: 13},
	Bible{ja: "ガラテヤの信徒への手紙", en: "Galatians", q: "GAL", page: 6},
	Bible{ja: "エフェソの信徒への手紙", en: "Ephesians", q: "EPH", page: 6},
	Bible{ja: "フィリピの信徒への手紙", en: "Philippians", q: "PHP", page: 4},
	Bible{ja: "コロサイの信徒への手紙", en: "Colossians", q: "COL", page: 4},
	Bible{ja: "テサロニケの信徒への手紙(1)", en: "Thessalonians(1)", q: "1TH", page: 5},
	Bible{ja: "テサロニケの信徒への手紙(2)", en: "Thessalonians(2)", q: "2TH", page: 3},
	Bible{ja: "テモテへの手紙(1)", en: "Timothy(1)", q: "1TI", page: 6},
	Bible{ja: "テモテへの手紙(2)", en: "Timothy(2)", q: "2TI", page: 4},
	Bible{ja: "テトスへの手紙", en: "Titus", q: "TIT", page: 3},
	Bible{ja: "フィレモンへの手紙", en: "Philemon", q: "PHM", page: 1},
	Bible{ja: "ペトロの手紙(1)", en: "Peter(1)", q: "1PE", page: 5},
	Bible{ja: "ペトロの手紙(2)", en: "Peter(2)", q: "2PE", page: 3},
	Bible{ja: "ヨハネの手紙(1)", en: "John(1)", q: "1JN", page: 5},
	Bible{ja: "ヨハネの手紙(2)", en: "John(2)", q: "2JN", page: 1},
	Bible{ja: "ヨハネの手紙(3)", en: "John(3)", q: "3JN", page: 1},
	Bible{ja: "ヘブライ人への手紙", en: "Hebrews", q: "HEB", page: 13},
	Bible{ja: "ヤコブの手紙", en: "James", q: "JAS", page: 5},
	Bible{ja: "ユダの手紙", en: "Jude", q: "JUD", page: 1},
	Bible{ja: "ヨハネの黙示録", en: "Revelation", q: "REV", page: 22},
}

const BASE_URL = "https://www.bible.com/ja/bible/1819"

var tpl *template.Template

func main() {

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func init() {
	tpl = template.Must(template.ParseFiles("web/index.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	idx := rand.Intn(len(dataArray))
	q := dataArray[idx].q
	page := rand.Intn(dataArray[idx].page) + 1

	var url = fmt.Sprintf("%s/%s.%d", BASE_URL, q, page)

	c := colly.NewCollector(
		colly.AllowedDomains("www.bible.com"),
	)

	d := Data{Ja: dataArray[idx].ja, En: dataArray[idx].en, Page: page}

	c.OnHTML("div[data-usfm]", func(e *colly.HTMLElement) {
		var num string
		var text string
		currentCount := -1
		e.ForEach("span[data-usfm]", func(_ int, el *colly.HTMLElement) {
			i := countDigit(el.Text)

			if i == 0 {
				d.Texts[currentCount] = d.Texts[currentCount] + "<br>" + el.Text
			} else {
				num = el.Text[0:i]
				text = el.Text[i:]

				d.Texts = append(d.Texts, fmt.Sprintf("<span class=\"number\">%s</span>%s", num, text))
				currentCount += 1
			}
		})
	})

	c.Visit(url)
	tpl.Execute(w, d)
}

func countDigit(s string) int {
	ss := []rune(s)
	i := 0
	for {
		if unicode.IsDigit(ss[i]) {
			i += 1
		} else {
			break
		}
	}
	return i
}
