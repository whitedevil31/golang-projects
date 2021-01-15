package main

import ("fmt"
"github.com/jung-kurt/gofpdf"
"io/ioutil"

)


func main(){
	
	content,err := ioutil.ReadFile("C:\\Users\\NANDA\\Desktop\\test.docx")
	if err !=nil {
	fmt.Println(err)
	}
	pdf := gofpdf.New("P","mm", "A4","")
	pdf.AddPage()
	pdf.SetFont("Arial","B",14)
	pdf.MultiCell(190,5,string(content),"0","0",false)
	_=pdf.OutputFileAndClose("test.pdf")
	fmt.Println("process completed")
}