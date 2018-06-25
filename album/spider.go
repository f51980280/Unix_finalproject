package main
import(
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var filenum int = 0
var repeated_time int = 0
var url string = ""
var host string = ""

var dir string = "picture"

func repeated(name string) bool {
	check_txt, err := os.OpenFile(dir+".txt", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("open file erro!")
	}
	defer check_txt.Close()
	b, err := ioutil.ReadFile(dir + ".txt")
	if err != nil {
		panic(err)
	}
	s := string(b)
	return strings.Contains(s, name)
}

func CrawlingImageTags(url string, imageChan chan []string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("open fail.")
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("open fail.")
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()

		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		reg, _ := regexp.Compile("<img .*>")
		imageTags := reg.FindAllString(resStr, -1)
		imageChan <- imageTags
	}
}
func getImageByUrl(url string) {
	if filenum < 1000 {
		reg, err := regexp.Compile("/.*jpg")
		if err != nil {
			fmt.Println("open fail.")
		}
		if reg.FindString(url) != "" {
			image := host + reg.FindString(url)
			imageRequest, err := http.Get(image)
			if err != nil {
				fmt.Println("open fail.")
			}
			data, err := ioutil.ReadAll(imageRequest.Body)
			defer imageRequest.Body.Close()
			if err != nil {
				fmt.Println("open fail.")
			}

			path := strings.Split(image, "/")
			var name string
			if len(path) > 1 {
				name = dir + "/" + path[len(path)-1]
				if repeated(name) == false {
					repeated_time = 0
					writeText(name)
//					name = fmt.Sprintf(dir+"/%03d.jpg", filenum)
					filenum++
					fmt.Println(name + " download success! ")
					os.MkdirAll(dir, os.ModePerm)
					out, err := os.Create(name)
					if err != nil {
						fmt.Println("open fail.")
					}
					io.Copy(out, bytes.NewReader(data))

				} else {
					repeated_time++
				}
			}

		}
	} else {
		fmt.Println("已下載完1000張圖片")
		var input string
		fmt.Scan(&input)

	}
}
func writeText(name string) {
	wr, err := os.OpenFile(dir+".txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}

	defer wr.Close()

	if _, err = wr.WriteString(name + "\n"); err != nil {
		panic(err)
	}

}

func getImageBySlice(imageSlice []string) {
	for _, value := range imageSlice {
		getImageByUrl(value)
	}
}
func main() {
	var i int = 0
	fmt.Println("請輸入0或1      (0)女優圖片網站   (1)imgur圖片網站")
	fmt.Scan(&i)
	if i == 0 {
		url = "http://www.ttpaihang.com/vote/rank.php?voteid=1089&page="
		host = "http://www.ttpaihang.com"
	} else {
		url = "https://imgur.com/search?q=face"
		host = "http:"
	}
	runtime.GOMAXPROCS(4)
	for {
//		if  filenum >= 1000 {
//			fmt.Println("Download complete!")
//			break
//		}
		imageChan := make(chan []string, 10)
		go CrawlingImageTags(url, imageChan)
		go getImageBySlice(<-imageChan)
	}
}

//程式碼範例：http://www.jianshu.com/p/6ab6e9727107
