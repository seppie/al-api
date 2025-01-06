package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	//"image/gif"
)

type GifsInfo struct {
	Data struct {
		URL string `json:"embed_url"`
	} `json:"data"`
}

type OrderInfo struct {
	Setting struct {
		Title string `json:"overrideTitle"`
		Brief string `json:"overrideBrief"`
	} `json:"setting"`
}

type DogeInfo struct {
	Message string `json:"message"`
}

type ChuckInfo struct {
	IconURL string `json:"icon_url"`
	URL     string `json:"url"`
	Value   string `json:"value"`
}

func Products(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func Chuck(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, "https://api.chucknorris.io/jokes/random", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("there was an error")
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	info := ChuckInfo{}
	json.Unmarshal(b, &info)
	w.Write([]byte(info.Value))
	w.Write([]byte("\n"))
}

func DogeCoin(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, "https://dog.ceo/api/breeds/image/random", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("there was an error")
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	info := DogeInfo{}
	json.Unmarshal(b, &info)
	imageresponse, _ := http.Get(info.Message)
	image, _ := io.ReadAll(imageresponse.Body)
	w.Write(image)
}

func Orders(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, "https://api.helldivers2.dev/raw/api/v2/Assignment/War/801", nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-Super-Client", "Democracy")
	req.Header.Add("X-Super-Contact", "Super-Al")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("there was an error")
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	info := []OrderInfo{}
	json.Unmarshal(b, &info)
	s := fmt.Sprintf("%s\nBrief: %s", info[0].Setting.Title, info[0].Setting.Brief)
	w.Write([]byte(s))
}

func GIF(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Not implemented yet. Coming soon!"))
	req, _ := http.NewRequest(http.MethodGet, "https://api.giphy.com/v1/gifs/random?api_key=Tl4gvuLYN12TErEvUNrWbI0NaBUuxuju&tag=&rating=g", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("there was an error")
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	//w.Write(b)
	info := GifsInfo{}
	json.Unmarshal(b, &info)
	//imageresponse, _ := http.Get(info.Data.URL)
	//image, _ := io.ReadAll(imageresponse.Body)
	//http.Redirect(w, r, info.Data.URL, http.StatusFound)
	//w.Write(image)
	/*	//w.Write(b)
			info := GifsInfo{}
			json.Unmarshal(b, &info)
			//http.Redirect(w, r, info.Data.URL, http.StatusFound)
			imageresponse, _ := http.Get(info.Data.URL)
			//image, _ := io.ReadAll(imageresponse.Body)
			image := bytes.NewReader(imageresponse.Body)
			gifImg, err := gif.DecodeAll(r)
		        if err != nil {
		            http.Error(w, "Error decoding GIF", http.StatusInternalServerError)
		            return
		        }
			w.Header().Set("Content-Type", "image/gif")
			gif.EncodeAll(w, gifImg)
			w.Write(image)
	*/
}

func Help(w http.ResponseWriter, r *http.Request) {
	s := "/help\t\tList all endpoints\n" +
		"/hello\t\tHello world\n" +
		"/chuck\t\tChuck Norris Fact\n" +
		"/doge\t\tDog pic\n" +
		"/liberty\tCurrent Helldivers II Brief\n" +
		"/gif\t\tDisplay a random GIF"
	w.Write([]byte(s))
}

func main() {
	log.Println("starting program")
	http.HandleFunc("/help", Help)
	http.HandleFunc("/hello", Products)
	http.HandleFunc("/chuck", Chuck)
	http.HandleFunc("/doge", DogeCoin)
	http.HandleFunc("/liberty", Orders)
	http.HandleFunc("/gif", GIF)
	http.ListenAndServe("192.168.1.4:8080", nil)
}
