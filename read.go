// Thi is the text file format the given text file name is :name.txt
// rahul vemuri
// aparna vemuri
// akanksha upasani
// srushti shahi

package main
import ( 
	"fmt"
	"strings"
	"os"
	"io/ioutil"
)
type name struct{
	fname string
	lname string
}
func main(){
	
	var a name
	
	sli:=make([]map[string]string,0)

	f, err := os.Open("name.txt")
	fmt.Println(err)
	
	rawBytes, err := ioutil.ReadAll(f)
	
	fmt.Println(err)
	lines := strings.Split(string(rawBytes), "\n")
	for  i,line := range lines{
		k:=0
		fmt.Printf("%d",i)
		for j,c:= range line{
			fmt.Printf("%d",j)
			if string(c)==" "{
				k=1
				continue
			}
			fmt.Println(k)
			if string(c)!=" " && k!=1{
				a.fname=a.fname+string(c)
			}
			if string(c)!=" " && k==1{
				
				a.lname=a.lname+string(c)
				
				
			}
			
			
			
		}
		names:=map[string]string{
			"first_name":a.fname,
			"last_name":a.lname}
		sli=append(sli,names)
		
		a.fname=" "
		a.lname=" "
		
	}
	for i,line:=range sli{
		fmt.Println(i,line)
	}

	
}