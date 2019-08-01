package main

//https://gist.github.com/julianshen/3940045#file-merge-go-L9
//https://maku77.github.io/hugo/go/file.html
//https://stackoverflow.com/questions/24972950/go-convert-strings-in-array-to-integer

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "time"
  
)

func main() {
  lines := fromFile("../ransu_100.txt")
  int_lines := stringToint(lines)
  //result := make([]int,0)

  start := time.Now()
  MergeSort(int_lines)//ar := MergeSort(int_lines,result)
  //fmt.Println(MergeSort(int_lines))
  end := time.Now()
  //wrightfile(ar)
fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}



//マージソート
func Merge(ldata []int, rdata []int) (result []int) {
    result = make([]int, len(ldata) + len(rdata))
    lidx, ridx := 0, 0

    for i:=0;i<cap(result);i++ {
        switch {
            case lidx >= len(ldata):
                result[i] = rdata[ridx]
                ridx++
            case ridx >= len(rdata):
                result[i] = ldata[lidx]
                lidx++
            case ldata[lidx] < rdata[ridx]:
                result[i] = ldata[lidx]
                lidx++
            default:
                result[i] = rdata[ridx]
                ridx++
        }
    }

    return result
}

func MergeSort(data []int) []int {
    if len(data) == 1 {
        return data
    } 
     //leftChan := make([]int,0)
     //rightChan := make([]int,0)
     middle := len(data)/2
     ldata := MergeSort(data[:middle])
     rdata := MergeSort(data[middle:])
     r := Merge(ldata, rdata)
     return r
    
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
func wrightfile(ar []int){
    fp, err := os.Create("result.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer fp.Close()

    fmt.Println(len(ar))
    var count3 int
    for _,v := range ar {
        fp.WriteString(strconv.Itoa(v)+"\n")
        count3++
        //fp.WriteString("\n")
    }
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