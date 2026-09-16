package advanced

import (
	"fmt"
	"os/exec"
)

func os_process() {

	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))

	// pr, pw := io.Pipe()

	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = pr

	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("food is good\nbar\nbaz\n"))
	// }()

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Output:", string(output))

	// cmd := exec.Command("printenv", "SHELL")

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Output:", string(output))

	// cmd := exec.Command("sleep", "60")

	// // Start the command
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error staring command:", err)
	// 	return
	// }

	// time.Sleep(2 * time.Second)

	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println("Error killing process:", err)
	// 	return
	// }
	// fmt.Println("Process killed")

	// // Waiting
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("Error waiting:", err)
	// 	return
	// }
	// fmt.Println("Process is complete")

	// cmd := exec.Command("grep", "foo")
	// // Set input for the command
	// cmd.Stdin = strings.NewReader("food is good\nbar\nbaz\n")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

	// cmd := exec.Command("echo", "Hello World!")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))
}
