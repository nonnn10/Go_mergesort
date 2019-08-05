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
  "log"
  "runtime"
)

func main() {
  lines := fromFile("./random_data/ransu_100.txt")  //file読み込みlinesに代入
  int_lines := stringToint(lines)   //読み込んだstring型をint型に変換しint_linesに代入
  result := make(chan []int)        //マージソートの結果を送受信するchannel型resultを作成
  
  start := time.Now()               //マージソートにかかる時間の計測開始
   go MergeSort(int_lines, result)  //並列でマージソートの実行
  end := time.Now()                 //時間の計測終了
  r := <- result                    //resultを受信しrに代入

  filewright(r)                     //rに入ったソート後のデータを”result.txt”fileに書き込み
    close(result)                   //channelのクローズ
fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())     //計測時間の標準出力
log.Println(runtime.NumGoroutine())                 //なくても良い(goroutine数の確認)
}


//fileを一行ずつ読み込む関数
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

//マージを行う関数
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


//マージソートを行う関数
func MergeSort(data []int, r chan []int){
    if len(data) == 1 {                     //データが1個になったら分割終了
        r <- data                           //データをrに送信
        return 
    }
    
    leftChan := make(chan []int)            //left側のデータを送受信するchannelの作成
    rightChan := make(chan []int)           //right側のデータを送受信するchannelの作成
    middle := len(data)/2                   //データを２分割した数を変数に代入
    
	go MergeSort(data[:middle], leftChan)   //left側のデータを並列にマージソート
    go MergeSort(data[middle:], rightChan)  //right側のデータを並列にマージソート
   // log.Println(runtime.NumGoroutine())
    ldata := <-leftChan                     //マージソートされたleft側のデータを受信しldata変数に代入
    rdata := <-rightChan                    //マージソートされたright側のデータを受信しrdata変数に代入

    close(leftChan)                         //channelのクローズ(クローズしなくても良いがリソースの解放)
    close(rightChan)                        //channelのクローズ
    r <- Merge(ldata, rdata)                //マージした結果をchannelに送信

    return 
}




//スライスの型変換(string>int)する関数
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

//fileに書き込む関数
func filewright(r []int){
    fp, err := os.Create("result.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer fp.Close()
    for _,v := range r {
        fp.WriteString(strconv.Itoa(v)+"\n")
        //fp.WriteString("\n")
    }
}