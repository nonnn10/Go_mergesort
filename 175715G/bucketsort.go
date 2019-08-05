package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	//"sync"
	//"log"
  )
 
func bucketSort(numbers []int) []int{
	max := 0
	for i := 0; i<len(numbers); i++ {
		if max < numbers[i] {max = numbers[i]}
	}
	bucket := make([]int , max+1)
	for i := 0; i<len(numbers); i++ {
		bucket[numbers[i]] += 1
	}
	arr := []int{}
	for i:=0; i<len(bucket); i++{
		for j:=0; j<bucket[i]; j++{
			arr = append(arr , i)
		}
	}
	return arr
}

func fromFile(filePath string) []string {
	// ファイルを開く
	f, err := os.Open(filePath)
	if err != nil {
	  fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
	  os.Exit(1)
	}
  
	// 関数return時に閉じる
	defer f.Close()
  
	// Scannerで読み込む
	lines := make([]string,0)//, 0) 3000000)  // ある程度行数が事前に見積もれるようであれば、makeで初期capacityを指定して予めメモリを確保しておくことが望ましい
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
	  // appendで追加
	  lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
	  fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
	}
  
	return lines
}
 
//スライスの型変換(string>int)
func stringToint(t []string) []int {
	t2 := make([]int, 0) //= []int{}
	for _, i := range t { 
		j, err := strconv.Atoi(i) 
		if err != nil {
			panic(err) 
		} 
			t2 = append(t2, j) 
	} 
	 //fmt.Println(t2) 
	return t2
}

/*
func main(){
	numbers := []int{1,1,3,2,4,6,1,2,4,6}
	nums := bucketSort(numbers)
	fmt.Println(nums)
}
*/

func main() {
	lines := fromFile("../ransu_10000000.txt")
	int_lines := stringToint(lines)
	//result := make(chan []int)
  
	start := time.Now()
	bucketSort(int_lines)
	end := time.Now()
	//r := <- result
  
	/*
	fp, err := os.Create("result_1000000.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
  
	defer fp.Close()
	for _,v := range nums {
		fp.WriteString(strconv.Itoa(v)+"\n")
		//fp.WriteString("\n")
	}
	*/
	//close(result)
  fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
  }
