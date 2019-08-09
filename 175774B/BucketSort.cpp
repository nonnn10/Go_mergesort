#include <omp.h>
#include <stdio.h>
#include <stdlib.h>
#include<iostream>
#include<fstream>

using namespace std;
void OutputTextFile(string filename,int a[],int elememns);

struct bucket {
	int n_elem; //要素数
	int index; //インデックス
	int start; //スタート地点
};
//小さい順に並び替える関数
int cmpfunc (const void * a, const void * b){return ( *(int*)a - *(int*)b );}

int main(int argc, char *argv[]) {

    int n_buckets,threads,elements,limit;//バケツの数，スレッド数，要素数，最大値
	int *unsorted_array, *sorted_array;
	int  i, w, j; 
	struct bucket* buckets; //バケツ
	double start_time,end_time,total_time;	//スタート時間，終わりの時間，トータルの時間

	printf("配列の要素数を入力してください \n");
	if (scanf("%d", &elements) != 1){
		printf("error\n");
		return -1;
	}
	printf("最大値を入力してください \n");
	if (scanf("%d", &limit) != 1){
		printf("error\n");
		return -1;
	}
	printf("バケツの数を入力してください \n");
	if (scanf("%d", &n_buckets) != 1){
		printf("error\n");
		return -1;
	}
    printf("スレッドの数を入力してください \n");
	if (scanf("%d", &threads) != 1){
		printf("error\n");
		return -1;
	}

	//バケツ１つの数の範囲
	w = (int)limit/n_buckets;

	//メモリ確保
	unsorted_array = (int *) malloc(sizeof(int)*elements);
	sorted_array = (int *) malloc(sizeof(int)*elements);
	buckets = (struct bucket *) calloc(n_buckets, sizeof(struct bucket));

    omp_set_num_threads(threads); //         ////////////////////////////////////スレッド数を変えるところ//////////////////////
    printf("スレッド数は%d\n",omp_get_max_threads());

	//乱数を生成して配列に格納
	for(i=0;i<elements;i++) {
		unsorted_array[i] = random() % limit+1;
	}
	//ソート前の配列をテキストファイルに出力
	OutputTextFile("unSorted.txt",unsorted_array,elements);


// ****************************
// バケツソート開始
// ****************************

	start_time = omp_get_wtime();//時間の計測を開始

	//配列を各バケツに分割
	for (i=0; i<elements;i++){
		j = unsorted_array[i]/w;
		if (j>n_buckets-1)
				j = n_buckets-1;
		buckets[j].n_elem++;
	}

	//バケツのインデックスとスタート地点を記憶させておく
	for (i=1; i<n_buckets;i++){
		buckets[i].index = buckets[i-1].index + buckets[i-1].n_elem;
		buckets[i].start = buckets[i-1].start + buckets[i-1].n_elem;
	}
	int b_index;

	//インデックスを使う準備
	for (i=0; i<elements;i++){
		j = unsorted_array[i]/w;
		if (j > n_buckets -1)
				j = n_buckets-1;
		b_index = buckets[j].index++;
		sorted_array[b_index] = unsorted_array[i];
	}

	//各バケツごとにソート（並列処理)
	#pragma omp parallel for 
	for(i=0; i<n_buckets; i++){
		qsort(sorted_array+buckets[i].start, buckets[i].n_elem, sizeof(int), cmpfunc);
    }

	end_time=omp_get_wtime();	//時間の計測を終了
	total_time = end_time - start_time;	//トータルの時間算出

	//ソートした配列をテキストファイルに出力
	OutputTextFile("Sorted.txt",sorted_array,elements);
	printf(" %d 個の要素をバケツソートするのに %f 秒\n", elements,total_time);

}

//配列をテキストファイルに出力する関数
void OutputTextFile(string filename,int a[],int elements)
{
    int i;
	ofstream outputfile(filename);
    for(i=0;i<elements;i++) {
        outputfile<<a[i];
        outputfile<<"\n";
	}
    outputfile.close();
    cout<<filename<<"に出力しました\n";
}