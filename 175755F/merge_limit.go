/*
このプログラムでは並列化におけるgoroutineの数を制限するようにしている
結果は並列化していないマージソートよりも遅くなるという結果だった
*/


package main

//https://gist.github.com/julianshen/3940045#file-merge-go-L9
//https://maku77.github.io/hugo/go/file.html
//https://stackoverflow.com/questions/24972950/go-convert-strings-in-array-to-integer
//参考にしたURLgoroutine数の制限
//https://qiita.com/woogy/items/b5cab817a64e071b48f4

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "time"
  "log"
  "runtime"
  //"sync"
)

func main() {
    //runtime.GOMAXPROCS(2)
    lines := fromFile("./random_data/ransu_100.txt")
    int_lines := stringToint(lines)
    start := time.Now()
       MergeSort(int_lines)
        //fmt.Println(MergeSort(int_lines))
    end := time.Now()
    log.Println(runtime.NumGoroutine())
    //wrightfile(r)
    fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
    //fmt.Println(runtime.GOMAXPROCS(8))
}

var (
    //GOMAXPROCS = runtime.GOMAXPROCS(2)
    ch       = make(chan bool, 2)      //goroutineが同時に並列できる数(グローバル変数)
    count int = 0 
)





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

    return
}

func MergeSort(data []int) (r []int){
   //log.Println(runtime.NumGoroutine()) 
   //fmt.Printf("len<cap = %d<%d\n", len(ch), cap(ch))
   if len(ch) < cap(ch)-2 {                           //chの容量以下なら並列化
        if len(data) == 1 {
            return data
        }
        //log.Println(runtime.NumGoroutine())         //並列化のgoroutine数のprint
        ch <- true
        ch <- true
        leftChan := make(chan []int)
        rightChan := make(chan []int)
        middle := len(data)/2
        go func() { leftChan <- MergeSort(data[:middle]); <-ch }()
        go func() { rightChan <- MergeSort(data[middle:]); <-ch }()
        l_data := <-leftChan    //変数名をelse部分と違うものにした
        r_data := <-rightChan
        close(leftChan)
        close(rightChan)
        r = Merge(l_data, r_data)
        return r
} else {                                             //chの容量以上なら並列化なし
    if len(data) == 1 {
        return data
    } 
    //count++
    //fmt.Println(count)
    //log.Println(runtime.NumGoroutine())
    middle := len(data)/2
    ldata := MergeSort(data[:middle])
    rdata := MergeSort(data[middle:])
    r := Merge(ldata, rdata)
    return r
}
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

//ソートしたデータを"result.txt"に書き込みする関数
func wrightfile(result []int){
	fp, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
  
	defer fp.Close()
	for _,v := range result {
		fp.WriteString(strconv.Itoa(v)+"\n")
		//fp.WriteString("\n")
	}
}

//引数のint型配列をfileに書き出す
func wrightfile_go(result []string){
	fp, err := os.Create("goroutine.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
  
	defer fp.Close()
	for _,v := range result {
		fp.WriteString(v+"\n")
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