mergesortの並列化
================
Overview  
Golangにおけるマージソートの並列化を行なったもの

##Description
merge.goファイルがマージソートの並列化を実践したもの  
merge_limitファイルがマージソートにおける同時に実行できるgoroutine数を制限したもの  
merge_not_concurrentファイルがマージソートを並列化していないもの  
それぞれのプログラムはマージソートした結果(データ)をresult.txtファイルに書き出す  
  
ソートするデータはrandom_dataファイルに入っている  
各プログラムの性能結果はエクセルファイルとしてmerge_resultファイルに入っている  