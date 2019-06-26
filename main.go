package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/tabwriter"
)

type LokerModel struct {
	no     int
	typeID string
	noID   string
}

var loker []LokerModel

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "init":
		if len(arrCommandStr) < 2 {
			return errors.New("Silahkan masukkan jumlah loker")
		} else {
			//arrLength := arrCommandStr[1]
			i, _ := strconv.Atoi(arrCommandStr[1])
			temp := make([]LokerModel, i)
			loker = temp
			fmt.Println("Berhasil membuat loker dengan jumlah", len(loker))
		}
		return nil

	case "status":
		w := new(tabwriter.Writer)

		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)

		defer w.Flush()

		fmt.Fprintf(w, "\n %s\t%s\t%s\t\n", "No Loker", "Tipe Identitas", "No Identitas")
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", "----", "----", "----")

		for i := 0; i < len(loker); i++ {
			fmt.Fprintf(w, "%d\t%s\t%s\t\n", i+1, loker[i].typeID, loker[i].noID)
		}
		return nil

	case "input":
		if len(arrCommandStr) < 3 {
			return errors.New("Silahkan masukkan tipe ID dan No. ID")
		}
		found := false
		for i := range loker {
			if loker[i].typeID == "" {
				loker[i] = LokerModel{i + 1, arrCommandStr[1], arrCommandStr[2]}
				fmt.Println("Kartu identitas tersimpan di loker nomor ", i+1)
				found = true
				break
			}
		}
		if !found {
			return errors.New("Maaf loker sudah penuh")
		}
		return nil

	case "leave":
		if len(arrCommandStr) < 2 {
			return errors.New("Silahkan masukkan nomor loker")
		}
		i, _ := strconv.Atoi(arrCommandStr[1])
		index := i - 1
		loker[index] = LokerModel{}
		fmt.Println("Loker nomor " + arrCommandStr[1] + " berhasil dikosongkan")
		return nil

	case "find":
		if len(arrCommandStr) < 2 {
			return errors.New("Silahkan masukkan nomor identitas")
		}
		found := false
		for i := range loker {
			if loker[i].noID == arrCommandStr[1] {
				fmt.Println("Kartu identitas tersebut berada di loker nomor ", loker[i].no)
				found = true
				break
			}
		}
		if !found {
			return errors.New("Nomor identitas tidak ditemukan")
		}
		return nil

	case "search":
		if len(arrCommandStr) < 2 {
			return errors.New("Silahkan masukkan tipe identitas")
		}
		for i := range loker {
			if loker[i].typeID == arrCommandStr[1] {
				fmt.Print(loker[i].noID, " ")
			}
		}
		fmt.Println()
		return nil
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
