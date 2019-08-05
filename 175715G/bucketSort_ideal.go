package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "time"
        "sync"
        "runtime"
  )

//バケットソートを行う関数
func bucketSort(numbers []int) ([]int){
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

//ソートしたいデータを2つに分割する関数
func splitData2(numbers []int, max int) ([]int,[]int){
	a := []int{}
	b := []int{}
	sp := max/2

	for i := 0; i<len(numbers); i++ {
		if numbers[i]<sp {
			a = append(a,numbers[i])
		//} else if numbers[i]>=sp && numbers[i]<sp*2{
		//	b = append(b,numbers[i])
		//} else if numbers[i]>=sp*2 && numbers[i]<sp*3{
		//	c = append(c,numbers[i])
		} else {
			b = append(b,numbers[i])
		}
	}
	
	return a,b
}

//二次元配列を一次元配列に直す
func mergeSlise(matrix [][]int) []int{
	resultData := []int{}
	for i:=0; i<len(matrix); i++{
		resultData = append(resultData,matrix[i]...)
	}
	return resultData
}

//ソートするデータのファイルを読み込む関数
func fromFile(filePath string) []string {
        f, err := os.Open(filePath)
        if err != nil {
          fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
          os.Exit(1)
        }

        defer f.Close()

        lines := make([]string,0)
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
          lines = append(lines, scanner.Text())
        }
        if serr := scanner.Err(); serr != nil {
          fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
        }

        return lines
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

//ソートしたデータを"result.txt"に書き込みする関数
func writefile(result []int){
        fp, err := os.Create("resultTest.txt")
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

func main() {
	lines := fromFile("rannum_1000000.txt")
	data := stringToint(lines)
	max := 0

	cpus := runtime.NumCPU()	//使用しているCPUの数
	wg := sync.WaitGroup{}
	ch := make(chan int, cpus)	//Goroutineの理想個数

	//ソートしたい配列内の最大値を決める
        for i := 0; i<len(data); i++ {
                if max < data[i] {max = data[i]}
        }

	a,b := splitData2(data,max)
	matrix := [][]int{a,b}

	//並列化の計測
	start := time.Now()
	for i := 0; i<len(matrix); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- 1
			matrix[i] = bucketSort(matrix[i])
			<-ch
		}(i)
	}
	wg.Wait()
	end := time.Now()
	
	sortdata := mergeSlise(matrix)
	writefile(sortdata)
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}
