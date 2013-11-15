package main 

import ("os";"bufio";"os/exec";"fmt")

func main(){
   	cmd := exec.Command("/bin/ls","-l")
    ot,_ := cmd.Output()
	os.Stdout.Write(ot[:])
 	

    fmt.Printf("\n\n\n")

	buf := make([]byte,1024)
	f,_ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
       n,_ := r.Read(buf)
       if n == 0 { break }
       w.Write(buf[:n])
	}

}