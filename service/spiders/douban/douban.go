package douban

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gonb/pkg/db"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var BaseUrl = "https://movie.douban.com/top250"

type DoubanMovie struct {
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

type Page struct {
	Page int
	Url  string
}

// 获取分页
func GetPages(url string) []DoubanMovie {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	return ParseMovies(doc)
}

func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{Page: 1, Url: ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")

		pages = append(pages, Page{
			Page: page,
			Url:  url,
		})
	})

	return pages
}

func ParseMovies(doc *goquery.Document) (movies []DoubanMovie) {
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find(".hd a span").Eq(0).Text()

		subTitle := selection.Find(".hd a span").Eq(1).Text()
		subTitle = strings.TrimLeft(subTitle, "  / ")

		other := selection.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(other, "  / ")

		desc := strings.TrimSpace(selection.Find(".bd p").Eq(0).Text())
		DescInfo := strings.Split(desc, "\n")
		desc = DescInfo[0]

		movieDesc := strings.Split(DescInfo[1], "/")
		year := strings.TrimSpace(movieDesc[0])
		area := strings.TrimSpace(movieDesc[1])
		tag := strings.TrimSpace(movieDesc[2])

		star := selection.Find(".bd .star .rating_num").Text()

		comment := strings.TrimSpace(selection.Find(".bd .star span").Eq(3).Text())
		compile := regexp.MustCompile("[0-9]")
		comment = strings.Join(compile.FindAllString(comment, -1), "")

		quote := selection.Find(".quote .inq").Text()

		movie := DoubanMovie{
			Title:    title,
			Subtitle: subTitle,
			Other:    other,
			Desc:     desc,
			Year:     year,
			Area:     area,
			Tag:      tag,
			Star:     star,
			Comment:  comment,
			Quote:    quote,
		}

		//log.Printf("i: %d, movie: %v", i, movie)
		movies = append(movies, movie)
	})

	return movies
}

func Start() {
	var movies []DoubanMovie
	req, err := http.NewRequest("GET", "http://movie.douban.com/top250", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")
	resp, err := (&http.Client{}).Do(req)

	//pages := GetPages(BaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)


	//for _, page := range pages {
	if err != nil {
		log.Println(err)
	}

	movies = append(movies, ParseMovies(doc)...)
	//}

	fmt.Println(movies)
}

func Add(movies []DoubanMovie) {
	for index, movie := range movies {
		if err := db.GormDb.Table("sp_douban_movie").Create(&movie).Error; err != nil {
			fmt.Printf("db.Create index: %s, err : %v", index, err)
		}
	}
}
