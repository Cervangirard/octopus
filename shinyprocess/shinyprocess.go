package shinyprocess

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ShinyProcess struct {
	Who string
	P   *os.Process // Not sure of this
	Url string
}

// Launch R process with shinyapp inside
func (p *ShinyProcess) LaunchApp(r *http.Request) (UrlShiny string) {

	fmt.Println("launch process -----")

	// Second try

	// /!\ NOT USED FOR NOW !!!! ---------------

	// Add array with random port

	// range_port := make([]int, 2000)
	// for i := 0; i < 2000; i++ {
	// 	range_port[i] = 30000 + i
	// }

	// // Set  random port
	// rand.Seed(time.Now().UnixNano())

	// random_port := rand.Intn(len(range_port))

	// port := range_port[random_port]

	// fmt.Println(port)
	// ---------------------

	host_shiny := fmt.Sprintf("http://localhost:%v", 8000)

	// Launch shinyapp Process

	cmdToRun := "/usr/local/bin/Rscript"

	args := []string{"/usr/local/bin/Rscript", "shiny_app.R", strconv.Itoa(8000)}

	procAttr := new(os.ProcAttr)

	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

	fmt.Println(host_shiny)

	// Get user cookie to know who is who
	user_cookie, err_c := r.Cookie("user")

	if err_c != nil {
		log.Fatal(err_c)
	}

	var err_process error

	// Set ShinyProcess

	p.P, err_process = os.StartProcess(cmdToRun, args, procAttr)

	if err_process != nil {
		log.Fatal(err_process)
	}

	p.Who = user_cookie.Value

	p.Url = host_shiny

	return host_shiny
}

// Kill ShinyProcess
func (p *ShinyProcess) KillSession() {
	p.P.Kill()

	fmt.Printf("Process killed !!!")
}
