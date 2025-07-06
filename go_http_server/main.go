package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	i := 1
	if len(args) > 1 {
		n, err := strconv.Atoi(args[1])
		if err != nil {
			n = 1
		}
		i = n
	}
	if i == 1 {
		server1()
	} else if i == 2 {
		server2()
	} else if i == 3 {
		server3()
	} else if i == 4 {
		server4()
	} else if i == 5 {
		server5()
	}
}

/**
 * Server1: 単一のハンドラ
 */
type Server1Handler struct{}

// Server1Handler型はhttp.Handlerインタフェースを実装している
// Handler…ServeHTTP(http.ResponseWriter, *http.Request)メソッドを持ったインタフェース
func (h Server1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server 1")
}

func server1() {
	// 空の構造体Server1Handlerのインスタンスを生成
	handler := Server1Handler{}

	server := &http.Server{
		Addr:    ":8080",
		Handler: &handler, // Handler型のインスタンスを引数として渡す
	}
	server.ListenAndServe()
}

/**
 * Server2: 複数のハンドラ
 */
type Server2aHandler struct{}
type Server2bHandler struct{}

func (h Server2aHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server 2a")
}
func (h Server2bHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server 2b")
}

func server2() {
	s2a := Server2aHandler{}
	s2b := Server2bHandler{}

	// それぞれのルートURLとハンドラをDefaultServeMux(*1)に紐付ける
	// (*1) net/httpパッケージ内でインスタンス化されているServeMux
	http.Handle("/s2a", s2a)
	http.Handle("/s2b", s2b)

	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}

/**
 * Server3: 複数のハンドラ関数(ServeMux)
 */
func server3() {
	// ServeMuxを生成
	mux := http.NewServeMux()

	// mux.HandleFunc()でルートURLをHandler関数(*1)と紐付ける
	// (*1) Handler関数…http.ServeHTTP()と同じシグネチャを持つ関数・メソッド
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server3 Document Root!\n"))
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server3 hello!\n"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

/**
 * Server4: 複数のハンドラ関数(net/http)
 */
func server4() {
	// http.HandleFunc()でルートURLをHandler関数(*1)と紐付ける
	// (*1) Handler関数…http.ServeHTTP()と同じシグネチャを持つ関数・メソッド
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server4 Document Root!\n"))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server4 hello!\n"))
	})

	server := &http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}


/**
 * Server5: 複数のハンドラ(http.HandlerFunc)
 */
// 
// 
func DocumentRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server 5 Document Root")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server 5 Hello")
}

func server5() {
	server := &http.Server{
		Addr:    ":8080",
	}
	http.HandleFunc("/", http.HandlerFunc(DocumentRoot))
	http.HandleFunc("/hello", http.HandlerFunc(Hello))
	server.ListenAndServe()
}
