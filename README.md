並列分散処理
================
Overview  
Golang、Pythonn、C++におけるマージソート、バケットソートの並列化を行なったもの

##Description  
それぞれ班員が作成したプログラムを班員別に学籍番号でまとめた  
下記にファイルの内容をまとめておく  
ソートするデータはrandom_dataファイルに入っている  

##175702E  
    MergeSort_Python.ipynbがPythonでマージソート を実装したもの
    実行すると、逐次処理・並列処理・sortedでの処理時間をそれぞれ出力する
    numprocでプロセス数を指定している
    result_merge.xlsxに結果を記載している
    使用したデータ配列は、別フォルダのらrandom_dataを使用している

##175715G  
　bucketSort.goファイルがGo言語でバケットソートを実装したもの
　bucketSort_parallel.goファイルはGo言語でのバケットソートの並列化を実装したもの
　関数SplitDataの数字を任意の数に変更することで非ソートデータの分割数を調整することができる
　bucketSort_Changeファイルはバケットソートの並列化を実装する際に作成されるGoroutineの数を
任意の数に設定できるようにしたもの
　変数chの配列数の変更により,Goroutineの数を調整することができる
　bucketSort.idealファイルは処理に使用しているCPU数に合わせて理想的なGoroutineの数を作成し,
実行してくれるもの
　
##175727A  
  merge-parallel.cppファイルはマージソートをC++で実装したものである.
  NUM_ITEMSでソートする要素数, NUM_LIMITで1要素の最大数を指定することができ、
  実行すると計測時間を出力する. また, #pragma omp parallel sections num_threads()の引数で
  スレッド数を指定することができる.
  merge_parallel_malloc.cppファイルは, merge-parallel.cppでNUM_ITEMSに1000万を指定した際に
  segmentation faultが発生する問題をmalloc()を用いて必要なメモリを確保することで改善したものである.

##175743B  

##175755F  
merge.goファイルがマージソートの並列化を実践したもの  
merge_limitファイルがマージソートにおける同時に実行できるgoroutine数を制限したもの  
merge_not_concurrentファイルがマージソートを並列化していないもの  
それぞれのプログラムはマージソートした結果(データ)をresult.txtファイルに書き出す  
ソートするデータはrandom_dataファイルに入っている  
各プログラムの性能結果はエクセルファイルとしてmerge_resultファイルに入っている  

##175774B  
